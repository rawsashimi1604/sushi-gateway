package gateway

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

const username = "mockUser"
const password = "mockPassword"

func handleBasicAuthRequest(t *testing.T, user string, pass string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request with basic auth header
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(user+":"+pass)))

	// Set the basic auth plugin data.
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"username": username,
		"password": password,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the basic auth plugin
	plugin := NewBasicAuthPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestBasicAuthSuccess(t *testing.T) {

	rr := handleBasicAuthRequest(t, username, password)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Header().Get("WWW-Authenticate") != "" {
		t.Errorf("Should not get WWW-Authenticate header when basic auth is passing")
	}

	if rr.Header().Get("Authorization") != "" {
		t.Errorf("Authorization header should be stripped when basic auth is passing")
	}
}

func TestBasicAuthFail(t *testing.T) {
	rr := handleBasicAuthRequest(t, "fakeUser", "fakePassword")

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Should return 401 Unauthorized when basic auth is failing")
	}
	if rr.Header().Get("WWW-Authenticate") == "" {
		t.Errorf("Should get WWW-Authenticate header when basic auth is failing")
	}
}

func TestBasicAuthPluginValidation(t *testing.T) {
	tests := []struct {
		name      string
		config    map[string]interface{}
		expectErr bool
	}{
		{
			name: "Valid configuration",
			config: map[string]interface{}{
				"username": "validUser",
				"password": "validPass",
			},
			expectErr: false,
		},
		{
			name: "Missing username",
			config: map[string]interface{}{
				"password": "validPass",
			},
			expectErr: true,
		},
		{
			name: "Missing password",
			config: map[string]interface{}{
				"username": "validUser",
			},
			expectErr: true,
		},
		{
			name:      "Missing both username and password",
			config:    map[string]interface{}{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin := BasicAuthPlugin{config: tt.config}
			err := plugin.Validate()
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}
