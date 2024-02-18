package basic_auth

import (
	plugins2 "github.com/rawsashimi1604/sushi-gateway/internal/plugins"
	"log/slog"
	"net/http"
)

type BasicAuthPlugin struct{}

var Plugin = NewBasicAuthPlugin()

func (plugin BasicAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing basic auth function...")

		// Parse out username and password from Authorization Header (b64)

		// Add WWW Authenticate Header

		// Check if API key valid...

		// Compare password...

		next.ServeHTTP(w, r)
	})
}

func NewBasicAuthPlugin() *plugins2.Plugin {
	return &plugins2.Plugin{
		Name:     "basic_auth",
		Priority: 10,
		Handler:  BasicAuthPlugin{},
	}
}
