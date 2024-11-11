package api

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceController_GetServices(t *testing.T) {
	IntegrationTestGuard(t)
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	serviceRepo := db.NewServiceRepository(database)
	serviceController := NewServiceController(serviceRepo)

	router := mux.NewRouter()
	serviceController.RegisterRoutes(router)
	t.Run("success - get all services", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler := serviceController.GetServices()
		handler.ServeHTTP(rr, req)
		slog.Info(rr.Body.String())
	})
}

func TestServiceController_AddService(t *testing.T) {
	IntegrationTestGuard(t)
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	serviceRepo := db.NewServiceRepository(database)
	serviceController := NewServiceController(serviceRepo)

	newService := model.Service{
		Name:                  "sushi-svc-3",
		BasePath:              "/sushi-service-3",
		Protocol:              "http",
		LoadBalancingStrategy: "round_robin",
		Upstreams: []model.Upstream{
			{Host: "localhost", Port: 8001},
		},
	}
	jsonPayload, _ := json.Marshal(newService)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonPayload))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := serviceController.AddService()
	handler.ServeHTTP(rr, req)
	slog.Info(rr.Body.String())
}

func TestServiceController_DeleteServiceByName(t *testing.T) {
	IntegrationTestGuard(t)
	// Setup database and repository for the controller.
	database, err := db.ConnectDb()
	if err != nil {
		t.Fatal("unable to connect to database")
	}

	serviceRepo := db.NewServiceRepository(database)
	serviceController := NewServiceController(serviceRepo)

	serviceName := "sushi-svc-3"
	req, err := http.NewRequest("DELETE", "/?name="+serviceName, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := serviceController.DeleteServiceByName()
	handler.ServeHTTP(rr, req)
	slog.Info(rr.Body.String())
}
