package egress

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type EgressController struct {
	proxyService *ProxyService
}

func NewEgressController(ps *ProxyService) *EgressController {
	return &EgressController{
		proxyService: ps,
	}
}

func (c *EgressController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", c.RouteRequest()).Methods(http.MethodGet)
}

func (c *EgressController) RouteRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Handing some route request.")
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		
	}
}
