package gateway

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

type BasicAuthPlugin struct {
	config map[string]interface{}
}

func NewBasicAuthPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_BASIC_AUTH,
		Priority: 1100,
		Phase:    AccessPhase,
		Handler: BasicAuthPlugin{
			config: config,
		},
		Validator: BasicAuthPlugin{
			config: config,
		},
	}
}

func (plugin BasicAuthPlugin) Validate() error {
	username, okUser := plugin.config["username"].(string)
	password, okPass := plugin.config["password"].(string)

	if !okUser || username == "" {
		return errors.New("username cannot be empty or missing")
	}
	if !okPass || password == "" {
		return errors.New("password cannot be empty or missing")
	}
	return nil
}

func (plugin BasicAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing basic auth function...")

		username, password, err := verifyAndParseAuthHeaderBasic(r)
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

		// Strip Basic auth header
		r.Header.Del("Authorization")

		next.ServeHTTP(w, r)
	})
}

func writeWWWAuthenticateHeader(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate",
		fmt.Sprintf("Basic realm=\"%s\", charset=%s", "Access to sushi gateway", constant.UTF_8))
}

func verifyAndParseAuthHeaderBasic(r *http.Request) (username string, password string, error *model.HttpError) {
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

	usernameFromConfig, _ := config["username"].(string)
	passwordFromConfig, _ := config["password"].(string)

	if username == usernameFromConfig && password == passwordFromConfig {
		return nil
	}
	return invalidCredsErr
}
