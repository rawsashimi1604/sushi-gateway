package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"log"
)

type ServiceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (serviceRepo *ServiceRepository) GetAllServices() ([]gateway.Service, error) {

	var services []gateway.Service
	serviceRows, err := serviceRepo.db.Query("SELECT name, base_path, protocol, load_balancing_alg FROM service")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch services: %w", err)
	}
	defer serviceRows.Close()

	for serviceRows.Next() {
		var service gateway.Service
		err := serviceRows.Scan(&service.Name, &service.BasePath, &service.Protocol, &service.LoadBalancingStrategy)
		if err != nil {
			log.Printf("failed to scan service: %v\n", err)
			continue
		}

		upstreamQuery := `SELECT id, host, port FROM upstream WHERE service_name = $1`
		upstreamRows, err := serviceRepo.db.Query(upstreamQuery, service.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch upstreams for service %s: %w", service.Name, err)
		}
		for upstreamRows.Next() {
			var upstream gateway.Upstream
			if err := upstreamRows.Scan(&upstream.Id, &upstream.Host, &upstream.Port); err != nil {
				log.Printf("failed to scan upstream: %v\n", err)
				continue
			}
			service.Upstreams = append(service.Upstreams, upstream)
		}
		upstreamRows.Close()

		routeQuery := `SELECT name, path FROM route WHERE service_name = $1`
		routeRows, err := serviceRepo.db.Query(routeQuery, service.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch routes for service %s: %w", service.Name, err)
		}
		for routeRows.Next() {
			var route gateway.Route
			if err := routeRows.Scan(&route.Name, &route.Path); err != nil {
				log.Printf("failed to scan route: %v\n", err)
				continue
			}

			methodQuery := `SELECT method FROM route_methods WHERE route_name = $1`
			methodRows, err := serviceRepo.db.Query(methodQuery, route.Name)
			if err != nil {
				log.Printf("failed to fetch methods for route %s: %v\n", route.Name, err)
				continue
			}
			for methodRows.Next() {
				var method string
				if err := methodRows.Scan(&method); err != nil {
					log.Printf("failed to scan method: %v\n", err)
					continue
				}
				route.Methods = append(route.Methods, method)
			}
			methodRows.Close()

			routePluginQuery := `SELECT p.id, p.name, p.config, p.enabled 
				FROM plugin p 
				JOIN route_plugin rp ON p.id = rp.plugin_id 
				WHERE rp.route_name = $1`
			pluginRows, err := serviceRepo.db.Query(routePluginQuery, route.Name)
			if err != nil {
				log.Printf("failed to fetch plugins for route %s: %v\n", route.Name, err)
				continue
			}
			for pluginRows.Next() {
				var plugin gateway.PluginConfig
				var configBytes []byte // Temporary byte slice to hold the JSON data

				// Scan the basic fields and the raw JSON into configBytes
				if err := pluginRows.Scan(&plugin.Id, &plugin.Name, &configBytes, &plugin.Enabled); err != nil {
					log.Printf("failed to scan plugin: %v\n", err)
					continue
				}

				// Unmarshal the JSON bytes into the plugin's Config map
				if err := json.Unmarshal(configBytes, &plugin.Config); err != nil {
					log.Printf("failed to unmarshal plugin config: %v\n", err)
					continue
				}
				route.Plugins = append(route.Plugins, plugin)
			}
			pluginRows.Close()

			service.Routes = append(service.Routes, route)
		}
		routeRows.Close()

		servicePluginQuery := `SELECT p.id, p.name, p.config, p.enabled 
				FROM plugin p 
				JOIN service_plugin sp ON p.id = sp.plugin_id 
				WHERE sp.service_name = $1`
		servicePluginRows, err := serviceRepo.db.Query(servicePluginQuery, service.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch plugins for service %s: %w", service.Name, err)
		}
		for servicePluginRows.Next() {
			var plugin gateway.PluginConfig
			var configBytes []byte // Temporary byte slice to hold the JSON data

			// Scan the basic fields and the raw JSON into configBytes
			if err := servicePluginRows.Scan(&plugin.Id, &plugin.Name, &configBytes, &plugin.Enabled); err != nil {
				log.Printf("failed to scan plugin: %v\n", err)
				continue
			}

			// Unmarshal the JSON bytes into the plugin's Config map
			if err := json.Unmarshal(configBytes, &plugin.Config); err != nil {
				log.Printf("failed to unmarshal plugin config: %v\n", err)
				continue
			}
			service.Plugins = append(service.Plugins, plugin)
		}
		servicePluginRows.Close()
		services = append(services, service)
	}

	// Ensure that plugins are not nil, and initialize as empty arrays where necessary
	for i := range services {
		if services[i].Plugins == nil {
			services[i].Plugins = []gateway.PluginConfig{}
		}
		for j := range services[i].Routes {
			if services[i].Routes[j].Plugins == nil {
				services[i].Routes[j].Plugins = []gateway.PluginConfig{}
			}
		}
	}

	return services, nil
}

