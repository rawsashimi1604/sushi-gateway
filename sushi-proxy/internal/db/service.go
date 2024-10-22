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

	return services, nil
}
