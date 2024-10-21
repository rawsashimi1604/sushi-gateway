package gateway

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func handleRequestSizeLimitReq(t *testing.T, body string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	// Set the request size limit plugin data.
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"max_payload_size": 15,
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the basic auth plugin
	plugin := NewRequestSizeLimitPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestRequestSizeLimit(t *testing.T) {
	rr := handleRequestSizeLimitReq(t, "MoreThan15BytesDefinitely")
	if rr.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusRequestEntityTooLarge)
	}
}

func TestRequestSizeLimitBypass(t *testing.T) {
	rr := handleRequestSizeLimitReq(t, "empty")
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
