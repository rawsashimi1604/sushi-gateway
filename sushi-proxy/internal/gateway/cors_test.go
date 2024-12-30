package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "valid config with minimal fields",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
			},
			expectError: false,
		},
		{
			name: "valid config with all fields",
			config: map[string]interface{}{
				"allow_origins":         []interface{}{"*"},
				"allow_methods":         []interface{}{"GET", "POST"},
				"allow_headers":         []interface{}{"Authorization", "Content-Type"},
				"expose_headers":        []interface{}{"Authorization"},
				"allow_credentials":     true,
				"allow_private_network": false,
				"preflight_continue":    false,
				"max_age":               float64(3600),
			},
			expectError: false,
		},
		{
			name: "missing allow_origins",
			config: map[string]interface{}{
				"allow_methods": []interface{}{"GET", "POST"},
			},
			expectError: true,
		},
		{
			name: "empty allow_origins array",
			config: map[string]interface{}{
				"allow_origins": []interface{}{},
			},
			expectError: true,
		},
		{
			name: "invalid allow_origins type",
			config: map[string]interface{}{
				"allow_origins": "*",
			},
			expectError: true,
		},
		{
			name: "invalid allow_origins entry type",
			config: map[string]interface{}{
				"allow_origins": []interface{}{123},
			},
			expectError: true,
		},
		{
			name: "empty allow_methods array",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"allow_methods": []interface{}{},
			},
			expectError: true,
		},
		{
			name: "invalid allow_methods entry",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"allow_methods": []interface{}{"INVALID"},
			},
			expectError: true,
		},
		{
			name: "invalid allow_methods entry type",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"allow_methods": []interface{}{123},
			},
			expectError: true,
		},
		{
			name: "empty allow_headers array",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"allow_headers": []interface{}{},
			},
			expectError: true,
		},
		{
			name: "invalid allow_headers entry type",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"allow_headers": []interface{}{123},
			},
			expectError: true,
		},
		{
			name: "empty expose_headers array",
			config: map[string]interface{}{
				"allow_origins":  []interface{}{"*"},
				"expose_headers": []interface{}{},
			},
			expectError: true,
		},
		{
			name: "invalid expose_headers entry type",
			config: map[string]interface{}{
				"allow_origins":  []interface{}{"*"},
				"expose_headers": []interface{}{123},
			},
			expectError: true,
		},
		{
			name: "invalid allow_credentials type",
			config: map[string]interface{}{
				"allow_origins":     []interface{}{"*"},
				"allow_credentials": "true",
			},
			expectError: true,
		},
		{
			name: "invalid allow_private_network type",
			config: map[string]interface{}{
				"allow_origins":         []interface{}{"*"},
				"allow_private_network": "false",
			},
			expectError: true,
		},
		{
			name: "invalid preflight_continue type",
			config: map[string]interface{}{
				"allow_origins":      []interface{}{"*"},
				"preflight_continue": "false",
			},
			expectError: true,
		},
		{
			name: "invalid max_age type",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"max_age":       "3600",
			},
			expectError: true,
		},
		{
			name: "negative max_age",
			config: map[string]interface{}{
				"allow_origins": []interface{}{"*"},
				"max_age":       float64(-1),
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

			plugin := NewCorsPlugin(pluginConfig)
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

func handleCorsRequest(t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the cors plugin data.
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"allow_origins":         []interface{}{"*"},
		"allow_methods":         []interface{}{"GET", "POST"},
		"allow_headers":         []interface{}{"Authorization", "Content-Type"},
		"expose_headers":        []interface{}{"Authorization"},
		"allow_credentials":     true,
		"allow_private_network": false,
		"preflight_continue":    false,
		"max_age":               float64(3600),
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the cors plugin
	plugin := NewCorsPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestCorsPlugin(t *testing.T) {
	rr := handleCorsRequest(t)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
	if rr.Header().Get("Access-Control-Allow-Origin") != "*" {
		t.Errorf("handler returned wrong Access-Control-Allow-Origin header: got %v want %v", rr.Header().Get("Access-Control-Allow-Origin"), "*")
	}
	if rr.Header().Get("Access-Control-Allow-Methods") != "GET,POST" {
		t.Errorf("handler returned wrong Access-Control-Allow-Methods header: got %v want %v", rr.Header().Get("Access-Control-Allow-Methods"), "GET,POST")
	}
	if rr.Header().Get("Access-Control-Allow-Headers") != "Authorization,Content-Type" {
		t.Errorf("handler returned wrong Access-Control-Allow-Headers header: got %v want %v", rr.Header().Get("Access-Control-Allow-Headers"), "Authorization,Content-Type")
	}
	if rr.Header().Get("Access-Control-Expose-Headers") != "Authorization" {
		t.Errorf("handler returned wrong Access-Control-Expose-Headers header: got %v want %v", rr.Header().Get("Access-Control-Expose-Headers"), "Authorization")
	}
	if rr.Header().Get("Access-Control-Allow-Credentials") != "true" {
		t.Errorf("handler returned wrong Access-Control-Allow-Credentials header: got %v want %v", rr.Header().Get("Access-Control-Allow-Credentials"), "true")
	}
	if rr.Header().Get("Access-Control-Max-Age") != "3600" {
		t.Errorf("handler returned wrong Access-Control-Max-Age header: got %v want %v", rr.Header().Get("Access-Control-Max-Age"), "3600")
	}
}
