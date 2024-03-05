package config

import (
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
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
	if err := validateGeneralConfigs(config); err != nil {
		return err
	}

	if err := validatePlugins(config); err != nil {
		return err
	}

	if err := validateServices(config); err != nil {
		return err
	}

	if err := validateRoutes(config); err != nil {
		return err
	}

	return nil
}

func validateGeneralConfigs(config *models.ProxyConfig) error {
	if config.Global.Name == "" {
		return fmt.Errorf("global name is required")
	}
	return nil
}

func validatePlugins(config *models.ProxyConfig) error {
	// Aggregate plugins in the config
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

		if !util.SliceContainsString(constant.AVAILABLE_PLUGINS, name) {
			return fmt.Errorf("plugin name is invalid. "+
				"Available plugins: %v", constant.AVAILABLE_PLUGINS)
		}
	}

	return nil
}

func validateServices(config *models.ProxyConfig) error {
	var serviceNames []string
	var servicePaths []string
	var availableProtocols = []string{"http", "https"}

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

		if !util.SliceContainsString(availableProtocols, service.Protocol) {
			return fmt.Errorf("service protocol is invalid, only http and https supported")
		}

		if len(service.Upstreams) == 0 {
			return fmt.Errorf("service must have at least one upstream")
		}

		serviceNames = append(serviceNames, service.Name)
		servicePaths = append(servicePaths, service.BasePath)
	}
	return nil
}

func validateRoutes(config *models.ProxyConfig) error {
	var validMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}

	for _, service := range config.Services {
		var routeNames []string
		for _, route := range service.Routes {
			// Path
			if util.SliceContainsString(routeNames, route.Path) {
				return fmt.Errorf("route path must be unique")
			}

			if !strings.HasPrefix(route.Path, "/") {
				return fmt.Errorf("route path must start with /")
			}

			if strings.HasSuffix(route.Path, "/") {
				return fmt.Errorf("route path must not end with /")
			}

			// Methods
			if len(route.Methods) == 0 {
				return fmt.Errorf("route methods must be specified")
			}

			for _, method := range route.Methods {
				if !util.SliceContainsString(validMethods, method) {
					return fmt.Errorf("route method is invalid")
				}
			}

			routeNames = append(routeNames, route.Path)
		}
	}

	return nil
}
