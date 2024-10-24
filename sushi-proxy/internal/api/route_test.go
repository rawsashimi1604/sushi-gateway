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

func TestRouteController_AddRoute(t *testing.T) {
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	routeRepo := db.NewRouteRepository(database)
	serviceRepo := db.NewServiceRepository(database)
	routeController := NewRouteController(routeRepo, serviceRepo)

	// Create the RouteDTO for testing
	newRouteDTO := RouteDTO{
		ServiceName: "sushi-svc",
		Route: model.Route{
			Name:    "get-sushi-2",
			Path:    "/v1/sushi",
			Methods: []string{"GET"},
			Plugins: []model.PluginConfig{
				{
					Name: "rate_limit",
					Config: map[string]interface{}{
						"limit_hour":   10,
						"limit_min":    10,
						"limit_second": 10,
					},
					Enabled: true,
				},
			},
		},
	}

	// Marshal the RouteDTO to JSON
	jsonPayload, _ := json.Marshal(newRouteDTO)
	slog.Info(string(jsonPayload))

	// Create a POST request for adding the route
	req, err := http.NewRequest("POST", "/routes", bytes.NewBuffer(jsonPayload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := routeController.AddRoute()
	handler.ServeHTTP(rr, req)
	slog.Info(rr.Body.String())
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Route created successfully")
}

func TestRouteController_DeleteRouteByName(t *testing.T) {
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	routeRepo := db.NewRouteRepository(database)
	serviceRepo := db.NewServiceRepository(database)
	routeController := NewRouteController(routeRepo, serviceRepo)

	routeName := "get-sushi-2"
	// Create a test HTTP request for DELETE /?name=get-sushi
	req, err := http.NewRequest("DELETE", "/?name="+routeName, nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the DeleteRouteByName handler
	handler := routeController.DeleteRouteByName()
	handler.ServeHTTP(rr, req)
	slog.Info(rr.Body.String())
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Route deleted successfully")

}
