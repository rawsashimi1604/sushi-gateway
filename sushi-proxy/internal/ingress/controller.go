package ingress

import (
	"encoding/json"
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
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		data, err := json.Marshal(map[string]interface{}{
			"message": "some route request",
		})
		if err != nil {
			slog.Info("Could not marshall json request.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(data); err != nil {
			slog.Info("Something went wrong writing json response.")
			slog.Info(err.Error())
		}
	}
}
