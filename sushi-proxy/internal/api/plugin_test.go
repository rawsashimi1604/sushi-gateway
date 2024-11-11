package api

import (
	"bytes"
	"encoding/json"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPluginController_AddPlugin(t *testing.T) {
	IntegrationTestGuard(t)

	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	pluginRepo := db.NewPluginRepository(database)
	pluginController := NewPluginController(pluginRepo)

	// Create the PluginDTO for testing
	newPluginDTO := PluginDTO{
		Scope: "route",
		Name:  "get-sushi-restaurants",
		Plugin: model.PluginConfig{
			Name: "basic_auth",
			Config: map[string]interface{}{
				"password": "changeme",
				"username": "admin",
			},
			Enabled: true,
		},
	}

	// Marshal the RouteDTO to JSON
	jsonPayload, _ := json.Marshal(newPluginDTO)
	slog.Info(string(jsonPayload))

	// Create a POST request for adding the route
	req, err := http.NewRequest("POST", "/plugin", bytes.NewBuffer(jsonPayload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := pluginController.AddPlugin()
	handler.ServeHTTP(rr, req)
	slog.Info(rr.Body.String())
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Plugin created successfully")
}

func TestPluginController_UpdatePlugin(t *testing.T) {
	IntegrationTestGuard(t)
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	pluginRepo := db.NewPluginRepository(database)
	pluginController := NewPluginController(pluginRepo)

	// Create the PluginDTO for testing
	newPluginDTO := PluginDTO{
		Scope: "route",
		Name:  "get-sushi-restaurants",
		Plugin: model.PluginConfig{
			Name: "basic_auth",
			Config: map[string]interface{}{
				"password": "changeme",
				"username": "admin",
			},
			Enabled: false,
		},
	}

	// Marshal the RouteDTO to JSON
	jsonPayload, _ := json.Marshal(newPluginDTO)
	slog.Info(string(jsonPayload))

	// Create a POST request for adding the route
	req, err := http.NewRequest("PUT", "/plugin", bytes.NewBuffer(jsonPayload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := pluginController.UpdatePlugin()
	handler.ServeHTTP(rr, req)
	slog.Info(rr.Body.String())
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Plugin updated successfully")
}

func TestPluginController_DeletePlugin(t *testing.T) {
	IntegrationTestGuard(t)
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	pluginRepo := db.NewPluginRepository(database)
	pluginController := NewPluginController(pluginRepo)

	// Define query parameters for the test
	scope := "route"
	name := "get-sushi-restaurants"
	pluginName := "basic_auth"

	// Create a DELETE request with query parameters instead of a JSON body
	req, err := http.NewRequest("DELETE", "/plugin?scope="+scope+"&name="+name+"&plugin_name="+pluginName, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := pluginController.DeletePlugin()
	handler.ServeHTTP(rr, req)

	// Output response for debugging
	slog.Info(rr.Body.String())

	// Check for successful deletion
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Plugin deleted successfully")
}
