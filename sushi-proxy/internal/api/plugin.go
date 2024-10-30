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
	router.Path("/").Methods("PUT").Handler(
		ProtectRouteUsingJWT(
			ProtectRouteWhenUsingDblessMode(p.UpdatePlugin())))
	router.Path("/").Methods("POST").Handler(
		ProtectRouteUsingJWT(
			ProtectRouteWhenUsingDblessMode(p.AddPlugin())))
	router.Path("/").Methods("DELETE").Handler(
		ProtectRouteUsingJWT(
			ProtectRouteWhenUsingDblessMode(p.DeletePlugin())))
}

/*
	As a developer i want to update plugin.

	update plugin to global configuration
	update plugin to service
	update plugin to route

	add new plugin to global configuration
	add new plugin to service
	add new plugin to route

	delete plugin from global configuration
	delete plugin from service
	delete plugin from route.
*/

// TODO: create these routes.
func (p *PluginController) UpdatePlugin() http.HandlerFunc {
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
