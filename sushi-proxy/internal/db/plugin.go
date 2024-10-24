package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"log"
)

type PluginRepository struct {
	db *sql.DB
}

func NewPluginRepository(db *sql.DB) *PluginRepository {
	return &PluginRepository{db: db}
}

// TODO: add methods
// GetPlugins fetches plugins based on the scope: global, service, or route
func (pluginRepo *PluginRepository) GetPlugins(scope string, targetName string) ([]model.PluginConfig, error) {
	var plugins []model.PluginConfig
	var query string

	switch scope {
	case "global":
		// Fetch global plugins
		query = `SELECT id, name, config, enabled FROM plugin WHERE scope = 'global'`
	case "service":
		// Fetch service-level plugins associated with a specific service
		query = `SELECT p.id, p.name, p.config, p.enabled 
				 FROM plugin p 
				 JOIN service_plugin sp ON p.id = sp.plugin_id 
				 WHERE sp.service_name = $1 AND p.scope = 'service'`
	case "route":
		// Fetch route-level plugins associated with a specific route
		query = `SELECT p.id, p.name, p.config, p.enabled 
				 FROM plugin p 
				 JOIN route_plugin rp ON p.id = rp.plugin_id 
				 WHERE rp.route_name = $1 AND p.scope = 'route'`
	default:
		return nil, fmt.Errorf("invalid scope: %s", scope)
	}

	rows, err := pluginRepo.db.Query(query, targetName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch plugins: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var plugin model.PluginConfig
		var configBytes []byte

		err := rows.Scan(&plugin.Id, &plugin.Name, &configBytes, &plugin.Enabled)
		if err != nil {
			log.Printf("failed to scan plugin: %v\n", err)
			continue
		}

		err = json.Unmarshal(configBytes, &plugin.Config)
		if err != nil {
			log.Printf("failed to unmarshal plugin config: %v\n", err)
			continue
		}

		plugins = append(plugins, plugin)
	}

	return plugins, nil
}

// AddPlugin adds a new plugin at the global, service, or route level
func (pluginRepo *PluginRepository) AddPlugin(scope string, plugin model.PluginConfig, targetName string) error {
	tx, err := pluginRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	pluginConfig, err := json.Marshal(plugin.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal plugin config: %w", err)
	}

	// Insert the plugin into the plugin table
	pluginInsertQuery := `INSERT INTO plugin (id, name, config, enabled, scope) 
						  VALUES ($1, $2, $3, $4, $5)
						  `
	_, err = tx.Exec(pluginInsertQuery, plugin.Id, plugin.Name, pluginConfig, plugin.Enabled, scope)
	if err != nil {
		return fmt.Errorf("failed to insert or update plugin: %w", err)
	}

	// Associate the plugin based on the scope
	switch scope {
	case "global":
		// No need to associate global plugins with a service or route
	case "service":
		// Insert into service_plugin table
		servicePluginInsertQuery := `INSERT INTO service_plugin (service_name, plugin_id) VALUES ($1, $2)`
		_, err = tx.Exec(servicePluginInsertQuery, targetName, plugin.Id)
		if err != nil {
			return fmt.Errorf("failed to associate plugin with service %s: %w", targetName, err)
		}
	case "route":
		// Insert into route_plugin table
		routePluginInsertQuery := `INSERT INTO route_plugin (route_name, plugin_id) VALUES ($1, $2)`
		_, err = tx.Exec(routePluginInsertQuery, targetName, plugin.Id)
		if err != nil {
			return fmt.Errorf("failed to associate plugin with route %s: %w", targetName, err)
		}
	default:
		return fmt.Errorf("invalid scope: %s", scope)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeletePlugin deletes a plugin by name from the global, service, or route level
func (pluginRepo *PluginRepository) DeletePlugin(scope string, pluginName string, targetName string) error {
	tx, err := pluginRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Delete plugin association based on the scope
	switch scope {
	case "global":
		// Global plugins are only in the plugin table, no associations
		pluginDeleteQuery := `DELETE FROM plugin WHERE name = $1 AND scope = 'global'`
		_, err = tx.Exec(pluginDeleteQuery, pluginName)
	case "service":
		// Delete service-level plugin association
		pluginDeleteQuery := `DELETE FROM service_plugin WHERE service_name = $1 AND plugin_id = 
							  (SELECT id FROM plugin WHERE name = $2 AND scope = 'service')`
		_, err = tx.Exec(pluginDeleteQuery, targetName, pluginName)
	case "route":
		// Delete route-level plugin association
		pluginDeleteQuery := `DELETE FROM route_plugin WHERE route_name = $1 AND plugin_id = 
							  (SELECT id FROM plugin WHERE name = $2 AND scope = 'route')`
		_, err = tx.Exec(pluginDeleteQuery, targetName, pluginName)
	default:
		return fmt.Errorf("invalid scope: %s", scope)
	}

	if err != nil {
		return fmt.Errorf("failed to delete plugin: %w", err)
	}

	// Delete the plugin itself from the plugin table if it exists
	pluginDeleteQuery := `DELETE FROM plugin WHERE name = $1 AND scope = $2`
	_, err = tx.Exec(pluginDeleteQuery, pluginName, scope)
	if err != nil {
		return fmt.Errorf("failed to delete plugin from plugin table: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
