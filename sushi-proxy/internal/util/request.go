package util

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"net/http"
	"strconv"
	"strings"
)

func GetServiceAndRouteFromRequest(proxyConfig *models.ProxyConfig, req *http.Request) (*models.Service, *models.Route, *errors.HttpError) {
	path := req.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return nil, nil, &errors.HttpError{
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
			return nil, nil, &errors.HttpError{
				Code:     "ROUTE_NOT_FOUND",
				Message:  "Route not found within the service",
				HttpCode: http.StatusNotFound,
			}
		}
	}

	return nil, nil, &errors.HttpError{
		Code:     "SERVICE_NOT_FOUND",
		Message:  "Service not found",
		HttpCode: http.StatusNotFound,
	}
}

func MatchRoute(route *models.Route, path string) bool {
	// TODO: Add url path params matching as well.
	return route.Path == path
}

func ParseContentLength(input string) int64 {
	if input == "" {
		return 0
	}
	conv, _ := strconv.ParseInt(input, 10, 64)
	return conv
}
