package plugins

type PluginManager struct {
	plugins []*Plugin
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		plugins: make([]*Plugin, 0),
	}
}

func (pm *PluginManager) RegisterPlugin(plugin *Plugin) {
	pm.plugins = append(pm.plugins, plugin)
}

func (pm *PluginManager) GetPlugins() []*Plugin {
	return pm.plugins
}
