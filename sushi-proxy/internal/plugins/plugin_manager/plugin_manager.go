package plugin_manager

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/acl"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/basic_auth"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"net/http"
	"sort"
)

type PluginManager struct {
	plugins []*plugins.Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]*plugins.Plugin, 0),
	}
}

func NewPluginManagerFromConfig(req *http.Request) (*PluginManager, *errors.HttpError) {
	// Load the plugin configuration from the config file
	// Based on the plugins loaded from http request
	// Order of precedence: route.plugins > service.plugins > global.plugins
	pm := NewPluginManager()

	globalPlugins := config.GlobalProxyConfig.Global.Plugins
	for _, pluginConfig := range globalPlugins {
		err := pm.loadConfig(pluginConfig)
		if err != nil {
			return nil, err
		}
	}

	// Search for service.plugins and route.plugins
	service, route, err := util.GetServiceAndRouteFromRequest(req)
	if err != nil {
		return nil, err
	}

	for _, pluginConfig := range route.Plugins {
		err := pm.loadConfig(pluginConfig)
		if err != nil {
			return nil, err
		}
	}

	for _, pluginConfig := range service.Plugins {
		err := pm.loadConfig(pluginConfig)
		if err != nil {
			return nil, err
		}
	}

	return pm, nil
}

// Load the plugin configuration from the config file
func (pm *PluginManager) loadConfig(pc models.PluginConfig) *errors.HttpError {
	name, ok := pc["name"].(string)
	if !ok {
		return errors.NewHttpError(http.StatusInternalServerError, "PLUGIN_CONFIG_ERROR",
			"Plugin name not found")
	}

	switch name {
	case constant.PLUGIN_BASIC_AUTH:
		pm.RegisterPlugin(basic_auth.NewBasicAuthPlugin(pc))
	case constant.PLUGIN_ACL:
		pm.RegisterPlugin(acl.NewAclPlugin(pc))
	}
	return nil
}

func (pm *PluginManager) RegisterPlugin(plugin *plugins.Plugin) {
	// If plugins already exists replcae
	exists := false
	for i, p := range pm.plugins {
		if p.Name == plugin.Name {
			pm.plugins[i] = plugin
			exists = true
			break
		}
	}

	if !exists {
		// Append the new plugin
		pm.plugins = append(pm.plugins, plugin)
	}

	// Sort the plugins by priority, higher priority executes first
	sort.Slice(pm.plugins, func(i, j int) bool {
		return pm.plugins[i].Priority < pm.plugins[j].Priority
	})
}

// ExecutePlugins chains the plugins and returns a single http.Handler
// finalHandler is the application's main handler that should execute after all plugins
func (pm *PluginManager) ExecutePlugins(finalHandler http.Handler) http.Handler {

	// Chain the plugins order to ensure the correct execution sequence
	for _, plugin := range pm.plugins {
		finalHandler = plugin.Handler.Execute(finalHandler)
	}

	return finalHandler
}

func (pm *PluginManager) GetPlugins() []*plugins.Plugin {
	return pm.plugins
}
