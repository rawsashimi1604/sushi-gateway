package rate_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/plugins"
	"log/slog"
	"net/http"
)

var RateLimitPlugin = plugins.Plugin{
	Name:     "rate_limit",
	Priority: 10,
	Handler:  execute(),
}

func execute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")
	}
}
