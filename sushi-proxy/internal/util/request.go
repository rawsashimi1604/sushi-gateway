package util

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"net/http"
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
				if strings.HasPrefix(routePath, route.Path) && SliceContainsString(route.Methods, req.Method) {
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
