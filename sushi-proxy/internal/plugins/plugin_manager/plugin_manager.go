package plugin_manager

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/acl"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/basic_auth"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/bot_protection"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/cors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/http_log"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/jwt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/key_auth"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/mtls"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/rate_limit"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/request_size_limit"
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
	service, route, err := util.GetServiceAndRouteFromRequest(&config.GlobalProxyConfig, req)
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

	enabled, enabledOk := pc["enabled"].(bool)
	if !enabledOk {
		return errors.NewHttpError(http.StatusInternalServerError, "PLUGIN_CONFIG_ERROR",
			"Plugin enabled flag not found")
	}

	// Skip as not enabled.
	if !enabled {
		return nil
	}

	switch name {
	case constant.PLUGIN_BASIC_AUTH:
		pm.RegisterPlugin(basic_auth.NewBasicAuthPlugin(pc))
	case constant.PLUGIN_ACL:
		pm.RegisterPlugin(acl.NewAclPlugin(pc))
	case constant.PLUGIN_BOT_PROTECTION:
		pm.RegisterPlugin(bot_protection.NewBotProtectionPlugin(pc))
	case constant.PLUGIN_KEY_AUTH:
		pm.RegisterPlugin(key_auth.NewKeyAuthPlugin(pc))
	case constant.PLUGIN_RATE_LIMIT:
		pm.RegisterPlugin(rate_limit.NewRateLimitPlugin(pc))
	case constant.PLUGIN_REQUEST_SIZE_LIMIT:
		pm.RegisterPlugin(request_size_limit.NewRequestSizeLimitPlugin(pc))
	case constant.PLUGIN_JWT:
		pm.RegisterPlugin(jwt.NewJwtPlugin(pc))
	case constant.PLUGIN_MTLS:
		pm.RegisterPlugin(mtls.NewMtlsPlugin())
	case constant.PLUGIN_HTTP_LOG:
		pm.RegisterPlugin(http_log.NewHttpLogPlugin(pc))
	case constant.PLUGIN_CORS:
		pm.RegisterPlugin(cors.NewCorsPlugin(pc))
	}
	return nil
}

func (pm *PluginManager) RegisterPlugin(plugin *plugins.Plugin) {
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

func (pm *PluginManager) GetPlugins() []*plugins.Plugin {
	return pm.plugins
}
