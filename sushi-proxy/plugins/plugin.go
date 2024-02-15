package plugins

import "net/http"

type Plugin struct {
	Name     string
	Priority uint
	Handler  http.HandlerFunc
}
