package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOAuth2IntrospectionPlugin_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  map[string]interface{}
		wantErr bool
	}{
		{
			name: "Valid configuration",
			config: map[string]interface{}{
				"introspection_url": "http://localhost:8090/realms/sushi-realm/protocol/openid-connect/token/introspect",
				"client_id":         "myclient",
				"client_secret":     "keyInYourSecretHere",
			},
			wantErr: false,
		},
		{
			name: "Missing introspection_url",
			config: map[string]interface{}{
				"client_id":     "myclient",
				"client_secret": "keyInYourSecretHere",
			},
			wantErr: true,
		},
		{
			name: "Invalid introspection_url",
			config: map[string]interface{}{
				"introspection_url": "invalid-url",
				"client_id":         "myclient",
				"client_secret":     "keyInYourSecretHere",
			},
			wantErr: true,
		},
		{
			name: "Missing client_id",
			config: map[string]interface{}{
				"introspection_url": "http://localhost:8090/realms/sushi-realm/protocol/openid-connect/token/introspect",
				"client_secret":     "keyInYourSecretHere",
			},
			wantErr: true,
		},
		{
			name: "Missing client_secret",
			config: map[string]interface{}{
				"introspection_url": "http://localhost:8090/realms/sushi-realm/protocol/openid-connect/token/introspect",
				"client_id":         "myclient",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin := OAuth2IntrospectionPlugin{
				config: tt.config,
			}
			err := plugin.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestOAuth2IntrospectionPlugin_Execute(t *testing.T) {
	// Create a mock server to simulate the introspection endpoint
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "myclient" || password != "keyInYourSecretHere" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := r.FormValue("token")
		if token == "valid_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"active": true}`))
		} else if token == "inactive_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"active": false}`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer mockServer.Close()

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{
			name:           "Valid token",
			authHeader:     "Bearer valid_token",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Inactive token",
			authHeader:     "Bearer inactive_token",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid token",
			authHeader:     "Bearer invalid_token",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Missing token",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Malformed auth header",
			authHeader:     "InvalidFormat",
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the plugin with mock server URL
			plugin := OAuth2IntrospectionPlugin{
				config: map[string]interface{}{
					"introspection_url": mockServer.URL,
					"client_id":         "myclient",
					"client_secret":     "keyInYourSecretHere",
				},
			}

			// Create a test handler that will be called if the plugin allows the request
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// Create a test request
			req := httptest.NewRequest("GET", "http://example.com", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			// Create a response recorder
			recorder := httptest.NewRecorder()

			// Execute the plugin
			handler := plugin.Execute(nextHandler)
			handler.ServeHTTP(recorder, req)

			// Check the response
			assert.Equal(t, tt.expectedStatus, recorder.Code)
		})
	}
}

func TestIntrospectToken(t *testing.T) {
	// Create a mock server to simulate the introspection endpoint
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "myclient" || password != "keyInYourSecretHere" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := r.FormValue("token")
		if token == "valid_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"active": true}`))
		} else if token == "inactive_token" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"active": false}`))
		} else if token == "malformed_response" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{malformed json`))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}))
	defer mockServer.Close()

	tests := []struct {
		name         string
		token        string
		wantErr      bool
		errorCode    string
		errorMessage string
	}{
		{
			name:    "Valid token",
			token:   "valid_token",
			wantErr: false,
		},
		{
			name:         "Inactive token",
			token:        "inactive_token",
			wantErr:      true,
			errorCode:    "TOKEN_EXPIRED",
			errorMessage: "Token has expired or is invalid",
		},
		{
			name:         "Invalid token",
			token:        "invalid_token",
			wantErr:      true,
			errorCode:    "TOKEN_INVALID",
			errorMessage: "Token validation failed",
		},
		{
			name:         "Malformed response",
			token:        "malformed_response",
			wantErr:      true,
			errorCode:    "INTROSPECTION_ERROR",
			errorMessage: "Failed to decode introspection response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := introspectToken(tt.token, mockServer.URL, "myclient", "keyInYourSecretHere")

			if tt.wantErr {
				assert.NotNil(t, err)
				httpErr := err
				assert.Equal(t, tt.errorCode, httpErr.Code)
				assert.Equal(t, tt.errorMessage, httpErr.Message)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
