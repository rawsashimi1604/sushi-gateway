package egress

import (
	"bytes"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type EgressService struct {
}

func NewEgressService() *EgressService {
	return &EgressService{}
}

// captureResponseWriter is used to capture the HTTP response
type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func newCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	// Default the status code to 200 in case WriteHeader is not called
	return &captureResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (s *EgressService) HandleProxyPass(w http.ResponseWriter, req *http.Request) ([]byte, int, *EgressError) {

	path, convertErr := s.convertPathToProxyPassUrl(req)
	if convertErr != nil {
		return nil, 0, convertErr
	}

	slog.Info("path: " + path)
	target, err := url.Parse(path)
	if err != nil {
		return nil, 0, &EgressError{
			Code:     "ERROR_PARSING_PROXY_URL",
			Message:  "Error parsing URL when creating proxy_pass",
			HttpCode: http.StatusInternalServerError,
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	captureWriter := newCaptureResponseWriter(w)

	req.URL.Path = target.Path
	req.URL.Host = target.Host
	req.URL.Scheme = target.Scheme
	req.RequestURI = target.RequestURI()

	// Set the X-Forwarded-For header to the original request IP
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Header.Set("X-Forwarded-For", req.RemoteAddr)
	req.Host = target.Host

	// Serve the proxy and capture the response
	proxy.ServeHTTP(captureWriter, req)

	return captureWriter.body.Bytes(), captureWriter.statusCode, nil
}

func (s *EgressService) convertPathToProxyPassUrl(req *http.Request) (string, *EgressError) {
	// Split the path to get the service base path
	path := req.URL.Path

	// Ensure there's at least a base path segment following the initial slash
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return "", &EgressError{
			Code:     "INVALID_PATH",
			Message:  "Invalid path format, needs at least /service_base_path/...",
			HttpCode: http.StatusBadRequest,
		}
	}

	serviceBasePath := "/" + parts[1]

	// Search proxy config json file to find a matching base_path
	for _, service := range config.GlobalProxyConfig.Services {
		if service.BasePath == serviceBasePath {
			// For simplicity, just use the first upstream as the proxy target
			// TODO: add load balancing logic.
			if len(service.Upstreams) > 0 {
				upstream := service.Upstreams[0]
				// Check for route match
				route := "/" + strings.Join(parts[2:], "/")
				for _, r := range service.Routes {
					if r.Path == route && checkMethodMatch(req.Method, r.Methods) {
						// Success service and route matched!
						proxyURL := fmt.Sprintf("%s://%s:%d%s", service.Protocol, upstream.Host, upstream.Port, route)
						return proxyURL, nil
					}
				}

				return "", &EgressError{
					Code:     "ROUTE_NOT_FOUND",
					Message:  "Route not found",
					HttpCode: http.StatusNotFound,
				}
			}

			return "", &EgressError{
				Code:     "NO_UPSTREAMS",
				Message:  "No upstreams found for service, needs at least 1 upstream",
				HttpCode: http.StatusInternalServerError,
			}
		}
	}

	return "", &EgressError{
		Code:     "SERVICE_NOT_FOUND",
		Message:  "Service not found",
		HttpCode: http.StatusNotFound,
	}
}

func checkMethodMatch(requestMethod string, allowedMethods []string) bool {
	if len(allowedMethods) == 0 {
		// If no methods are specified, assume all methods are allowed
		return true
	}

	for _, method := range allowedMethods {
		if requestMethod == method {
			return true
		}
	}
	return false
}
