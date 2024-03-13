package mtls

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type MtlsPlugin struct{}

func NewMtlsPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_MTLS,
		Priority: 12,
		Handler:  MtlsPlugin{},
	}
}

func (plugin MtlsPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing mtls function...")

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
