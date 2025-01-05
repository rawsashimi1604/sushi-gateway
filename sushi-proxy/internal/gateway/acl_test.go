package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleAclRequest(t *testing.T, clientIP string, config map[string]interface{}) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the client IP
	req.RemoteAddr = clientIP

	// Convert the input to how we would expect it to be in the gateway file
	pluginConfig, err := CreatePluginConfigInput(config)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the plugin
	plugin := NewAclPlugin(pluginConfig)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rr, req)
	return rr
}

func TestWhitelistSuccess(t *testing.T) {
	config := map[string]interface{}{
		"whitelist": []string{"127.0.0.1"},
	}
	rr := handleAclRequest(t, "127.0.0.1:8080", config)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestWhitelistBlocked(t *testing.T) {
	config := map[string]interface{}{
		"whitelist": []string{"127.0.0.1"},
	}
	rr := handleAclRequest(t, "192.168.1.1:8080", config)
	if rr.Code != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusForbidden)
	}
}

func TestBlacklistBlocked(t *testing.T) {
	config := map[string]interface{}{
		"blacklist": []string{"192.168.1.1"},
	}
	rr := handleAclRequest(t, "192.168.1.1:8080", config)
	if rr.Code != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusForbidden)
	}
}

func TestBlacklistSuccess(t *testing.T) {
	config := map[string]interface{}{
		"blacklist": []string{"192.168.1.1"},
	}
	rr := handleAclRequest(t, "127.0.0.1:8080", config)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestAclValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "valid whitelist",
			config: map[string]interface{}{
				"whitelist": []string{"127.0.0.1"},
			},
			expectError: false,
		},
		{
			name: "valid blacklist",
			config: map[string]interface{}{
				"blacklist": []string{"192.168.1.1"},
			},
			expectError: false,
		},
		{
			name: "both lists present",
			config: map[string]interface{}{
				"whitelist": []string{"127.0.0.1"},
				"blacklist": []string{"192.168.1.1"},
			},
			expectError: true,
		},
		{
			name:        "no lists present",
			config:      map[string]interface{}{},
			expectError: true,
		},
		{
			name: "invalid whitelist type",
			config: map[string]interface{}{
				"whitelist": 123,
			},
			expectError: true,
		},
		{
			name: "invalid whitelist entry type",
			config: map[string]interface{}{
				"whitelist": []interface{}{123},
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

			plugin := NewAclPlugin(pluginConfig)
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
