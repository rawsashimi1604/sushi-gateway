package api

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
)

func NewAdminApiRouter() http.Handler {
	slog.Info("Creating new admin api router...")
	router := mux.NewRouter()

	gatewayController := NewGatewayController()
	gatewayController.RegisterRoutes(router)

	authController := NewAuthController()
	authController.RegisterRoutes(router)

	corsRouter := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	slog.Info("Successfully created admin api router...")
	return corsRouter.Handler(router)
}
