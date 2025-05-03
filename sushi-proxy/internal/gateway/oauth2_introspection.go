package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
)

type OAuth2IntrospectionPlugin struct {
	config map[string]interface{}
}

func NewOAuth2IntrospectionPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_OAUTH2_INTROSPECTION,
		Priority: 1700,
		Phase:    AccessPhase,
		Handler: OAuth2IntrospectionPlugin{
			config: config,
		},
		Validator: OAuth2IntrospectionPlugin{
			config: config,
		},
	}
}

// TODO: add validation
func (plugin OAuth2IntrospectionPlugin) Validate() error {

	// Introspection URL is required and must be a valid URL
	if plugin.config["introspection_url"] == nil {
		return fmt.Errorf("introspection_url is required")
	}

	introspectionURL, ok := plugin.config["introspection_url"].(string)
	if !ok {
		return fmt.Errorf("introspection_url must be a string")
	}

	if _, err := url.Parse(introspectionURL); err != nil {
		return fmt.Errorf("introspection_url must be a valid URL")
	}

	if plugin.config["client_id"] == nil {
		return fmt.Errorf("client_id is required")
	}

	if plugin.config["client_secret"] == nil {
		return fmt.Errorf("client_secret is required")
	}

	return nil
}

func (plugin OAuth2IntrospectionPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO: implement logic.
		slog.Info("Hello world from oauth 2 introspection plugin")
		next.ServeHTTP(w, r)
	})
}
