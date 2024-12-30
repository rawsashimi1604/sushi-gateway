package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const mockJwtSecret = "123secret456"
const mockJwtIssuer = "someIssuerKey"

func TestJwtValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "valid config",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    mockJwtIssuer,
				"secret": mockJwtSecret,
			},
			expectError: false,
		},
		{
			name: "missing alg",
			config: map[string]interface{}{
				"iss":    mockJwtIssuer,
				"secret": mockJwtSecret,
			},
			expectError: true,
		},
		{
			name: "empty alg",
			config: map[string]interface{}{
				"alg":    "",
				"iss":    mockJwtIssuer,
				"secret": mockJwtSecret,
			},
			expectError: true,
		},
		{
			name: "invalid alg",
			config: map[string]interface{}{
				"alg":    "RS256",
				"iss":    mockJwtIssuer,
				"secret": mockJwtSecret,
			},
			expectError: true,
		},
		{
			name: "missing iss",
			config: map[string]interface{}{
				"alg":    "HS256",
				"secret": mockJwtSecret,
			},
			expectError: true,
		},
		{
			name: "empty iss",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    "",
				"secret": mockJwtSecret,
			},
			expectError: true,
		},
		{
			name: "missing secret",
			config: map[string]interface{}{
				"alg": "HS256",
				"iss": mockJwtIssuer,
			},
			expectError: true,
		},
		{
			name: "empty secret",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    mockJwtIssuer,
				"secret": "",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pluginConfig, err := CreatePluginConfigInput(tt.config)
			if err != nil {
				t.Fatal(err)
			}

			plugin := NewJwtPlugin(pluginConfig)
			err = plugin.Validator.Validate()

			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func handleHS256JwtRequest(t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Generate a JWT token using mockJwtSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": mockJwtIssuer,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	signedToken, err := token.SignedString([]byte(mockJwtSecret))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+signedToken)

	// Setup the jwt plugin
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"alg":    "HS256",
		"iss":    mockJwtIssuer,
		"secret": mockJwtSecret,
	})
	if err != nil {
		t.Fatal(err)
	}

	plugin := NewJwtPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)
	return rr
}

func TestJwtHS256AuthSuccess(t *testing.T) {
	rr := handleHS256JwtRequest(t)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Header().Get("Authorization") != "" {
		t.Errorf("Authorization header should be stripped when JWT auth is passing")
	}
}
