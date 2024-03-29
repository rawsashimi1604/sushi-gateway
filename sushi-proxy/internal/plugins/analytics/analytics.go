package analytics

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type AnalyticsPlugin struct{}

func NewAnalyticsPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_ANALYTICS,
		Priority: 12,
		Handler:  AnalyticsPlugin{},
	}
}

func (plugin AnalyticsPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing analytics function...")

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
