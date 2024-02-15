package plugins

import "net/http"

type Plugin struct {
	name     string
	priority uint
	handler  http.HandlerFunc
}
