package api

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServiceController_GetServices(t *testing.T) {
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
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := serviceController.GetServices()
		handler.ServeHTTP(rr, req)
		slog.Info(rr.Body.String())
	})
}
