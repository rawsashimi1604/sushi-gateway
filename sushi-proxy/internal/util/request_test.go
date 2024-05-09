package util

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"net/http"
	"testing"
)

func TestGetServiceAndRouteFromRequest(t *testing.T) {

	// Load the proxy config
	config.GlobalAppConfig = config.LoadGlobalConfig()
	config.LoadProxyConfig(config.GlobalAppConfig.ConfigFilePath)
	proxyConfig := config.GlobalProxyConfig

	req, err := http.NewRequest("GET", "/sushi/restaurant", nil)
	if err != nil {
		t.Fatal(err)
	}

	service, route, err := GetServiceAndRouteFromRequest(&proxyConfig, req)
	if err != nil {
		t.Fatal(err)
	}
	if service.Name != "sushi" {
		t.Errorf("service name is wrong: got %v want %v", service.Name, "sushi")
	}
	if route.Path != "/restaurant" {
		t.Errorf("route path is wrong: got %v want %v", route.Path, "/restaurant")
	}
}
