package ingress

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type Controller struct {
}

func (c *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/route", c.RouteRequest()).Methods(http.MethodGet)
}

func (c *Controller) RouteRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Handing some route request.")
	}
}
