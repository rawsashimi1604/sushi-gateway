package key_auth

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type KeyAuthPlugin struct {
	config map[string]interface{}
}

func NewKeyAuthPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "key_auth",
		Priority: 10,
		Handler: KeyAuthPlugin{
			config: config,
		},
	}
}

func (plugin KeyAuthPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing key_auth function...")

		apiKey, err := extractAPIKey(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		err = plugin.validateAPIKey(apiKey)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func (plugin KeyAuthPlugin) validateAPIKey(apiKey string) *errors.HttpError {
	data := plugin.config["data"].(map[string]interface{})
	key := data["key"].(string) // Assert to []interface{} first
	if key == apiKey {
		return nil
	} else {
		return errors.NewHttpError(http.StatusUnauthorized,
			"INVALID_CREDENTIALS", "Invalid credentials.")
	}
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
