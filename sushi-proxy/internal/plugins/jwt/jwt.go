package jwt

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/cache"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type JwtPlugin struct{}

var Plugin = NewJwtPlugin()

// JwtCache TODO: add caching mechanisms, persist between page views, per realm
var JwtCache = cache.New(5, 100)

func (plugin JwtPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing jwt auth function...")
		next.ServeHTTP(w, r)
	})
}

func NewJwtPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "jwt",
		Priority: 15,
		Handler:  JwtPlugin{},
	}
}
