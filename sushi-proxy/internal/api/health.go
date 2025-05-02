package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) RegisterRoutes(router *mux.Router) {
	router.Path("/healthz").Methods("GET").Handler(c.CheckHealth())
}

func (c *HealthController) CheckHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("HealthController:: Admin API - Gateway is healthy.")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	}
}
