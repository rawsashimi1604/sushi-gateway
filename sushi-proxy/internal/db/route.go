package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

type RouteRepository struct {
	db *sql.DB
}

func NewRouteRepository(db *sql.DB) *RouteRepository {
	return &RouteRepository{db: db}
}

// GetAllRoutes retrieves all routes from the database, including their methods and plugins.
func (routeRepo *RouteRepository) GetAllRoutes(serviceName string) ([]model.Route, error) {
	var routes []model.Route

	// 1. Query to get all routes associated with the service
	routeQuery := `SELECT name, path FROM route WHERE service_name = $1`
	routeRows, err := routeRepo.db.Query(routeQuery, serviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch routes: %w", err)
	}
	defer routeRows.Close()

	for routeRows.Next() {
		var route model.Route
		err := routeRows.Scan(&route.Name, &route.Path)
		if err != nil {
			return nil, fmt.Errorf("failed to scan route: %w", err)
		}

		// 2. Query to get the methods for the current route
		methodQuery := `SELECT method FROM route_methods WHERE route_name = $1`
		methodRows, err := routeRepo.db.Query(methodQuery, route.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch methods for route %s: %w", route.Name, err)
		}
		defer methodRows.Close()

		for methodRows.Next() {
			var method string
			err := methodRows.Scan(&method)
			if err != nil {
				return nil, fmt.Errorf("failed to scan method: %w", err)
			}
			route.Methods = append(route.Methods, method)
		}

		// 3. Query to get the plugins for the current route
		pluginQuery := `
			SELECT p.id, p.name, p.config, p.enabled
			FROM plugin p
			JOIN route_plugin rp ON p.id = rp.plugin_id
			WHERE rp.route_name = $1`
		pluginRows, err := routeRepo.db.Query(pluginQuery, route.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch plugins for route %s: %w", route.Name, err)
		}
		defer pluginRows.Close()

		for pluginRows.Next() {
			var plugin model.PluginConfig
			var configBytes []byte // Temporary variable to hold JSON bytes

			// Scan the plugin data and config JSON
			if err := pluginRows.Scan(&plugin.Id, &plugin.Name, &configBytes, &plugin.Enabled); err != nil {
				return nil, fmt.Errorf("failed to scan plugin for route %s: %w", route.Name, err)
			}

			// Unmarshal JSON config into plugin's config map
			if err := json.Unmarshal(configBytes, &plugin.Config); err != nil {
				return nil, fmt.Errorf("failed to unmarshal plugin config for route %s: %w", route.Name, err)
			}

			route.Plugins = append(route.Plugins, plugin)
		}

		// Add the route with its methods and plugins to the list of routes
		routes = append(routes, route)
	}

	return routes, nil
}

func (routeRepo *RouteRepository) AddRoute(serviceName string, route model.Route) error {
	tx, err := routeRepo.db.Begin() // Start a transaction
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Insert route
	routeInsertQuery := `INSERT INTO route (name, service_name, path) VALUES ($1, $2, $3)`
	_, err = tx.Exec(routeInsertQuery, route.Name, serviceName, route.Path)
	if err != nil {
		return fmt.Errorf("failed to insert route: %w", err)
	}

	// 2. Insert methods for the route
	methodInsertQuery := `INSERT INTO route_methods (route_name, method) VALUES ($1, $2)`
	for _, method := range route.Methods {
		_, err = tx.Exec(methodInsertQuery, route.Name, method)
		if err != nil {
			return fmt.Errorf("failed to insert method for route %s: %w", route.Name, err)
		}
	}

	// 3. Insert plugins for the route
	routePluginInsertQuery := `INSERT INTO plugin (id, name, config, enabled, scope) 
                               VALUES ($1, $2, $3, $4, 'route')
                               `
	routePluginMappingQuery := `INSERT INTO route_plugin (route_name, plugin_id) VALUES ($1, $2)`
	for _, plugin := range route.Plugins {
		pluginConfig, err := json.Marshal(plugin.Config)
		if err != nil {
			return fmt.Errorf("failed to marshal plugin config for route %s: %w", route.Name, err)
		}

		_, err = tx.Exec(routePluginInsertQuery, plugin.Id, plugin.Name, pluginConfig, plugin.Enabled)
		if err != nil {
			return fmt.Errorf("failed to insert plugin for route %s: %w", route.Name, err)
		}

		_, err = tx.Exec(routePluginMappingQuery, route.Name, plugin.Id)
		if err != nil {
			return fmt.Errorf("failed to associate plugin with route %s: %w", route.Name, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeleteRoute deletes a route, including its methods and plugins.
func (routeRepo *RouteRepository) DeleteRoute(routeName string) error {
	tx, err := routeRepo.db.Begin() // Start a transaction
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Delete plugins associated with the route
	routePluginDeleteQuery := `DELETE FROM route_plugin WHERE route_name = $1`
	_, err = tx.Exec(routePluginDeleteQuery, routeName)
	if err != nil {
		return fmt.Errorf("failed to delete plugins for route %s: %w", routeName, err)
	}

	// 2. Delete methods associated with the route
	methodDeleteQuery := `DELETE FROM route_methods WHERE route_name = $1`
	_, err = tx.Exec(methodDeleteQuery, routeName)
	if err != nil {
		return fmt.Errorf("failed to delete methods for route %s: %w", routeName, err)
	}

	// 3. Delete the route
	routeDeleteQuery := `DELETE FROM route WHERE name = $1`
	_, err = tx.Exec(routeDeleteQuery, routeName)
	if err != nil {
		return fmt.Errorf("failed to delete route %s: %w", routeName, err)
	}

	// 4. Cleanup orphaned plugins
	orphanedPluginCleanupQuery := `
		DELETE FROM plugin 
		WHERE id NOT IN (SELECT plugin_id FROM route_plugin)
		AND id NOT IN (SELECT plugin_id FROM service_plugin)
		AND scope != 'global'`
	_, err = tx.Exec(orphanedPluginCleanupQuery)
	if err != nil {
		return fmt.Errorf("failed to cleanup orphaned plugins: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
