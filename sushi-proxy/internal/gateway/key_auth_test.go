package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const mockValidKey = "test-api-key-123"

func TestKeyAuthValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "valid config",
			config: map[string]interface{}{
				"key": mockValidKey,
			},
			expectError: false,
		},
		{
			name:        "missing key",
			config:      map[string]interface{}{},
			expectError: true,
		},
		{
			name: "empty key",
			config: map[string]interface{}{
				"key": "",
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

			plugin := NewKeyAuthPlugin(pluginConfig)
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

func handleKeyAuthRequest(t *testing.T, withValidKey bool) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Setup the key auth plugin
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"key": mockValidKey,
	})
	if err != nil {
		t.Fatal(err)
	}

	plugin := NewKeyAuthPlugin(config)

	// Set the API key if testing valid key case
	if withValidKey {
		req.Header.Set("apiKey", mockValidKey)
	}

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)
	return rr
}

func TestKeyAuthSuccess(t *testing.T) {
	rr := handleKeyAuthRequest(t, true)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if apiKeyHeader := rr.Header().Get("apiKey"); apiKeyHeader != "" {
		t.Errorf("apiKey header not stripped: got %v want %v", apiKeyHeader, "")
	}

}

func TestKeyAuthFailure(t *testing.T) {
	rr := handleKeyAuthRequest(t, false)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}
