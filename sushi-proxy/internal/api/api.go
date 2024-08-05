package api

import (
	"github.com/gorilla/mux"
	"log/slog"
)

func NewAdminApiRouter() *mux.Router {
	slog.Info("Creating new admin api router...")
	router := mux.NewRouter()

	gatewayController := NewGatewayController()
	gatewayController.RegisterRoutes(router)

	authController := NewAuthController()
	authController.RegisterRoutes(router)

	slog.Info("Successfully created admin api router...")
	return router
}
