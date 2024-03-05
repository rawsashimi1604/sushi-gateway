package config

import (
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"log/slog"
)

func validateAndParseSchema(raw []byte) (*ProxyConfig, error) {
	var config ProxyConfig
	err := json.Unmarshal(raw, &config)
	if err != nil {
		slog.Info("Error parsing config file", err)
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *ProxyConfig) error {
	err := validateGeneralConfigs(config)
	err = validatePlugins(config)
	return err
}

func validateGeneralConfigs(config *ProxyConfig) error {
	if config.Global.Name == "" {
		return fmt.Errorf("global name is required")
	}
	return nil
}

func validatePlugins(config *ProxyConfig) error {
	// Get all plugins in the config
	var plugins []models.PluginConfig
	for _, globalPlugin := range config.Global.Plugins {
		plugins = append(plugins, globalPlugin)
	}
	for _, service := range config.Services {
		for _, servicePlugin := range service.Plugins {
			plugins = append(plugins, servicePlugin)
		}

		for _, route := range service.Routes {
			for _, routePlugin := range route.Plugins {
				plugins = append(plugins, routePlugin)
			}
		}
	}

	// Validate each plugin
	for _, plugin := range plugins {
		name, nameOk := plugin["name"].(string)
		if !nameOk || name == "" {
			return fmt.Errorf("plugin name is required")
		}
	}

	return nil
}
