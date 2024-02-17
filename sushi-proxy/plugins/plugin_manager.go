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
	// TODO: sort when adding in place...
	pm.plugins = append(pm.plugins, plugin)
}

func (pm *PluginManager) GetPlugins() []*Plugin {
	return pm.plugins
}
