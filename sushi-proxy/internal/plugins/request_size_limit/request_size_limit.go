package request_size_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type RequestSizeLimitPlugin struct{}

var Plugin = NewRequestSizeLimitPlugin()

func NewRequestSizeLimitPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "request_size_limit",
		Priority: 10,
		Handler:  RequestSizeLimitPlugin{},
	}
}

func (plugin RequestSizeLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing request_size_limit function...")

		next.ServeHTTP(w, r)
	})
}
