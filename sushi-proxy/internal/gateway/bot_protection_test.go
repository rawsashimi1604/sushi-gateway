package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleBotRequest(t *testing.T, agent string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Simulate a request with basic auth header
	req.Header.Set("User-Agent", agent)

	// Set the bot protection plugin data.
	config, err := CreatePluginConfigJsonInput(map[string]interface{}{
		"data": map[string]interface{}{
			"blacklist": []string{"googlebot", "bingbot", "yahoobot"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new instance of the basic auth plugin
	plugin := NewBotProtectionPlugin(config)

	rr := httptest.NewRecorder()
	handler := plugin.Handler.Execute(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(rr, req)

	return rr
}

func TestBotProtectionSuccess(t *testing.T) {
	// Valid Request
	rr := handleBotRequest(t, "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func TestBotProtectionBlacklisted(t *testing.T) {
	// Invalid Request
	rr := handleBotRequest(t, "googlebot")
	if rr.Code != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
