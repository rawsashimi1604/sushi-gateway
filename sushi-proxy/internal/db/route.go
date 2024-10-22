package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
)

type RouteRepository struct {
	db *sql.DB
}

func NewRouteRepository(db *sql.DB) *RouteRepository {
	return &RouteRepository{db: db}
}

func (routeRepo *RouteRepository) AddRoute(serviceName string, route gateway.Route) error {
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
	routePluginInsertQuery := `INSERT INTO plugin (id, name, config, enabled) 
                               VALUES ($1, $2, $3, $4)
                               ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, config = EXCLUDED.config, enabled = EXCLUDED.enabled`
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

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
