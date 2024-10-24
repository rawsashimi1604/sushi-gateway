package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"log/slog"
	"net/http"
)

type RouteController struct {
	routeRepo *db.RouteRepository
}

func NewRouteController(routeRepo *db.RouteRepository) *RouteController {
	return &RouteController{routeRepo: routeRepo}
}

func (r *RouteController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/route").Methods("POST").Handler(ProtectRouteUsingJWT(r.AddRoute()))
	router.PathPrefix("/route").Methods("DELETE").Handler(ProtectRouteUsingJWT(r.DeleteRoute()))
}

func (r *RouteController) AddRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// RouteDTO represents the structure for adding a new route, including the service name.
		type RouteDTO struct {
			ServiceName string      `json:"service_name"`
			Route       model.Route `json:"route"`
		}

		var routeDTO RouteDTO
		if err := json.NewDecoder(req.Body).Decode(&routeDTO); err != nil {
			slog.Info("Failed to decode route DTO from request: " + err.Error())
			httperr := &gateway.HttpError{
				Code:     "CREATE_SERVICE_ERR",
				Message:  "failed to decode route from request body",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate the route data
		if routeDTO.ServiceName == "" || routeDTO.Route.Name == "" || routeDTO.Route.Path == "" || len(routeDTO.Route.Methods) == 0 {
			slog.Info("Invalid route data, missing required fields")
			http.Error(w, "Missing required route fields", http.StatusBadRequest)
			return
		}

		// Call the repository to add the route
		err := r.routeRepo.AddRoute(routeDTO.ServiceName, routeDTO.Route)
		if err != nil {
			slog.Info("Failed to add route to the repository: " + err.Error())
			http.Error(w, "Failed to add route", http.StatusInternalServerError)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Route created successfully",
		})
	}
}

func (r *RouteController) DeleteRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
