package util

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"net/http"
	"strconv"
	"strings"
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
