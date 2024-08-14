package gateway

import (
	"log/slog"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	slog.Info("Creating new router...")
	router := mux.NewRouter()

	egressController := NewSushiProxy()
	egressController.RegisterRoutes(router)

	slog.Info("Successfully created router...")
	return router
}
