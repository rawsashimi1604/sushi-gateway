package service

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type ServiceController struct {
}

func NewServiceController() *ServiceController {
	return &ServiceController{}
}

func (c *ServiceController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/api/v1/service").HandlerFunc(c.HandleIndex())
}

func (c *ServiceController) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("hello world from index.")
	}
}
