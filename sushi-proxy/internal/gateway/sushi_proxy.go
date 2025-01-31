package gateway

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

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

		// We execute the plugins in different phases.
		// Access phase > Log Phase

		// Chain the plugins with the final handler where the request is forwarded.
		finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// After executing all the plugins, handle the end result here.
			err := proxy.HandleProxyPass(w, r)
			if err != nil {
				slog.Info(err.Error())
				err.WriteJSONResponse(w)
				return
			}
		})

		chainedHandler := pluginManager.ExecutePlugins(AccessPhase, finalHandler)
		chainedHandler = pluginManager.ExecutePlugins(LogPhase, chainedHandler)

		// Execute the request (plugins + proxying).
		chainedHandler.ServeHTTP(captureWriter, req)

		// After whole request lifecycle, write the response from the upstream API to the client.
		w.Write(captureWriter.body.Bytes())
	}
}

func (s *SushiProxy) HandleProxyPass(w http.ResponseWriter, req *http.Request) *model.HttpError {

	path, convertErr := s.convertPathToProxyPassUrl(req)
	if convertErr != nil {
		return convertErr
	}
	target, err := url.Parse(path)

	if err != nil {
		return &model.HttpError{
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

		// Set Api Gateway headers
		req.Header.Set(constant.X_FORWARDED_HOST, req.Header.Get("Host"))
		req.Header.Set(constant.X_FORWARDED_FOR, req.RemoteAddr)

		req.Host = target.Host
	}
	proxy.ServeHTTP(w, req)
	return nil
}

// Routing logic, get the URL to proxy the request to.
func (s *SushiProxy) convertPathToProxyPassUrl(req *http.Request) (string, *model.HttpError) {
	matchedService, _, err := util.GetServiceAndRouteFromRequest(&GlobalProxyConfig, req)
	if err != nil {
		return "", err
	}

	// Handle load balancing
	loadBalancer := NewLoadBalancer(GlobalHealthChecker)
	remoteIpAddress, err := util.GetHostIp(req.RemoteAddr)
	if err != nil {
		slog.Error("Error getting remote IP address: " + err.Error())
		remoteIpAddress = "default"
		return "", err
	}

	upstreamIndex := loadBalancer.GetNextUpstream(*matchedService, remoteIpAddress)
	if upstreamIndex == model.NoUpstreamsAvailable {
		return "", &model.HttpError{
			Code:     "ERROR_NO_UPSTREAMS_AVAILABLE",
			Message:  "No upstreams available for service: " + matchedService.Name,
			HttpCode: http.StatusServiceUnavailable,
		}
	}

	upstream := matchedService.Upstreams[upstreamIndex]

	// Get the proxy URL...
	path := retrieveFullPathFromRequest(req.URL.Path)
	proxyURL := fmt.Sprintf("%s://%s:%d%s", matchedService.Protocol, upstream.Host, upstream.Port, path)
	return proxyURL, nil
}

// Should get the path after splitting URL from request, join them, and return the full path.
func retrieveFullPathFromRequest(path string) string {
	// Convert /service/route/v1/anything to /route/v1/anything
	parts := strings.Split(path, "/")
	return "/" + strings.Join(parts[2:], "/")
}
