package basic_auth

import (
	"encoding/base64"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/cache"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"strings"
)

type BasicAuthPlugin struct{}

var Plugin = NewBasicAuthPlugin()

// BasicAuthCache TODO: add caching mechanisms, persist between page views, per realm
var BasicAuthCache = cache.New(5, 100)

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
		err = authorize(username, password, r)
		if err != nil {
			writeWWWAuthenticateHeader(w)
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

func verifyAndParseAuthHeader(r *http.Request) (username string, password string, error *errors.HttpError) {
	authHeader := r.Header.Get("Authorization")
	bits := strings.Split(authHeader, " ")

	// valid format : Basic user:pass(base64 encoded)
	isValidAuthFormat := authHeader != "" && len(bits) == 2
	if !isValidAuthFormat {
		slog.Info("Invalid basic auth format passed in.")
		return "", "", errors.NewHttpError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid auth format passed in.")
	}

	decoded, err := base64.StdEncoding.DecodeString(bits[1])
	if err != nil {
		slog.Info("Unable to decode base64 token.")
		return "", "", errors.NewHttpError(http.StatusUnauthorized,
			"DECODE_TOKEN_ERROR", "Unable to decode base64 token.")
	}

	tokenVals := strings.Split(string(decoded), ":")
	if len(tokenVals) != 2 {
		slog.Info("Invalid basic auth format passed in.")
		return "", "", errors.NewHttpError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid basic auth format passed in.")
	}

	return tokenVals[0], tokenVals[1], nil
}

func authorize(username string, password string, r *http.Request) *errors.HttpError {
	service, _, err := util.GetServiceAndRouteFromRequest(r)
	if err != nil {
		return err
	}

	for _, cred := range service.Credentials {
		if cred.Plugin == "basic_auth" {
			// Assert string type and get from cred map
			usernameFromCred, okUser := cred.Data["username"].(string)
			passwordFromCred, okPass := cred.Data["password"].(string)
			if !okUser || !okPass {
				return errors.NewHttpError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "invalid credentials, please try again.")
			}
			if username == usernameFromCred && password == passwordFromCred {
				return nil
			}
		}
	}
	return errors.NewHttpError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "invalid credentials, please try again.")
}
