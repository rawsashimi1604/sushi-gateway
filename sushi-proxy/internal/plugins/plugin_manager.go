package plugins

import (
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

func NewPluginManagerFromConfig(req *http.Request) *PluginManager {
	// Load the plugin configuration from the config file
	// Based on the plugins loaded from http request

	// TODO: add this functionality
	return &PluginManager{
		plugins: make([]*Plugin, 0),
	}
}

func (pm *PluginManager) RegisterPlugin(plugin *Plugin) {
	// TODO: probably add error handling, if plugin already exists, throw error...
	pm.plugins = append(pm.plugins, plugin)

	// Sort the plugins by priority
	// Higher priority executes first
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

func (pm *PluginManager) GetPlugins() []*Plugin {
	return pm.plugins
}
