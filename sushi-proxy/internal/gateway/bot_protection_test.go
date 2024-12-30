package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleBotRequest(t *testing.T, agent string, config map[string]interface{}) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the User-Agent header
	req.Header.Set("User-Agent", agent)

	// Convert the input to how we would expect it to be in the gateway file
	pluginConfig, err := CreatePluginConfigInput(config)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the plugin
	plugin := NewBotProtectionPlugin(pluginConfig)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestBotProtectionAllowNormalBrowser(t *testing.T) {
	config := map[string]interface{}{
		"blacklist": []string{"googlebot", "bingbot", "yahoobot"},
	}
	rr := handleBotRequest(t, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)", config)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestBotProtectionBlockBot(t *testing.T) {
	config := map[string]interface{}{
		"blacklist": []string{"googlebot", "bingbot", "yahoobot"},
	}
	rr := handleBotRequest(t, "googlebot", config)
	if rr.Code != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusForbidden)
	}
}

func TestBotProtectionAllowEmptyUserAgent(t *testing.T) {
	config := map[string]interface{}{
		"blacklist": []string{"googlebot", "bingbot", "yahoobot"},
	}
	rr := handleBotRequest(t, "", config)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestBotProtectionValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "valid config",
			config: map[string]interface{}{
				"blacklist": []string{"googlebot", "bingbot"},
			},
			expectError: false,
		},
		{
			name:        "missing blacklist",
			config:      map[string]interface{}{},
			expectError: true,
		},
		{
			name: "empty blacklist",
			config: map[string]interface{}{
				"blacklist": []string{},
			},
			expectError: true,
		},
		{
			name: "invalid blacklist type",
			config: map[string]interface{}{
				"blacklist": "not an array",
			},
			expectError: true,
		},
		{
			name: "invalid blacklist entry type",
			config: map[string]interface{}{
				"blacklist": []interface{}{123},
			},
			expectError: true,
		},
		{
			name: "empty string in blacklist",
			config: map[string]interface{}{
				"blacklist": []string{"googlebot", ""},
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

			plugin := NewBotProtectionPlugin(pluginConfig)
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
