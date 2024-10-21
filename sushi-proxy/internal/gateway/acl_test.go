package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWhitelistedIPs(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request from a whitelisted IP
	req.RemoteAddr = "127.0.0.1"

	// Convert the input to how we would expect it to be in the gateway file
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"whitelist": []string{"127.0.0.1"},
		"blacklist": []string{},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the plugin
	plugin := NewAclPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestBlacklistedIPs(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request from a blacklisted IP
	req.RemoteAddr = "192.168.1.1"

	config, err := CreatePluginConfigInput(map[string]interface{}{
		"whitelist": []string{},
		"blacklist": []string{"192.168.1.1"},
	})
	if err != nil {
		t.Fatal(err)
	}

	plugin := NewAclPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
	}
}
