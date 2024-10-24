package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type ServiceController struct {
}

func NewServiceController() *ServiceController {
	return &ServiceController{}
}

func (s *ServiceController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").Methods("GET").Handler(
		ProtectRouteUsingJWT(s.GetServices()))
}

func (s *ServiceController) GetServices() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (s *ServiceController) AddService() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func (s *ServiceController) DeleteServiceByName() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := json.Marshal("hello")
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
