package gateway

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

func setupMockProxyConfig() {
	GlobalProxyConfig = model.ProxyConfig{
		Global: model.Global{
			Name: "test-gateway",
			Plugins: []model.PluginConfig{
				{
					Id:      "plugin_1",
					Name:    "http_log",
					Enabled: true,
					Config: map[string]interface{}{
						"http_endpoint": "http://localhost:3000/v1/log",
						"method":        "POST",
						"content_type":  "application/json",
					},
				},
			},
		},
		Services: []model.Service{
			{
				Name:                  "test-service",
				Protocol:              "http",
				BasePath:              "/test",
				LoadBalancingStrategy: model.RoundRobin,
				Plugins:               []model.PluginConfig{},

				Upstreams: []model.Upstream{
					{
						Id:   "upstream_id",
						Host: "localhost",
						Port: 8080,
					},
				},
				Routes: []model.Route{
					{
						Name:    "test-route",
						Path:    "/test-path",
						Methods: []string{"GET"},
						Plugins: []model.PluginConfig{},
					},
				},
			},
		},
	}
}

func TestHttpLogValidation(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "valid config",
			config: map[string]interface{}{
				"http_endpoint": "http://localhost:8080/logs",
				"method":        "POST",
				"content_type":  "application/json",
			},
			expectError: false,
		},
		{
			name: "missing http_endpoint",
			config: map[string]interface{}{
				"method":       "POST",
				"content_type": "application/json",
			},
			expectError: true,
		},
		{
			name: "empty http_endpoint",
			config: map[string]interface{}{
				"http_endpoint": "",
				"method":        "POST",
				"content_type":  "application/json",
			},
			expectError: true,
		},
		{
			name: "missing method",
			config: map[string]interface{}{
				"http_endpoint": "http://localhost:8080/logs",
				"content_type":  "application/json",
			},
			expectError: true,
		},
		{
			name: "invalid method",
			config: map[string]interface{}{
				"http_endpoint": "http://localhost:8080/logs",
				"method":        "DELETE",
				"content_type":  "application/json",
			},
			expectError: true,
		},
		{
			name: "missing content_type",
			config: map[string]interface{}{
				"http_endpoint": "http://localhost:8080/logs",
				"method":        "POST",
			},
			expectError: true,
		},
		{
			name: "invalid content_type",
			config: map[string]interface{}{
				"http_endpoint": "http://localhost:8080/logs",
				"method":        "POST",
				"content_type":  "text/plain",
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

			plugin := NewHttpLogPlugin(pluginConfig)
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

func TestHttpLogExecution(t *testing.T) {
	// Setup mock proxy config
	setupMockProxyConfig()

	// Create a test server to receive logs
	logServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Verify request method
		if r.Method != "POST" {
			t.Errorf("expected POST request, got %s", r.Method)
		}

		// Verify log structure
		var logData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&logData); err != nil {
			t.Errorf("failed to decode log data: %v", err)
			return
		}

		slog.Info("Received log", "log", logData)

		// Check required fields
		requiredFields := []string{"service", "route", "request", "client_ip", "started_at"}
		for _, field := range requiredFields {
			if _, ok := logData[field]; !ok {
				t.Errorf("missing required field: %s", field)
			}
		}

		// Verify service data
		service, ok := logData["service"].(map[string]interface{})
		if !ok {
			t.Error("service field is not a map")
			return
		}
		if service["name"] != "test-service" {
			t.Errorf("expected service name test-service, got %v", service["name"])
		}

		// Verify route data
		route, ok := logData["route"].(map[string]interface{})
		if !ok {
			t.Error("route field is not a map")
			return
		}
		if route["path"] != "/test-path" {
			t.Errorf("expected route path /test-path, got %v", route["path"])
		}

		// Verify request data
		request, ok := logData["request"].(map[string]interface{})
		if !ok {
			t.Error("request field is not a map")
			return
		}
		requiredRequestFields := []string{"protocol", "method", "path", "url", "uri", "headers"}
		for _, field := range requiredRequestFields {
			if _, ok := request[field]; !ok {
				t.Errorf("missing required request field: %s", field)
			}
		}

		// Verify response data
		response, ok := logData["response"].(map[string]interface{})
		if !ok {
			t.Error("response field is not a map")
			return
		}
		if response["status"].(float64) != float64(http.StatusOK) {
			t.Errorf("expected status %d, got %v", http.StatusOK, response["status"])
		}
		if response["size"].(float64) != float64(0) {
			t.Errorf("expected size 0, got %v", response["size"])
		}

		// Verify timestamps and latency
		startedAt, ok := logData["started_at"].(float64)
		if !ok {
			t.Error("started_at is not a number")
			return
		}
		endedAt, ok := logData["ended_at"].(float64)
		if !ok {
			t.Error("ended_at is not a number")
			return
		}
		latency, ok := logData["latency"].(string)
		if !ok {
			t.Error("latency is not a string")
			return
		}

		// Verify latency calculation
		expectedLatency := endedAt - startedAt
		if latency != strconv.Itoa(int(expectedLatency))+"ms" {
			t.Errorf("expected latency %v, got %v", expectedLatency, latency)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer logServer.Close()

	// Create plugin config
	config := map[string]interface{}{
		"http_endpoint": logServer.URL,
		"method":        "POST",
		"content_type":  "application/json",
	}

	pluginConfig, err := CreatePluginConfigInput(config)
	if err != nil {
		t.Fatal(err)
	}

	// Create plugin
	plugin := NewHttpLogPlugin(pluginConfig)

	// Create test request with mock context values
	req := httptest.NewRequest("GET", "/test/test-path", nil)
	req.RemoteAddr = "127.0.0.1:1234"

	// Add mock context values
	startTime := time.Now()
	endTime := startTime.Add(100 * time.Millisecond)
	mockHeaders := map[string][]string{"Content-Type": {"application/json"}}

	ctx := context.WithValue(req.Context(), constant.CONTEXT_START_TIME, startTime)
	ctx = context.WithValue(ctx, constant.CONTEXT_END_TIME, endTime)
	ctx = context.WithValue(ctx, constant.CONTEXT_RESPONSE_STATUS, http.StatusOK)
	ctx = context.WithValue(ctx, constant.CONTEXT_RESPONSE_SIZE, 0)
	ctx = context.WithValue(ctx, constant.CONTEXT_RESPONSE_HEADERS, mockHeaders)
	req = req.WithContext(ctx)

	// Execute plugin
	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rr, req)

	// Verify that the main request was successful
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	// Give some time for the log to be sent
	time.Sleep(100 * time.Millisecond)
}
