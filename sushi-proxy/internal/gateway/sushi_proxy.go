package gateway

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// TODO: refactor code for routing logic.

type SushiProxy struct {
}

func NewSushiProxy() *SushiProxy {
	return &SushiProxy{}
}

func (proxy *SushiProxy) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(proxy.RouteRequest())
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
	w.Header().Set("Content-Length", string(rune(w.size)))
	return size, err
}

func (w *captureResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (proxy *SushiProxy) RouteRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("Handing request: " + req.URL.Path)
		// TODO: check if necessary to add content-type
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		captureWriter := newCaptureResponseWriter(w)

		// Register plugins from global, service and routes using the plugin manager.
		pluginManager, err := NewPluginManagerFromConfig(req)
		if err != nil {
			slog.Info(err.Error())
			err.WriteJSONResponse(w)
			return
		}

		// Chain the plugins with the final handler where the request is forwarded.
		chainedHandler := pluginManager.ExecutePlugins(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// After executing all the plugins, handle the end result here.
			err := proxy.HandleProxyPass(w, r)
			if err != nil {
				slog.Info(err.Error())
				err.WriteJSONResponse(w)
				return
			}
		}))

		// Execute the request (plugins + proxying).
		chainedHandler.ServeHTTP(captureWriter, req)

		// After whole request lifecycle, write the response from the upstream API to the client.
		w.Write(captureWriter.body.Bytes())
	}
}

func (s *SushiProxy) HandleProxyPass(w http.ResponseWriter, req *http.Request) *HttpError {

	path, convertErr := s.convertPathToProxyPassUrl(req)
	if convertErr != nil {
		return convertErr
	}
	target, err := url.Parse(path)

	if err != nil {
		return &HttpError{
			Code:     "ERROR_PARSING_PROXY_URL",
			Message:  "Error parsing URL when handling request.",
			HttpCode: http.StatusInternalServerError,
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Customize the Director to modify request before forwarding
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		// Call the original Director to preserve other behaviors
		originalDirector(req)
		req.URL.Path = target.Path
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Header.Set(constant.X_FORWARDED_HOST, req.Header.Get("Host"))
		req.Header.Set(constant.X_FORWARDED_FOR, req.RemoteAddr)
		req.Host = target.Host
	}
	proxy.ServeHTTP(w, req)
	return nil
}

// Routing logic, get the URL to proxy the request to.
func (s *SushiProxy) convertPathToProxyPassUrl(req *http.Request) (string, *HttpError) {
	// TODO: check the protocol http or https
	matchedService, matchedRoute, err := GetServiceAndRouteFromRequest(&GlobalProxyConfig, req)
	if err != nil {
		return "", err
	}

	// Handle load balancing
	loadBalancer := NewLoadBalancer()
	upstreamIndex := loadBalancer.GetNextUpstream(*matchedService)
	upstream := matchedService.Upstreams[upstreamIndex]

	proxyURL := fmt.Sprintf("%s://%s:%d%s", matchedService.Protocol, upstream.Host, upstream.Port, matchedRoute.Path)
	return proxyURL, nil
}