func (serviceRepo *ServiceRepository) AddService(service gateway.Service) error {
	tx, err := serviceRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Insert the service into the service table
	serviceInsertQuery := `INSERT INTO service (name, base_path, protocol, load_balancing_alg) 
						   VALUES ($1, $2, $3, $4)`
	_, err = tx.Exec(serviceInsertQuery, service.Name, service.BasePath, service.Protocol, service.LoadBalancingStrategy)
	if err != nil {
		return fmt.Errorf("failed to insert service: %w", err)
	}

	// Insert upstreams for the service
	upstreamInsertQuery := `INSERT INTO upstream (id, service_name, host, port) VALUES ($1, $2, $3, $4)`
	for _, upstream := range service.Upstreams {
		_, err = tx.Exec(upstreamInsertQuery, upstream.Id, service.Name, upstream.Host, upstream.Port)
		if err != nil {
			return fmt.Errorf("failed to insert upstream for service %s: %w", service.Name, err)
		}
	}

	// Insert routes for the service
	routeInsertQuery := `INSERT INTO route (name, service_name, path) VALUES ($1, $2, $3)`
	for _, route := range service.Routes {
		_, err = tx.Exec(routeInsertQuery, route.Name, service.Name, route.Path)
		if err != nil {
			return fmt.Errorf("failed to insert route for service %s: %w", service.Name, err)
		}

		// Insert route methods
		methodInsertQuery := `INSERT INTO route_methods (route_name, method) VALUES ($1, $2)`
		for _, method := range route.Methods {
			_, err = tx.Exec(methodInsertQuery, route.Name, method)
			if err != nil {
				return fmt.Errorf("failed to insert method for route %s: %w", route.Name, err)
			}
		}

		// Insert route-level plugins using the route_plugin table
		routePluginMappingQuery := `INSERT INTO route_plugin (route_name, plugin_id) VALUES ($1, $2)`
		for _, plugin := range route.Plugins {
			pluginConfig, err := json.Marshal(plugin.Config)
			if err != nil {
				return fmt.Errorf("failed to marshal plugin config for route %s: %w", route.Name, err)
			}

			// Insert or update the plugin in the plugin table
			routePluginInsertQuery := `INSERT INTO plugin (id, name, config, enabled) 
									   VALUES ($1, $2, $3, $4)
									   `
			_, err = tx.Exec(routePluginInsertQuery, plugin.Id, plugin.Name, pluginConfig, plugin.Enabled)
			if err != nil {
				return fmt.Errorf("failed to insert plugin for route %s: %w", route.Name, err)
			}

			// Associate plugin with route
			_, err = tx.Exec(routePluginMappingQuery, route.Name, plugin.Id)
			if err != nil {
				return fmt.Errorf("failed to associate plugin with route %s: %w", route.Name, err)
			}
		}
	}

	// Insert service-level plugins using the service_plugin table
	servicePluginMappingQuery := `INSERT INTO service_plugin (service_name, plugin_id) VALUES ($1, $2)`
	for _, plugin := range service.Plugins {
		pluginConfig, err := json.Marshal(plugin.Config)
		if err != nil {
			return fmt.Errorf("failed to marshal plugin config for service %s: %w", service.Name, err)
		}

		// Insert or update the plugin in the plugin table
		servicePluginInsertQuery := `INSERT INTO plugin (id, name, config, enabled) 
									 VALUES ($1, $2, $3, $4)
									 `
		_, err = tx.Exec(servicePluginInsertQuery, plugin.Id, plugin.Name, pluginConfig, plugin.Enabled)
		if err != nil {
			return fmt.Errorf("failed to insert plugin for service %s: %w", service.Name, err)
		}

		// Associate plugin with service
		_, err = tx.Exec(servicePluginMappingQuery, service.Name, plugin.Id)
		if err != nil {
			return fmt.Errorf("failed to associate plugin with service %s: %w", service.Name, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeleteServiceByName deletes a service and all associated data (handled by DELETE CASCADE in DB)
func (serviceRepo *ServiceRepository) DeleteServiceByName(serviceName string) error {
	tx, err := serviceRepo.db.Begin() // Start a transaction
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Rollback if an error occurs
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Delete the service; CASCADE will delete associated upstreams, routes, plugins, etc.
	serviceDeleteQuery := `DELETE FROM service WHERE name = $1`
	_, err = tx.Exec(serviceDeleteQuery, serviceName)
	if err != nil {
		return fmt.Errorf("failed to delete service %s: %w", serviceName, err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
