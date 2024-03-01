package key_auth

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
)

type KeyAuthPlugin struct{}

var Plugin = NewKeyAuthPlugin()

func NewKeyAuthPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "key_auth",
		Priority: 10,
		Handler:  KeyAuthPlugin{},
	}
}

func (plugin KeyAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing KeyAuth auth function...")

		apiKey, err := extractAPIKey(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		err = validateAPIKey(r, apiKey)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func validateAPIKey(r *http.Request, apiKey string) *errors.HttpError {
	service, _, err := util.GetServiceAndRouteFromRequest(r)
	if err != nil {
		return err
	}

	for _, cred := range service.Credentials {
		if cred.Plugin == "key_auth" {
			// Get from cred map
			key, ok := cred.Data["key"].(string)
			if !ok {
				return errors.NewHttpError(http.StatusUnauthorized,
					"INVALID_CREDENTIALS", "Invalid credentials.")
			}

			if key == apiKey {
				return nil
			}
		}
	}
	return errors.NewHttpError(http.StatusUnauthorized,
		"INVALID_CREDENTIALS", "Invalid credentials.")

}

func extractAPIKey(r *http.Request) (string, *errors.HttpError) {
	// From query parameter
	apiKey := r.URL.Query().Get("apiKey")

	if apiKey != "" {
		return apiKey, nil
	}

	// From header
	apiKey = r.Header.Get("apiKey")

	if apiKey != "" {
		return apiKey, nil
	}

	return "", errors.NewHttpError(http.StatusUnauthorized,
		"MISSING_API_KEY", "API key is missing.")
}
