package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/validator"
	"log/slog"
)

func ValidateAndParseSchema(raw []byte) (*model.ProxyConfig, error) {
	var config model.ProxyConfig
	err := json.Unmarshal(raw, &config)
	if err != nil {
		slog.Info("Error parsing gateway file", err)
		return nil, err
	}

	return &config, nil
}

func ValidateConfig(config *model.ProxyConfig) error {
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

func validateGeneralConfigs(config *model.ProxyConfig) error {
	if config.Global.Name == "" {
		return fmt.Errorf("global name is required")
	}
	return nil
}

func validatePlugins(config *model.ProxyConfig) error {
	// Aggregate plugins in the gateway
	var plugins []model.PluginConfig
	pluginValidator := validator.NewPluginValidator()

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
		err := pluginValidator.ValidatePlugin(plugin)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateServices(config *model.ProxyConfig) error {
	var serviceNames []string
	var servicePaths []string
	serviceValidator := validator.NewServiceValidator()

	for _, service := range config.Services {
		// Name
		if util.SliceContainsString(serviceNames, service.Name) {
			return fmt.Errorf("service name: %s must be unique", service.Name)
		}

		// Path
		if util.SliceContainsString(servicePaths, service.BasePath) {
			return fmt.Errorf("service path: %s must be unique", service.BasePath)
		}

		// Generic service validations
		if err := serviceValidator.ValidateService(service); err != nil {
			return err
		}

		serviceNames = append(serviceNames, service.Name)
		servicePaths = append(servicePaths, service.BasePath)
	}
	return nil
}

func validateRoutes(config *model.ProxyConfig) error {

	for _, service := range config.Services {
		var routePaths []string
		var routeNames []string
		routeValidator := validator.NewRouteValidator()

		for _, route := range service.Routes {
			// Name
			if util.SliceContainsString(routeNames, route.Name) {
				return fmt.Errorf("route name: %s must be unique", route.Name)
			}

			// Generic route validations
			if err := routeValidator.ValidateRoute(route); err != nil {
				return err
			}

			routePaths = append(routePaths, route.Path)
			routeNames = append(routeNames, route.Name)
		}
	}

	return nil
}
