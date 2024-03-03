package plugins

import (
	"net/http"
)

type PluginExecutor interface {
	Execute(next http.Handler) http.Handler
}

type Plugin struct {
	Name     string
	Priority uint
	Handler  PluginExecutor
}
