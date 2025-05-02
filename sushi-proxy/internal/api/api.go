package api

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"github.com/rs/cors"
)

const DEFAULT_CORS_ORIGIN = "http://localhost:5173"

func NewAdminApiRouter() http.Handler {
	slog.Info("Creating new admin api router...")
	router := mux.NewRouter()

	gatewayController := NewGatewayController()
	gatewayController.RegisterRoutes(router)

	authController := NewAuthController()
	authController.RegisterRoutes(router)

	healthController := NewHealthController()
	healthController.RegisterRoutes(router)

	corsOrigin := gateway.GlobalAppConfig.AdminCorsOrigin
	if corsOrigin == "" {
		corsOrigin = DEFAULT_CORS_ORIGIN
	}

	corsRouter := cors.New(cors.Options{
		AllowedOrigins:   []string{corsOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	slog.Info("Successfully created admin api router...")
	return corsRouter.Handler(router)
}
