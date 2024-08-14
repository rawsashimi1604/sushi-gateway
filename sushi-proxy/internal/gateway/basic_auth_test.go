package gateway

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

const username = "mockUser"
const password = "mockPassword"

func handleRequest(t *testing.T, user string, pass string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request with basic auth header
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(user+":"+pass)))

	// Set the basic auth plugin data.
	config, err := CreatePluginConfigJsonInput(map[string]interface{}{
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

	return rr
}

func TestBasicAuthSuccess(t *testing.T) {

	rr := handleRequest(t, username, password)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	if rr.Header().Get("WWW-Authenticate") != "" {
		t.Errorf("Should not get WWW-Authenticate header when basic auth is passing")
	}
}

func TestBasicAuthFail(t *testing.T) {
	rr := handleRequest(t, "fakeUser", "fakePassword")

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Should return 401 Unauthorized when basic auth is failing")
	}
	if rr.Header().Get("WWW-Authenticate") == "" {
		t.Errorf("Should get WWW-Authenticate header when basic auth is failing")
	}
}
