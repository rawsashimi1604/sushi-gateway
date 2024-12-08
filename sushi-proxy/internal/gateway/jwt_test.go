package gateway

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TODO: add test
const mockJwtSecret = "123secret456"
const mockJwtIssuer = "someIssuerKey"

func handleHS256JwtRequest(t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Generate a JWT token using mockJwtSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": mockJwtIssuer,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	signedToken, err := token.SignedString([]byte(mockJwtSecret))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+signedToken)

	// Setup the jwt plugin
	config, err := CreatePluginConfigInput(map[string]interface{}{
		"alg":    "HS256",
		"iss":    mockJwtIssuer,
		"secret": "123secret456",
	})
	if err != nil {
		t.Fatal(err)
	}

	plugin := NewJwtPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)
	return rr
}

func TestJwtAuthSuccess(t *testing.T) {
	rr := handleHS256JwtRequest(t)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Header().Get("Authorization") != "" {
		t.Errorf("Authorization header should be stripped when basic auth is passing")
	}
}
