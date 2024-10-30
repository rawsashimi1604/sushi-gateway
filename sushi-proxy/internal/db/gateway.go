package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

type GatewayRepository struct {
	db *sql.DB
}

func NewGatewayRepository(db *sql.DB) *GatewayRepository {
	return &GatewayRepository{db: db}
}

// GetGatewayInfo retrieves the gateway information and associated global plugins from the "gateway" table.
func (gatewayRepo *GatewayRepository) GetGatewayInfo() (model.Global, error) {
	var gatewayInfo model.Global

	// Query to fetch the gateway name
	query := `SELECT name FROM gateway LIMIT 1`
	err := gatewayRepo.db.QueryRow(query).Scan(&gatewayInfo.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Global{}, fmt.Errorf("no gateway information found")
		}
		return model.Global{}, fmt.Errorf("failed to fetch gateway information: %w", err)
	}

	// Initialize Plugins as an empty slice to prevent nil references
	gatewayInfo.Plugins = []model.PluginConfig{}

	// Query to fetch all global plugins
	pluginQuery := `SELECT id, name, config, enabled FROM plugin WHERE scope = 'global'`
	rows, err := gatewayRepo.db.Query(pluginQuery)
	if err != nil {
		return model.Global{}, fmt.Errorf("failed to fetch global plugins: %w", err)
	}
	defer rows.Close()

	// Populate global plugins
	for rows.Next() {
		var plugin model.PluginConfig
		var configBytes []byte

		if err := rows.Scan(&plugin.Id, &plugin.Name, &configBytes, &plugin.Enabled); err != nil {
			return model.Global{}, fmt.Errorf("failed to scan plugin: %w", err)
		}

		// Unmarshal JSON config into the plugin's Config map
		if err := json.Unmarshal(configBytes, &plugin.Config); err != nil {
			return model.Global{}, fmt.Errorf("failed to unmarshal plugin config: %w", err)
		}

		gatewayInfo.Plugins = append(gatewayInfo.Plugins, plugin)
	}

	// Check for any errors during iteration
	if err = rows.Err(); err != nil {
		return model.Global{}, fmt.Errorf("error during global plugins iteration: %w", err)
	}

	return gatewayInfo, nil
}
