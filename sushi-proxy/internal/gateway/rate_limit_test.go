package gateway

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleRateLimitRequest(t *testing.T, ip string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Simulate a request with IP addr
	req.RemoteAddr = ip
	req.URL.Path = "/mockService/mockRoute"

	// Set the rate limit plugin data.
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"limit_second": 10,
		"limit_min":    10,
		"limit_hour":   10,
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create a mock proxy config
	mockProxyConfig := createMockProxyConfig(t)

	// Create a new instance of the basic auth plugin
	plugin := NewRateLimitPlugin(config, mockProxyConfig)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestRateLimitPluginTooManyRequests(t *testing.T) {
	var rr *httptest.ResponseRecorder
	// Execute 11 requests since limit is 10
	for i := 0; i <= 11; i++ {
		rr = handleRateLimitRequest(t, "tooManyReqIp")
	}
	if rr.Code != http.StatusTooManyRequests {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusTooManyRequests)
	}
}

func TestRateLimitPluginBypass(t *testing.T) {
	var rr *httptest.ResponseRecorder
	// Execute 5 requests since limit is 10, dont go above the limit
	for i := 0; i <= 5; i++ {
		rr = handleRateLimitRequest(t, "successIp")
	}
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func createMockProxyConfig(t *testing.T) *model.ProxyConfig {

	//rateLimitPlugin, err := CreatePluginConfigInput(map[string]interface{}{
	//	"name":    "rate_limit",
	//	"enabled": true,
	//	"data": map[string]interface{}{
	//		"limit_second": 10,
	//		"limit_min":    10,
	//		"limit_hour":   10,
	//	},
	//})

	rateLimitPlugin := model.PluginConfig{
		Id:      "someId",
		Name:    "rate_limit",
		Enabled: true,
		Config: map[string]interface{}{
			"limit_second": 10,
			"limit_min":    10,
			"limit_hour":   10,
		},
	}
	//if err != nil {
	//	t.Fatal(err)
	//}

	service := model.Service{
		Name:                  "mockService",
		BasePath:              "/mockService",
		Protocol:              "http",
		LoadBalancingStrategy: RoundRobin,
		Upstreams:             make([]model.Upstream, 0),
		Plugins: []model.PluginConfig{
			rateLimitPlugin,
		},
		Routes: []model.Route{
			{
				Name:    "mockRoute",
				Path:    "/mockRoute",
				Methods: []string{"GET"},
				Plugins: make([]model.PluginConfig, 0),
			},
		},
	}

	proxyConfig := &model.ProxyConfig{}
	proxyConfig.Global.Name = "mockProxy"
	proxyConfig.Services = []model.Service{service}

	return proxyConfig
}
