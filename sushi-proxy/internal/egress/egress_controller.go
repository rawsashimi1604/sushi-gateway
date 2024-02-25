package egress

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/analytics"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/basic_auth"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/rate_limit"
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
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("Handing request: " + req.URL.Path)
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")

		// Configure, register new plugins...
		pluginManager := plugins.NewPluginManager()
		pluginManager.RegisterPlugin(rate_limit.Plugin)
		pluginManager.RegisterPlugin(analytics.Plugin)
		pluginManager.RegisterPlugin(basic_auth.Plugin)

		// Chain the plugins with the final handler where the request is forwarded.
		chainedHandler := pluginManager.ExecutePlugins(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			slog.Info("Forwarding request...")

			// After executing all the plugins, handle the end result here.
			body, _, err := c.proxyService.HandleProxyPass(w, r)
			if err != nil {
				slog.Info(err.Error())
				slog.Info("Handle some error here...")
			}

			w.Write(body)
		}))

		// Execute the request (plugins + proxying).
		chainedHandler.ServeHTTP(w, req)
	}
}
