package gateway

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

type KeyAuthPlugin struct {
	config map[string]interface{}
}

func NewKeyAuthPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_KEY_AUTH,
		Priority: 1250,
		Phase:    AccessPhase,
		Handler: KeyAuthPlugin{
			config: config,
		},
		Validator: KeyAuthPlugin{
			config: config,
		},
	}
}

func (plugin KeyAuthPlugin) Validate() error {
	key, ok := plugin.config["key"].(string)
	if !ok || key == "" {
		return fmt.Errorf("key must be a non-empty string")
	}
	return nil
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

		// Strip header
		r.Header.Del("apiKey")

		next.ServeHTTP(w, r)
	})
}
func (plugin KeyAuthPlugin) validateAPIKey(apiKey string) *model.HttpError {
	config := plugin.config
	key := config["key"].(string) // Assert to []interface{} first
	if key == apiKey {
		return nil
	} else {
		return model.NewHttpError(http.StatusUnauthorized,
			"INVALID_CREDENTIALS", "Invalid credentials.")
	}
}

func extractAPIKey(r *http.Request) (string, *model.HttpError) {
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

	return "", model.NewHttpError(http.StatusUnauthorized,
		"MISSING_API_KEY", "API key is missing.")
}
