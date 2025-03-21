package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
)

type GatewayController struct {
}

func NewGatewayController() *GatewayController {
	return &GatewayController{}
}

func (c *GatewayController) RegisterRoutes(router *mux.Router) {
	router.Path("/gateway/config").Methods("GET").Handler(
		ProtectRouteUsingJWT(c.GetGatewayConfig()))
	router.Path("/gateway").Methods("GET").Handler(
		ProtectRouteUsingJWT(c.GetGatewayInformation()))
}

func (c *GatewayController) GetGatewayInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("GatewayController:: Admin API - Getting gateway information")
		gatewayConfig := gateway.GlobalProxyConfig
		payload, _ := json.Marshal(gatewayConfig)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (c *GatewayController) GetGatewayConfig() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("GatewayController:: Admin API - Getting gateway configuration")
		appConfig := gateway.GlobalAppConfig
		payload, _ := json.Marshal(appConfig)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
