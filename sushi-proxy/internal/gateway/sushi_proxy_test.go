package gateway

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: complete a simple e2e test for gateway.
func TestHandleProxyRequest(t *testing.T) {

	req, err := http.NewRequest("GET", "/mockService/mockRoute", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewSushiProxy().RouteRequest()
	handler.ServeHTTP(rr, req)
}

func Test_HandleServiceNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/mockService/mockRoute", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := NewSushiProxy().RouteRequest()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", rr.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	slog.Info(rr.Body.String())

	expectedCode := "SERVICE_NOT_FOUND"
	if errorField, ok := response["error"].(map[string]interface{}); ok {
		if message, ok := errorField["Code"].(string); !ok || message != expectedCode {
			t.Errorf("Expected message '%s', got '%v'", expectedCode, message)
		}
	} else {
		t.Errorf("Expected error field in response body")
	}
}

// TODO: add test for route not found
