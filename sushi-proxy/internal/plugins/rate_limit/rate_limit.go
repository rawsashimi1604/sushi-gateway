package rate_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type RateLimitPlugin struct {
	config map[string]interface{}
}

func NewRateLimitPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_RATE_LIMIT,
		Priority: 10,
		Handler: RateLimitPlugin{
			config: config,
		},
	}
}

func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
