package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"log/slog"
	"net/http"
)

type ServiceController struct {
	serviceRepo *db.ServiceRepository
}

func NewServiceController(serviceRepo *db.ServiceRepository) *ServiceController {
	return &ServiceController{serviceRepo: serviceRepo}
}

func (s *ServiceController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Methods("GET").Handler(
		ProtectRouteUsingJWT(s.GetServices()))
}

func (s *ServiceController) GetServices() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		services, err := s.serviceRepo.GetAllServices()

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			slog.Info("something went wrong when getting the services.")
			httperr := &gateway.HttpError{
				Code:     "GET_SERVICE_ERR",
				Message:  "something went wrong when getting the services.",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}
		payload, _ := json.Marshal(services)
		w.Write(payload)
	}
}

func (s *ServiceController) AddService() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (s *ServiceController) DeleteServiceByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
