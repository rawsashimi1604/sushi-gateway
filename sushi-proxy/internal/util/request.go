package util

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

func GetServiceAndRouteFromRequest(proxyConfig *model.ProxyConfig, req *http.Request) (*model.Service, *model.Route, *model.HttpError) {
	path := req.URL.Path
	parts := strings.Split(path, "/")

	// Needs to have at least 3 parts for path to be valid:
	// 1. empty string, 2. service base path, 3. route path
	if len(parts) < 3 {
		return nil, nil, &model.HttpError{
			Code:     "INVALID_PATH",
			Message:  "Invalid path format, needs at least /service_base_path/...",
			HttpCode: http.StatusBadRequest,
		}
	}

	serviceBasePath := "/" + parts[1]
	routePath := "/" + strings.Join(parts[2:], "/")

	for _, service := range proxyConfig.Services {
		if service.BasePath == serviceBasePath {
			for _, route := range service.Routes {
				routeContainsMethod := SliceContainsString(route.Methods, req.Method)
				if MatchRoute(&route, routePath) && routeContainsMethod {
					return &service, &route, nil
				}
			}

			return nil, nil, &model.HttpError{
				Code:     "ROUTE_NOT_FOUND",
				Message:  fmt.Sprintf("Route not found for path: %s. Check your HTTP Method and Route path", routePath),
				HttpCode: http.StatusNotFound,
			}
		}
	}

	return nil, nil, &model.HttpError{
		Code:     "SERVICE_NOT_FOUND",
		Message:  "Service not found",
		HttpCode: http.StatusNotFound,
	}
}

// Gets the IP address from a remote address
func GetHostIp(remoteAddress string) (string, *model.HttpError) {

	ipAddr, _, err := net.SplitHostPort(remoteAddress)
	if err != nil {
		slog.Error("unable to get the host from ip address: " + err.Error())
		return "", model.NewHttpError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "something went wrong in the server.")
	}
	return ipAddr, nil
}

// GetHostAndPortFromURL parses a URL string (with scheme, example: http://127.0.0.1:8080) and returns the host and port
// If no port is specified in the URL:
// - HTTP defaults to 80
// - HTTPS defaults to 443
func GetHostAndPortFromURL(urlStr string) (string, int, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse URL: %w", err)
	}

	// Get host and port from URL
	host := parsedURL.Hostname()
	port := parsedURL.Port()

	// If port is not specified, use default ports based on scheme
	if port == "" {
		switch parsedURL.Scheme {
		case "http":
			port = "80"
		case "https":
			port = "443"
		default:
			return "", 0, fmt.Errorf("unsupported scheme: %s", parsedURL.Scheme)
		}
	}

	// Convert port string to integer
	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port number: %w", err)
	}

	return host, portNum, nil
}

// Check whether the route exists in the service, match either static or dynamic routes
func MatchRoute(route *model.Route, requestPath string) bool {
	// if not contains any {param} in route path, do a simple static match
	isStaticRoute := !strings.Contains(route.Path, "{") && !strings.Contains(route.Path, "}")
	if isStaticRoute {
		return route.Path == requestPath
	}

	// if contains {param} in route path, then do a dynamic match
	return matchDynamicRoute(route.Path, requestPath)
}

func matchDynamicRoute(routePath string, requestPath string) bool {
	routeParts := strings.Split(routePath, "/")
	requestPathParts := strings.Split(requestPath, "/")
	if len(routeParts) != len(requestPathParts) {
		return false
	}

	// Compare each segment
	for i := range routeParts {
		if strings.HasPrefix(routeParts[i], "{") && strings.HasSuffix(routeParts[i], "}") {
			// This is a dynamic segment, so continue without checking
			continue
		}

		// For static segments, they must match exactly
		if routeParts[i] != requestPathParts[i] {
			return false
		}
	}

	// All segments match for dynamic values
	return true
}

func GetContentLength(input string) int64 {
	if input == "" {
		return 0
	}
	conv, _ := strconv.ParseInt(input, 10, 64)
	return conv
}
