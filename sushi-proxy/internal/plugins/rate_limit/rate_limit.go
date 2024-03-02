package rate_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type RateLimitPlugin struct{}

var Plugin = NewRateLimitPlugin()

func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}

func NewRateLimitPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "rate_limit",
		Priority: 10,
		Handler:  RateLimitPlugin{},
	}
}
