package rate_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/plugins"
	"log/slog"
	"net/http"
)

type RateLimitPlugin struct{}

var Plugin = NewRateLimitPlugin()

func (plugin RateLimitPlugin) Execute(req *http.Request) {
	slog.Info("Executing rate limit function...")
	// Add your rate limit logic here
}

func NewRateLimitPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "rate_limit",
		Priority: 10,
		Handler:  RateLimitPlugin{},
	}
}
