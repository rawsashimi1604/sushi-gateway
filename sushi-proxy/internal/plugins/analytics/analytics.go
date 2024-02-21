package analytics

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type AnalyticsPlugin struct{}

var Plugin = NewAnalyticsPlugin()

func (plugin AnalyticsPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing analytics function...")
		slog.Info("Performing some analytics:: 1 + 1 = 2")
		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}

func NewAnalyticsPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "analytics",
		Priority: 12,
		Handler:  AnalyticsPlugin{},
	}
}
