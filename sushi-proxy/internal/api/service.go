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
	router.PathPrefix("/").
		Methods("GET").Handler(ProtectRouteUsingJWT(s.GetServices())).
		Methods("POST").Handler(ProtectRouteUsingJWT(s.AddService())).
		Methods("DELETE").Handler(ProtectRouteUsingJWT(s.DeleteServiceByName()))
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

		// Decode the incoming request body to a Service struct
		var newService gateway.Service
		if err := json.NewDecoder(req.Body).Decode(&newService); err != nil {
			slog.Info(err.Error())
			httperr := &gateway.HttpError{
				Code:     "CREATE_SERVICE_ERR",
				Message:  "failed to decode service from request",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Call the service repository to add the service
		services, err := s.serviceRepo.GetAllServices()
		for _, svc := range services {
			if svc.Name == newService.Name {
				httperr := &gateway.HttpError{
					Code:     "CREATE_SERVICE_ERR",
					Message:  "name already exists.",
					HttpCode: http.StatusBadRequest,
				}
				httperr.WriteLogMessage()
				httperr.WriteJSONResponse(w)
				return
			}

			if svc.BasePath == newService.BasePath {
				httperr := &gateway.HttpError{
					Code:     "CREATE_SERVICE_ERR",
					Message:  "base_path already exists.",
					HttpCode: http.StatusBadRequest,
				}
				httperr.WriteLogMessage()
				httperr.WriteJSONResponse(w)
				return
			}
		}

		// Generate UUIDs
		uuid := gateway.NewUUIDGenerator()
		uuid.GenerateUUIDForService(newService)

		// Add to the database
		err = s.serviceRepo.AddService(newService)
		if err != nil {
			slog.Info(err.Error())
			httperr := &gateway.HttpError{
				Code:     "CREATE_SERVICE_ERR",
				Message:  "failed to add service into database",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Service created successfully",
		})
	}
}

func (s *ServiceController) DeleteServiceByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Extract the service name from the query parameters
		serviceName := req.URL.Query().Get("name")

		if serviceName == "" {
			httperr := &gateway.HttpError{
				Code:     "DELETE_SERVICE_ERR",
				Message:  "service name is missing in the request",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Call the repository to delete the service by name
		err := s.serviceRepo.DeleteServiceByName(serviceName)
		if err != nil {
			slog.Info("failed to delete service: " + err.Error())
			httperr := &gateway.HttpError{
				Code:     "DELETE_SERVICE_ERR",
				Message:  "Failed to delete service from the database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Service deleted successfully",
		})
	}
}
