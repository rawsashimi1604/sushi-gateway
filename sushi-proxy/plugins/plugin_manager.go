package plugins

type PluginManager struct {
	plugins []*Plugin
}

func (pm *PluginManager) registerPlugin(plugin *Plugin) {
	pm.plugins = append(pm.plugins, plugin)
}
