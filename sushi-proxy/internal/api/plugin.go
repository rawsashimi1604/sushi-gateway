package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"net/http"
)

type PluginController struct {
	pluginRepo *db.PluginRepository
}

func NewPluginController(pluginRepo *db.PluginRepository) *PluginController {
	return &PluginController{pluginRepo: pluginRepo}
}

func (p *PluginController) RegisterRoutes(router *mux.Router) {
	router.Path("/").Methods("GET").Handler(
		ProtectRouteUsingJWT(p.GetPlugin()))
}

// TODO: create these routes.
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
