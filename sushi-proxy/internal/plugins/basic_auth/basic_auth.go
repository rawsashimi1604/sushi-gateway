package basic_auth

import (
	"encoding/base64"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/cache"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
	"strings"
)

type BasicAuthPlugin struct{}

var Plugin = NewBasicAuthPlugin()

var BasicAuthCache = cache.New(5, 100)

// TODO: add credentials in db...
// TODO: add caching mechanisms, persist between page views, per realm

func (plugin BasicAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing basic auth function...")

		username, password, err := verifyAndParseAuthHeader(r)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		slog.Info(fmt.Sprintf("basicAuth:: username: %s, password: %s", username, password))

		// Compare password...
		mockUser := "admin"
		mockPw := "changeme"
		if username != mockUser || password != mockPw {
			slog.Info("invalid credentials")
			writeWWWAuthenticateHeader(w)
			err = plugins.NewPluginError(http.StatusUnauthorized, "INVALID_CREDENTIALS",
				"invalid credentials, please try again.")
			err.WriteJSONResponse(w)
			return
		}

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

func writeWWWAuthenticateHeader(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate",
		fmt.Sprintf("Basic realm=\"%s\", charset=%s", "Access to sushi gateway", constant.UTF_8))
}

func verifyAndParseAuthHeader(r *http.Request) (username string, password string, error *plugins.PluginError) {
	authHeader := r.Header.Get("Authorization")

	bits := strings.Split(authHeader, " ")

	// valid format : Basic user:pass(base64 encoded)
	isValidAuthFormat := authHeader != "" && len(bits) == 2
	if !isValidAuthFormat {
		slog.Info("Invalid basic auth format passed in.")
		return "", "", plugins.NewPluginError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid basic auth format passed in.")
	}

	// Decode the base64 string
	decoded, err := base64.StdEncoding.DecodeString(bits[1])
	if err != nil {
		slog.Info("Unable to decode base64 token.")
		return "", "", plugins.NewPluginError(http.StatusUnauthorized,
			"DECODE_TOKEN_ERROR", "Unable to decode base64 token.")
	}

	tokenVals := strings.Split(string(decoded), ":")
	if len(tokenVals) != 2 {
		slog.Info("Invalid basic auth format passed in.")
		return "", "", plugins.NewPluginError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid basic auth format passed in.")
	}

	return tokenVals[0], tokenVals[1], nil
}
