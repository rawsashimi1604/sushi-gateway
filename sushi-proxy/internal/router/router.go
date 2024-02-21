package router

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"log/slog"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/egress"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	slog.Info("Creating new router...")
	router := mux.NewRouter()

	egressService := egress.NewEgressService(config.GlobalAppConfig.ReverseProxyHttpUrl)
	egressController := egress.NewEgressController(egressService)
	egressController.RegisterRoutes(router)

	slog.Info("Successfully created router...")
	return router
}
