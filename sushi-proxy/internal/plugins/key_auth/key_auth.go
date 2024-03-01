package key_auth

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
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

		if !validateAPIKey(apiKey) {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func validateAPIKey(apiKey string) bool {

	return apiKey == "valid-api-key"
}

func extractAPIKey(r *http.Request) (string, *errors.HttpError) {
	// From query parameter
	apiKey := r.URL.Query().Get("api_key")

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
