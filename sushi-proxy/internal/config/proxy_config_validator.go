package config

import (
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"strings"
)

func ValidateAndParseSchema(raw []byte) (*models.ProxyConfig, error) {
	var config models.ProxyConfig
	err := json.Unmarshal(raw, &config)
	if err != nil {
		slog.Info("Error parsing config file", err)
		return nil, err
	}

	return &config, nil
}

func ValidateConfig(config *models.ProxyConfig) error {
	err := validateGeneralConfigs(config)
	err = validatePlugins(config)
	err = validateServices(config)
	err = validateRoutes(config)
	return err
}

func validateGeneralConfigs(config *models.ProxyConfig) error {
	if config.Global.Name == "" {
		return fmt.Errorf("global name is required")
	}
	return nil
}

func validatePlugins(config *models.ProxyConfig) error {
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

func validateServices(config *models.ProxyConfig) error {
	var serviceNames []string
	var servicePaths []string
	for _, service := range config.Services {
		if util.SliceContainsString(serviceNames, service.Name) {
			return fmt.Errorf("service name must be unique")
		}

		if util.SliceContainsString(servicePaths, service.BasePath) {
			return fmt.Errorf("service path must be unique")
		}

		if !strings.HasPrefix(service.BasePath, "/") {
			return fmt.Errorf("service path must start with /")
		}

		if strings.HasSuffix(service.BasePath, "/") {
			return fmt.Errorf("service path must not end with /")
		}

		serviceNames = append(serviceNames, service.Name)
		servicePaths = append(servicePaths, service.BasePath)
	}
	return nil
}

func validateRoutes(config *models.ProxyConfig) error {
	for _, service := range config.Services {
		var routeNames []string
		for _, route := range service.Routes {
			if util.SliceContainsString(routeNames, route.Path) {
				return fmt.Errorf("route name must be unique")
			}

			if len(route.Methods) == 0 {
				return fmt.Errorf("route methods must be specified")
			}

			if !strings.HasPrefix(route.Path, "/") {
				return fmt.Errorf("route path must start with /")
			}

			if strings.HasSuffix(route.Path, "/") {
				return fmt.Errorf("route path must not end with /")
			}

			routeNames = append(routeNames, route.Path)
		}
	}

	return nil
}
