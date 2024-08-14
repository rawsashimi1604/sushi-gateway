package gateway

import (
	"bytes"
	"github.com/gorilla/mux"
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

// captureResponseWriter is used to capture the HTTP response
type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
	body       bytes.Buffer
}

func newCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	// Default the status code to 200 in case WriteHeader is not called
	return &captureResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (w *captureResponseWriter) Write(data []byte) (int, error) {
	size, err := w.ResponseWriter.Write(data)
	w.size += size
	w.Header().Set("Content-Length", string(w.size))
	return size, err
}

func (w *captureResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (c *EgressController) RouteRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("Handing request: " + req.URL.Path)
		// TODO: check if necessary to add content-type
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		captureWriter := newCaptureResponseWriter(w)

		// Configure, register new plugins...
		pluginManager, err := NewPluginManagerFromConfig(req)
		if err != nil {
			slog.Info(err.Error())
			err.WriteJSONResponse(w)
			return
		}

		// Chain the plugins with the final handler where the request is forwarded.
		chainedHandler := pluginManager.ExecutePlugins(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// After executing all the plugins, handle the end result here.
			err := c.proxyService.HandleProxyPass(w, r)
			if err != nil {
				slog.Info(err.Error())
				err.WriteJSONResponse(w)
				return
			}
		}))

		// Execute the request (plugins + proxying).
		chainedHandler.ServeHTTP(captureWriter, req)

		// After whole request lifecycle.
		w.Write(captureWriter.body.Bytes())
	}
}
