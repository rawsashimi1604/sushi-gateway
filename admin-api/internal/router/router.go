package router

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/admin-api/internal/api/service"
	"log/slog"
)

func NewRouter() *mux.Router {
	slog.Info("Creating new router...")
	router := mux.NewRouter()

	serviceController := service.NewServiceController()
	serviceController.RegisterRoutes(router)
	
	slog.Info("Successfully created router...")
	return router
}
