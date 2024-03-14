package http_log

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type HttpLogPlugin struct {
	config map[string]interface{}
}

func NewHttpLogPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_HTTP_LOG,
		Priority: 12,
		Handler: HttpLogPlugin{
			config: config,
		},
	}
}

func (plugin HttpLogPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing httplog function...")
		next.ServeHTTP(w, r)
	})
}
