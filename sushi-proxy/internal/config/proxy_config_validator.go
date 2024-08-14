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

// TODO: validate load balancer configs

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

		// TODO: validate each plugin data/config schema
	}

	return nil
}

func validateServices(config *models.ProxyConfig) error {
	var serviceNames []string
	var servicePaths []string
	var availableProtocols = []string{"http", "https"}

	for _, service := range config.Services {
		// Name
		if util.SliceContainsString(serviceNames, service.Name) {
			return fmt.Errorf("service name: %s must be unique", service.Name)
		}

		// Path
		if util.SliceContainsString(servicePaths, service.BasePath) {
			return fmt.Errorf("service path: %s must be unique", service.BasePath)
		}

		if !strings.HasPrefix(service.BasePath, "/") {
			return fmt.Errorf("service path: %s must start with /", service.BasePath)
		}

		if strings.HasSuffix(service.BasePath, "/") {
			return fmt.Errorf("service path: %s must not end with /", service.BasePath)
		}

		// Protocol
		if !util.SliceContainsString(availableProtocols, service.Protocol) {
			return fmt.Errorf("service protocol: %s is invalid, "+
				"only http and https supported", service.Protocol)
		}

		// Upstreams
		if len(service.Upstreams) == 0 {
			return fmt.Errorf("service :%s must have at least one upstream", service.Name)
		}

		serviceNames = append(serviceNames, service.Name)
		servicePaths = append(servicePaths, service.BasePath)
	}
	return nil
}

func validateRoutes(config *models.ProxyConfig) error {

	var validMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}

	for _, service := range config.Services {
		var routePaths []string
		var routeNames []string
		for _, route := range service.Routes {
			// Path
			if util.SliceContainsString(routePaths, route.Path) {
				return fmt.Errorf("route path: %s must be unique", route.Path)
			}

			// Name
			if util.SliceContainsString(routeNames, route.Name) {
				return fmt.Errorf("route name: %s must be unique", route.Name)
			}

			if !strings.HasPrefix(route.Path, "/") {
				return fmt.Errorf("route path: %s must start with /", route.Path)
			}

			if strings.HasSuffix(route.Path, "/") {
				return fmt.Errorf("route path: %s must not end with /", route.Path)
			}

			// Methods
			if len(route.Methods) == 0 {
				return fmt.Errorf("route methods must be specified")
			}

			for _, method := range route.Methods {
				if !util.SliceContainsString(validMethods, method) {
					return fmt.Errorf("route method: %s is invalid", method)
				}
			}

			routePaths = append(routePaths, route.Path)
			routeNames = append(routeNames, route.Name)
		}
	}

	return nil
}
