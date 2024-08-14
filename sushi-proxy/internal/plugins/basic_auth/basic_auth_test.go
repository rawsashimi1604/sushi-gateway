package basic_auth

import (
	"encoding/base64"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBasicAuthSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request with basic auth header
	username := "mockUser"
	password := "mockPassword"
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))

	// Set the basic auth plugin data.
	config, err := util.CreatePluginConfigJsonInput(map[string]interface{}{
		"data": map[string]interface{}{
			"username": username,
			"password": password,
		},
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

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Header().Get("WWW-Authenticate") != "" {
		t.Errorf("Should not get WWW-Authenticate header when basic auth is passing")
	}
}

func TestBasicAuthFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request with failing basic auth header
	username := "mockUser"
	password := "mockPassword"
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("wrongUser"+":wrongPassword")))

	// Set the basic auth plugin data.
	config, err := util.CreatePluginConfigJsonInput(map[string]interface{}{
		"data": map[string]interface{}{
			"username": username,
			"password": password,
		},
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

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Should return 401 Unauthorized when basic auth is failing")
	}
	if rr.Header().Get("WWW-Authenticate") == "" {
		t.Errorf("Should get WWW-Authenticate header when basic auth is failing")
	}
}
