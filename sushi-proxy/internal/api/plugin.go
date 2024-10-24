package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

package api

import (
"encoding/json"
"github.com/gorilla/mux"
"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
"net/http"
)

type PluginController struct {
}

func NewPluginController() *PluginController {
	return &PluginController{}
}

func (p *PluginController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Methods("GET").Handler(
		ProtectRouteUsingJWT(p.GetPlugin()))
}

func (p *PluginController) GetPlugin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (p *PluginController) AddPlugin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (p *PluginController) DeletePlugin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}