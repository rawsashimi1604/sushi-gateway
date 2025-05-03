package gateway

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
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

// TODO: add different token types like JWT and Opaque for the oauth access token
// TODO: add caching support to reduce api calls to authorization server (ttl)
// TODO: add support to strip the the token from the request after succeeding
func (plugin OAuth2IntrospectionPlugin) Validate() error {

	// Introspection URL is required and must be a valid URL
	if plugin.config["introspection_url"] == nil {
		return fmt.Errorf("introspection_url is required")
	}

	introspectionURL, ok := plugin.config["introspection_url"].(string)
	if !ok {
		return fmt.Errorf("introspection_url must be a string")
	}

	parsedURL, err := url.Parse(introspectionURL)
	// checks for valid url and scheme.
	if err != nil || !parsedURL.IsAbs() || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return fmt.Errorf("introspection_url must be a valid absolute URL starting with http:// or https://")
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

		slog.Info("Executing oauth 2 introspection function...")

		// TODO: Refactor it out to a service or something.
		// Check if JWT token is valid
		tokenString, err := verifyAndParseAuthHeaderJwt(r)
		if err != nil {
			writeWWWAuthenticateHeaderJwt(w)
			err.WriteJSONResponse(w)
			return
		}

		// Introspect the token
		introspectionUrl := plugin.config["introspection_url"].(string)
		clientId := plugin.config["client_id"].(string)
		clientSecret := plugin.config["client_secret"].(string)

		err = introspectToken(tokenString, introspectionUrl, clientId, clientSecret)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		// Strip Authorization header
		r.Header.Del("Authorization")

		next.ServeHTTP(w, r)
	})
}

func introspectToken(tokenString string, introspectionURL string, clientID string, clientSecret string) *model.HttpError {
	// Hit the introspection endpoint
	// Create the request body with the token
	data := url.Values{}
	data.Set("token", tokenString)

	// Create the request to the introspection endpoint
	req, err := http.NewRequest("POST", introspectionURL, strings.NewReader(data.Encode()))
	if err != nil {
		slog.Error("Failed to create introspection request", "error", err)
		return model.NewHttpError(http.StatusInternalServerError,
			"INTROSPECTION_ERROR", "Failed to create introspection request")
	}

	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Introspection request failed", "error", err)
		return model.NewHttpError(http.StatusBadGateway,
			"INTROSPECTION_ERROR", "Failed to make introspection request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("Introspection request returned non-200 status", "status", resp.StatusCode)
		return model.NewHttpError(http.StatusUnauthorized,
			"TOKEN_INVALID", "Token validation failed")
	}

	// Parse response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		slog.Error("Failed to decode introspection response", "error", err)
		return model.NewHttpError(http.StatusInternalServerError,
			"INTROSPECTION_ERROR", "Failed to decode introspection response")
	}

	slog.Info("Introspection response", "result", result)

	// Check if token is active
	active, ok := result["active"].(bool)
	if !ok || !active {
		slog.Error("Token is not active or expired")
		return model.NewHttpError(http.StatusUnauthorized,
			"TOKEN_EXPIRED", "Token has expired or is invalid")
	}

	return nil
}
