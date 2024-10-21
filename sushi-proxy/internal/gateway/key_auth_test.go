package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleKeyAuthRequest(t *testing.T, key string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the apiKey Header
	req.Header.Set("apiKey", key)

	// Set the cors plugin data.
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"key": "validApiKey",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the basic auth plugin
	plugin := NewKeyAuthPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestKeyAuthPlugin(t *testing.T) {
	rr := handleKeyAuthRequest(t, "validApiKey")
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestKeyAuthWrongKey(t *testing.T) {
	rr := handleKeyAuthRequest(t, "invalidKey")
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusUnauthorized)
	}
}
