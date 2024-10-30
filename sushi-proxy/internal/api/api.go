package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
)

func NewAdminApiRouter(database *sql.DB) http.Handler {
	slog.Info("Creating new admin api router...")
	router := mux.NewRouter()

	gatewayController := NewGatewayController()
	gatewayController.RegisterRoutes(router)

	authController := NewAuthController()
	authController.RegisterRoutes(router)

	// Admin API only add routes if hosted in db mode.
	// Service Resource
	var serviceController *ServiceController
	if database == nil {
		serviceController = NewServiceController(nil)
	} else {
		serviceController = NewServiceController(db.NewServiceRepository(database))
	}
	serviceController.RegisterRoutes(router)

	// Route Resource
	var routeController *RouteController
	if database == nil {
		routeController = NewRouteController(nil, nil)
	} else {
		routeController = NewRouteController(db.NewRouteRepository(database),
			db.NewServiceRepository(database))
	}
	routeController.RegisterRoutes(router)

	// Plugin Resource
	var pluginController *PluginController
	if database == nil {
		pluginController = NewPluginController(nil)
	} else {
		pluginController = NewPluginController(db.NewPluginRepository(database))
	}
	pluginController.RegisterRoutes(router)

	corsRouter := cors.New(cors.Options{
		// TODO: externalize manager url
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	slog.Info("Successfully created admin api router...")
	return corsRouter.Handler(router)
}

// If dbless mode, we shouldn't execute any persistance logic.
// Configuration should only be injected from config.json.
func ProtectRouteWhenUsingDblessMode(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if gateway.GlobalAppConfig.PersistenceConfig == constant.DBLESS_MODE {
			slog.Info("Dbless mode detected, request not available!")
			w.Header().Set("Content-Type", "application/json")
			httperr := &model.HttpError{
				Code:     "ROUTE_NOT_AVAILABLE_ERR",
				Message:  "Gateway running in db-less mode, request not available.",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}
		next.ServeHTTP(w, req)
	})
}
