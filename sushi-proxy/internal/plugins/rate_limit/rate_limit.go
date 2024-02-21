package rate_limit

import (
	plugins2 "github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type RateLimitPlugin struct{}

var Plugin = NewRateLimitPlugin()

func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")
		//slog.Info("Too many requests!!! Rate limit reached!!!")
		//pluginErr := plugins2.NewPluginError(http.StatusBadRequest, "RATE_LIMIT_MOCK_ERROR", "rate limit reached!")
		//pluginErr.WriteJSONResponse(w)

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}

func NewRateLimitPlugin() *plugins2.Plugin {
	return &plugins2.Plugin{
		Name:     "rate_limit",
		Priority: 10,
		Handler:  RateLimitPlugin{},
	}
}
