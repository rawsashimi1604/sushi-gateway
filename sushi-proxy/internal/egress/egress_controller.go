package egress

import (
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins/plugin_manager"
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
		pluginManager, err := plugin_manager.NewPluginManagerFromConfig(req)
		if err != nil {
			slog.Info(err.Error())
			err.WriteJSONResponse(w)
			return
		}

		// Chain the plugins with the final handler where the request is forwarded.
		chainedHandler := pluginManager.ExecutePlugins(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// After executing all the plugins, handle the end result here.
			body, _, err := c.proxyService.HandleProxyPass(w, r)
			if err != nil {
				slog.Info(err.Error())
			}
			w.Write(body)
		}))

		// Execute the request (plugins + proxying).
		chainedHandler.ServeHTTP(w, req)
	}
}
