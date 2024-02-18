package basic_auth

import (
	"encoding/base64"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/internal/cache"
	"github.com/rawsashimi1604/sushi-gateway/internal/plugins"
	"log/slog"
	"net/http"
	"strings"
)

type BasicAuthPlugin struct{}

var Plugin = NewBasicAuthPlugin()

var BasicAuthCache = cache.New(5, 100)

func (plugin BasicAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing basic auth function...")

		username, password, err := verifyAndParseAuthHeader(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		slog.Info(fmt.Sprintf("username: %s, password: %s", username, password))

		// Add WWW Authenticate Header
		// Compare password...

		next.ServeHTTP(w, r)
	})
}

func NewBasicAuthPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "basic_auth",
		Priority: 15,
		Handler:  BasicAuthPlugin{},
	}
}

func verifyAndParseAuthHeader(r *http.Request) (username string, password string, error *plugins.PluginError) {
	authHeader := r.Header.Get("Authorization")

	bits := strings.Split(authHeader, " ")

	// valid format : Basic user:pass(base64 encoded)
	isValidAuthFormat := authHeader != "" && len(bits) == 2
	if !isValidAuthFormat {
		return "", "", plugins.NewPluginError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid basic auth format passed in.")
	}

	// Decode the base64 string
	decoded, err := base64.StdEncoding.DecodeString(bits[1])
	if err != nil {
		return "", "", plugins.NewPluginError(http.StatusUnauthorized,
			"DECODE_TOKEN_ERROR", "Unable to decode base64 token.")
	}

	tokenVals := strings.Split(string(decoded), ":")
	if len(tokenVals) != 2 {
		return "", "", plugins.NewPluginError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid basic auth format passed in.")
	}

	return tokenVals[0], tokenVals[1], nil
}
