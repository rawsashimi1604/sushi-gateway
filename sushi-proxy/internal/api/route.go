package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/validator"
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
			httperr := &model.HttpError{
				Code:     "CREATE_SERVICE_ERR",
				Message:  "failed to decode route from request body",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate the route
		// Check route does not exist...
		routes, err := r.routeRepo.GetAllRoutes(routeDTO.ServiceName)
		if err != nil {
			slog.Info(err.Error())
			httperr := &model.HttpError{
				Code:     "CREATE_ROUTE_ERR",
				Message:  "failed to add route into database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		for _, route := range routes {
			if route.Name == routeDTO.Route.Name {
				httperr := &model.HttpError{
					Code:     "CREATE_ROUTE_ERR",
					Message:  "route name already exits.",
					HttpCode: http.StatusBadRequest,
				}
				httperr.WriteLogMessage()
				httperr.WriteJSONResponse(w)
				return
			}
		}

		// Generic route validations
		routeValidator := validator.NewRouteValidator()
		if err := routeValidator.ValidateRoute(routeDTO.Route); err != nil {
			slog.Info("service validation failed")
			httperr := &model.HttpError{
				Code:     "CREATE_SERVICE_ERR",
				Message:  err.Error(),
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Call the repository to add the route
		err = r.routeRepo.AddRoute(routeDTO.ServiceName, routeDTO.Route)
		if err != nil {
			slog.Info("Failed to add route to the repository: " + err.Error())
			httperr := &model.HttpError{
				Code:     "CREATE_SERVICE_ERR",
				Message:  "failed to add route to the repository.",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
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

// DeleteRoute handles deleting a route by its name (DELETE request)
func (r *RouteController) DeleteRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Extract the route name from the query parameters
		routeName := req.URL.Query().Get("name")

		if routeName == "" {
			httperr := &model.HttpError{
				Code:     "DELETE_ROUTE_ERR",
				Message:  "route name is missing in the request",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Call the repository to delete the route by name
		err := r.routeRepo.DeleteRoute(routeName)
		if err != nil {
			slog.Info("failed to delete route: " + err.Error())
			httperr := &model.HttpError{
				Code:     "DELETE_ROUTE_ERR",
				Message:  "Failed to delete route from the database",
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
			"message": "Route deleted successfully",
		})
	}
}
