package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleCorsRequest(t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set the cors plugin data.
	config, err := CreatePluginConfigJsonInput(map[string]interface{}{
		"data": map[string]interface{}{
			"allow_origins":         []string{"*"},
			"allow_methods":         []string{"GET", "POST"},
			"allow_headers":         []string{"Authorization", "Content-Type"},
			"expose_headers":        []string{"Authorization"},
			"allow_credentials":     true,
			"allow_private_network": false,
			"preflight_continue":    false,
			"max_age":               3600,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the basic auth plugin
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
