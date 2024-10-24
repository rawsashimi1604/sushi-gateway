package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"net/http"
)

type RouteController struct {
	routeRepo *db.RouteRepository
}

func NewRouteController(routeRepo *db.RouteRepository) *RouteController {
	return &RouteController{routeRepo: routeRepo}
}

func (r *RouteController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Methods("GET").Handler(
		ProtectRouteUsingJWT(r.AddRoute()))
}

func (r *RouteController) AddRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (r *RouteController) DeleteRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
