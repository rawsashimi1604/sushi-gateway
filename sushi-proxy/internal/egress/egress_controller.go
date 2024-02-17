package egress

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/plugins"
	"github.com/rawsashimi1604/sushi-gateway/plugins/analytics"
	"github.com/rawsashimi1604/sushi-gateway/plugins/rate_limit"
	"log/slog"
	"net/http"
)

type EgressController struct {
	proxyService *EgressService
}

func NewEgressController(ps *EgressService) *EgressController {
	return &EgressController{
		proxyService: ps,
	}
}

func (c *EgressController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(c.RouteRequest())
}

func (c *EgressController) RouteRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Handing request: " + r.URL.Path)
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")

		// Configure, register new plugins...
		pluginManager := plugins.NewPluginManager()
		pluginManager.RegisterPlugin(rate_limit.Plugin)
		pluginManager.RegisterPlugin(analytics.Plugin)

		// Chain the plugins with the final handler where the request is forwarded.
		chainedHandler := pluginManager.ExecutePlugins(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// After executing all the plugins, handle the end result here.
			body, code, err := c.proxyService.ForwardRequest(r)
			if err != nil {
				slog.Info("Handle some HaProxy error here...")
				if err.Code == constant.READ_HAPROXY_RESPONSE_BODY_ERROR_CODE {
					w.WriteHeader(err.HttpCode)
					jsonData, _ := json.Marshal(err)
					w.Write(jsonData)
					return // Ensure to return here to prevent further execution after writing the error response.
				}
			}

			w.WriteHeader(code)
			w.Write(body)
		}))

		// Execute the chained handler.
		chainedHandler.ServeHTTP(w, r)
	}
}
