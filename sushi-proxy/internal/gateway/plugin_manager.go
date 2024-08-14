package gateway

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"net/http"
	"sort"
)

type PluginManager struct {
	plugins []*Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]*Plugin, 0),
	}
}

func NewPluginManagerFromConfig(req *http.Request) (*PluginManager, *HttpError) {
	// Load the plugin configuration from the gateway file
	// Based on the plugins loaded from http request
	// Order of precedence: route.plugins > service.plugins > global.plugins
	pm := NewPluginManager()

	globalPlugins := GlobalProxyConfig.Global.Plugins
	for _, pluginConfig := range globalPlugins {
		err := pm.loadConfig(pluginConfig)
		if err != nil {
			return nil, err
		}
	}

	// Search for service.plugins and route.plugins
	service, route, err := GetServiceAndRouteFromRequest(&GlobalProxyConfig, req)
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

// Load the plugin configuration from the gateway file
func (pm *PluginManager) loadConfig(pc PluginConfig) *HttpError {
	name, ok := pc["name"].(string)
	if !ok {
		return NewHttpError(http.StatusInternalServerError, "PLUGIN_CONFIG_ERROR",
			"Plugin name not found")
	}

	enabled, enabledOk := pc["enabled"].(bool)
	if !enabledOk {
		return NewHttpError(http.StatusInternalServerError, "PLUGIN_CONFIG_ERROR",
			"Plugin enabled flag not found")
	}

	// Skip as not enabled.
	if !enabled {
		return nil
	}

	switch name {
	case constant.PLUGIN_BASIC_AUTH:
		pm.RegisterPlugin(NewBasicAuthPlugin(pc))
	case constant.PLUGIN_ACL:
		pm.RegisterPlugin(NewAclPlugin(pc))
	case constant.PLUGIN_BOT_PROTECTION:
		pm.RegisterPlugin(NewBotProtectionPlugin(pc))
	case constant.PLUGIN_KEY_AUTH:
		pm.RegisterPlugin(NewKeyAuthPlugin(pc))
	case constant.PLUGIN_RATE_LIMIT:
		pm.RegisterPlugin(NewRateLimitPlugin(pc))
	case constant.PLUGIN_REQUEST_SIZE_LIMIT:
		pm.RegisterPlugin(NewRequestSizeLimitPlugin(pc))
	case constant.PLUGIN_JWT:
		pm.RegisterPlugin(NewJwtPlugin(pc))
	case constant.PLUGIN_MTLS:
		pm.RegisterPlugin(NewMtlsPlugin())
	case constant.PLUGIN_HTTP_LOG:
		pm.RegisterPlugin(NewHttpLogPlugin(pc))
	case constant.PLUGIN_CORS:
		pm.RegisterPlugin(NewCorsPlugin(pc))
	}
	return nil
}

func (pm *PluginManager) RegisterPlugin(plugin *Plugin) {
	// If plugins already exists replace
	exists := false
	for i, p := range pm.plugins {
		if p.Name == plugin.Name {
			pm.plugins[i] = plugin
			exists = true
			break
		}
	}

	if !exists {
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
	for _, plugin := range pm.plugins {
		finalHandler = plugin.Handler.Execute(finalHandler)
	}
	return finalHandler
}

func (pm *PluginManager) GetPlugins() []*Plugin {
	return pm.plugins
}
