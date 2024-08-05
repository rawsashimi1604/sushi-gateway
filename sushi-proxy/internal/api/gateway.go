package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"log/slog"
	"net/http"
)

type GatewayController struct {
}

func NewGatewayController() *GatewayController {
	return &GatewayController{}
}

func (c *GatewayController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Methods("GET").HandlerFunc(c.GetGatewayInformation())
}

func (c *GatewayController) GetGatewayInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("Getting Gateway Information...")
		gatewayConfig := config.GlobalProxyConfig
		payload, _ := json.Marshal(gatewayConfig)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
