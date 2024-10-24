package validator

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"strings"
)

type RouteValidator struct {
}

func NewRouteValidator() *RouteValidator {
	return &RouteValidator{}
}

func (rv *RouteValidator) ValidateRoute(route model.Route) error {
	if err := validatePath(route); err != nil {
		return err
	}
	if err := validateMethod(route); err != nil {
		return err
	}
	return nil
}

func validatePath(route model.Route) error {
	if !strings.HasPrefix(route.Path, "/") {
		return fmt.Errorf("route path: %s must start with /", route.Path)
	}

	if strings.HasSuffix(route.Path, "/") {
		return fmt.Errorf("route path: %s must not end with /", route.Path)
	}
	return nil
}

func validateMethod(route model.Route) error {

	if len(route.Methods) == 0 {
		return fmt.Errorf("route methods must be specified")
	}

	for _, method := range route.Methods {
		if !util.SliceContainsString(constant.VALID_METHODS, method) {
			return fmt.Errorf("route method: %s is invalid", method)
		}
	}

	return nil
}
