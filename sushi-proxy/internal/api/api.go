package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
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
	if database != nil {
		routeController = NewRouteController(nil)
	} else {
		routeController = NewRouteController(db.NewRouteRepository(database))
	}
	routeController.RegisterRoutes(router)

	// Plugin Resource
	var pluginController *PluginController
	if database != nil {
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
