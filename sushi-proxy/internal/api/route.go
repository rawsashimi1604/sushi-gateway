package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type RouteController struct {
}

func NewRouteController() *RouteController {
	return &RouteController{}
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
