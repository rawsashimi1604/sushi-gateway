package internal

import (
	"log/slog"

	"github.com/rawsashimi1604/sushi-gateway/internal/egress"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	slog.Info("Creating new router...")
	router := mux.NewRouter()

	proxyService := egress.NewProxyService("http://localhost:8080")
	egressController := egress.NewEgressController(proxyService)
	egressController.RegisterRoutes(router)

	slog.Info("Successfully created router...")
	return router
}
