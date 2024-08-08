package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"net/http"
)

type GatewayController struct {
}

func NewGatewayController() *GatewayController {
	return &GatewayController{}
}

func (c *GatewayController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Methods("GET").Handler(
		ProtectRouteUsingJWT(c.GetGatewayInformation()))
}

func (c *GatewayController) GetGatewayInformation() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		gatewayConfig := config.GlobalProxyConfig
		payload, _ := json.Marshal(gatewayConfig)
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
