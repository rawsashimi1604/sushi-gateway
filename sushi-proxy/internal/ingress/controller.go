package ingress

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
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
