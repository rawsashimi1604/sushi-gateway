package plugins

import "net/http"

type PluginExecutor interface {
	Execute(req *http.Request)
}

type Plugin struct {
	Name     string
	Priority uint
	Handler  PluginExecutor
}
