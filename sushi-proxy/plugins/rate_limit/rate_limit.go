package rate_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/plugins"
	"log/slog"
	"net/http"
)

type RateLimitPlugin struct{}

var Plugin = NewRateLimitPlugin()

func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")

		slog.Info("Too many requests!!! Rate limit reached!!!")
		pluginErr := plugins.NewPluginError(http.StatusBadRequest, "RATE_LIMIT_MOCK_ERROR", "rate limit reached!")
		pluginErr.WriteJSONResponse(w)
		return

		// call the next plugin.
		//next.ServeHTTP(w, r)
	})
}

func NewRateLimitPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "rate_limit",
		Priority: 10,
		Handler:  RateLimitPlugin{},
	}
}
