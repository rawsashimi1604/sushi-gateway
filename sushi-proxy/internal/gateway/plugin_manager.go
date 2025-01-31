package gateway

import (
	"net/http"
	"sort"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

type PluginManager struct {
	plugins []*Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]*Plugin, 0),
	}
}

func NewPluginManagerFromConfig(req *http.Request) (*PluginManager, *model.HttpError) {
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
	service, route, err := util.GetServiceAndRouteFromRequest(&GlobalProxyConfig, req)
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
func (pm *PluginManager) loadConfig(pc model.PluginConfig) *model.HttpError {
	// Skip as not enabled.
	if !pc.Enabled {
		return nil
	}

	switch pc.Name {
	case constant.PLUGIN_BASIC_AUTH:
		pm.RegisterPlugin(NewBasicAuthPlugin(pc.Config))
	case constant.PLUGIN_ACL:
		pm.RegisterPlugin(NewAclPlugin(pc.Config))
	case constant.PLUGIN_BOT_PROTECTION:
		pm.RegisterPlugin(NewBotProtectionPlugin(pc.Config))
	case constant.PLUGIN_KEY_AUTH:
		pm.RegisterPlugin(NewKeyAuthPlugin(pc.Config))
	case constant.PLUGIN_RATE_LIMIT:
		pm.RegisterPlugin(NewRateLimitPlugin(pc.Config, &GlobalProxyConfig))
	case constant.PLUGIN_REQUEST_SIZE_LIMIT:
		pm.RegisterPlugin(NewRequestSizeLimitPlugin(pc.Config))
	case constant.PLUGIN_JWT:
		pm.RegisterPlugin(NewJwtPlugin(pc.Config))
	case constant.PLUGIN_MTLS:
		pm.RegisterPlugin(NewMtlsPlugin(pc.Config))
	case constant.PLUGIN_HTTP_LOG:
		pm.RegisterPlugin(NewHttpLogPlugin(pc.Config))
	case constant.PLUGIN_CORS:
		pm.RegisterPlugin(NewCorsPlugin(pc.Config))
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
func (pm *PluginManager) ExecutePlugins(phase PluginPhase, finalHandler http.Handler) http.Handler {
	for _, plugin := range pm.plugins {
		if plugin.Phase == phase {
			finalHandler = plugin.Handler.Execute(finalHandler)
		}
	}
	return finalHandler
}

func (pm *PluginManager) GetPlugins() []*Plugin {
	return pm.plugins
}
