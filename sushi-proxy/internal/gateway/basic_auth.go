package gateway

import (
	"encoding/base64"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"log/slog"
	"net/http"
	"strings"
)

type BasicAuthPlugin struct {
	config map[string]interface{}
}

func NewBasicAuthPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_BASIC_AUTH,
		Priority: 15,
		Handler: BasicAuthPlugin{
			config: config,
		},
	}
}

func (plugin BasicAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing basic auth function...")

		username, password, err := verifyAndParseAuthHeader(r)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		err = plugin.authorize(username, password)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeWWWAuthenticateHeader(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate",
		fmt.Sprintf("Basic realm=\"%s\", charset=%s", "Access to sushi gateway", constant.UTF_8))
}

func verifyAndParseAuthHeader(r *http.Request) (username string, password string, error *model.HttpError) {
	authHeader := r.Header.Get("Authorization")
	bits := strings.Split(authHeader, " ")

	// valid format : Basic user:pass(base64 encoded)
	isValidAuthFormat := authHeader != "" && len(bits) == 2
	if !isValidAuthFormat {
		slog.Info("Invalid basic auth format passed in.")
		return "", "", model.NewHttpError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid auth format passed in.")
	}

	decoded, err := base64.StdEncoding.DecodeString(bits[1])
	if err != nil {
		slog.Info("Unable to decode base64 token.")
		return "", "", model.NewHttpError(http.StatusUnauthorized,
			"DECODE_TOKEN_ERROR", "Unable to decode base64 token.")
	}

	tokenVals := strings.Split(string(decoded), ":")
	if len(tokenVals) != 2 {
		slog.Info("Invalid basic auth format passed in.")
		return "", "", model.NewHttpError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid basic auth format passed in.")
	}

	return tokenVals[0], tokenVals[1], nil
}

// Get from configurations
func (plugin BasicAuthPlugin) authorize(username string, password string) *model.HttpError {

	invalidCredsErr := model.NewHttpError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "invalid credentials, please try again.")

	config := plugin.config

	usernameFromConfig, okUser := config["username"].(string)
	passwordFromConfig, okPass := config["password"].(string)

	// Invalid configuration
	// TODO: handle this better, do validation in the gateway file
	if !okUser || !okPass {
		return invalidCredsErr
	}

	if username == usernameFromConfig && password == passwordFromConfig {
		return nil
	}
	return invalidCredsErr
}
