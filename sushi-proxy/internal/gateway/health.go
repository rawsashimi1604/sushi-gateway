package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// Health status of a service
// Healthy = Service is available
// Unhealthy = Service is not available
// NotAvailable = Service does not have health check endpoint turned on, or health check has not started yet
type HealthStatus string

const (
	Healthy      HealthStatus = "healthy"
	Unhealthy    HealthStatus = "unhealthy"
	NotAvailable HealthStatus = "not_available"
)

var GlobalHealthChecker = NewHealthChecker()

// service -> upstream -> health status
type HealthChecker struct {
	serviceHealthMap map[string]map[string]HealthStatus
}

func NewHealthChecker() *HealthChecker {
	serviceHealthMap := make(map[string]map[string]HealthStatus)
	return &HealthChecker{
		serviceHealthMap: serviceHealthMap,
	}
}

func (hc *HealthChecker) Initialize() {
	for _, service := range GlobalProxyConfig.Services {
		hc.serviceHealthMap[service.Name] = make(map[string]HealthStatus)
		for _, upstream := range service.Upstreams {
			hc.serviceHealthMap[service.Name][upstream.Id] = NotAvailable
		}
	}
}

func (hc *HealthChecker) CheckHealthForAllServices() {

	servicesToCheck := GlobalProxyConfig.Services

	// Ping all services /health route asynchronously
	var wg sync.WaitGroup

	// Count number of upstreams to check. Need to do this before starting the goroutines to sync up with waitGroup
	upstreamsToCheck := 0
	for _, service := range servicesToCheck {
		upstreamsToCheck += len(service.Upstreams)
	}
	wg.Add(upstreamsToCheck)

	slog.Info("Checking health for all services defined in proxy configuration...")
	for _, service := range servicesToCheck {
		slog.Info("Checking health for service: " + service.Name)

		if !service.Health.Enabled {
			slog.Info("Skipping health check since not enabled for service: " + service.Name)
			continue
		}

		for _, upstream := range service.Upstreams {
			s := service  // Have to redeclare variable or goroutine wont work async
			u := upstream // Have to redeclare variable or goroutine wont work async

			go func(s *model.Service, u *model.Upstream) {
				defer wg.Done()

				healthCheckPath := fmt.Sprintf("%s://%s:%d%s", s.Protocol, u.Host, u.Port, s.Health.Path)

				res, err := http.Get(healthCheckPath)
				if err != nil {
					hc.serviceHealthMap[s.Name][u.Id] = Unhealthy
					slog.Error("Service: " + s.Name + " upstream: " + u.Id + " is unhealthy. Something went wrong, unable to ping health path: " + healthCheckPath)
					return
				}

				if res.StatusCode != http.StatusOK {
					hc.serviceHealthMap[s.Name][u.Id] = Unhealthy
					slog.Error("Service: " + s.Name + " upstream: " + u.Id + " is unhealthy. Status not 200 OK for health path: " + healthCheckPath)
					return
				}

				hc.serviceHealthMap[s.Name][u.Id] = Healthy
				slog.Info("Service: " + s.Name + " upstream: " + u.Id + " is healthy. Status 200 OK for health path: " + healthCheckPath)
			}(&s, &u)
		}
	}
	wg.Wait()
}

func (hc *HealthChecker) GetHealthyUpstreams(service model.Service) []model.Upstream {

	// Skip if health check is not enabled.
	isHealthCheckEnabled := service.Health.Enabled
	if !isHealthCheckEnabled {
		slog.Info("Skip getting healthy upstreams, health check is not enabled for service: " + service.Name)
		return service.Upstreams
	}

	healthyUpstreams := make([]model.Upstream, 0)
	for _, upstream := range service.Upstreams {
		if hc.serviceHealthMap[service.Name][upstream.Id] == Healthy {
			healthyUpstreams = append(healthyUpstreams, upstream)
		}
	}

	return healthyUpstreams
}
