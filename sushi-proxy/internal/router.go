package internal

import (
	"log/slog"

	"github.com/rawsashimi1604/sushi-gateway/internal/ingress"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	slog.Info("Creating new router...")
	router := mux.NewRouter()

	ingressController := ingress.Controller{}
	ingressController.RegisterRoutes(router)

	slog.Info("Successfully created router...")
	return router
}
