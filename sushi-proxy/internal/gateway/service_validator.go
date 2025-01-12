package gateway

import (
	"fmt"
	"strings"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

type ServiceValidator struct {
}

func NewServiceValidator() *ServiceValidator {
	return &ServiceValidator{}
}

func (sv *ServiceValidator) ValidateService(service model.Service) error {
	if err := validateServiceLoadBalancing(&service); err != nil {
		return err
	}
	if err := validateBasePath(&service); err != nil {
		return err
	}
	if err := validateProtocol(&service); err != nil {
		return err
	}
	if err := validateUpstream(&service); err != nil {
		return err
	}
	if err := validateHealthCheckPath(&service); err != nil {
		return err
	}
	return nil
}

func validateServiceLoadBalancing(service *model.Service) error {
	isLoadBalancingAlgValid := service.LoadBalancingStrategy.IsValid()
	if !isLoadBalancingAlgValid {
		return fmt.Errorf("service load balancing strategy: %s is invalid", service.LoadBalancingStrategy)
	}
	return nil
}

func validateBasePath(service *model.Service) error {
	if !strings.HasPrefix(service.BasePath, "/") {
		return fmt.Errorf("service path: %s must start with /", service.BasePath)
	}
	if strings.HasSuffix(service.BasePath, "/") {
		return fmt.Errorf("service path: %s must not end with /", service.BasePath)
	}
	return nil
}

func validateProtocol(service *model.Service) error {
	if !util.SliceContainsString(constant.AVAILABLE_PROXY_PROTOCOLS, service.Protocol) {
		return fmt.Errorf("service protocol: %s is invalid, "+
			"only http and https supported", service.Protocol)
	}
	return nil
}

func validateUpstream(service *model.Service) error {
	if len(service.Upstreams) == 0 {
		return fmt.Errorf("service :%s must have at least one upstream", service.Name)
	}
	return nil
}

func validateHealthCheckPath(service *model.Service) error {
	// SKip as health check is not enabled
	if !service.Health.Enabled {
		return nil
	}
	if service.Health.Path == "" {
		return fmt.Errorf("service: %s health check path is required", service.Name)
	}
	if !strings.HasPrefix(service.Health.Path, "/") {
		return fmt.Errorf("service path: %s must start with /", service.Health.Path)
	}
	if strings.HasSuffix(service.Health.Path, "/") {
		return fmt.Errorf("service path: %s must not end with /", service.Health.Path)
	}

	return nil
}
