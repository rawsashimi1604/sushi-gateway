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

// service -> upstream -> health status
type HealthChecker struct {
	serviceHealthMap map[string]map[string]HealthStatus
}

func NewHealtherChecker() *HealthChecker {
	// Initialize service health map
	serviceHealthMap := make(map[string]map[string]HealthStatus)
	for _, service := range GlobalProxyConfig.Services {
		serviceHealthMap[service.Name] = make(map[string]HealthStatus)
		for _, upstream := range service.Upstreams {
			serviceHealthMap[service.Name][upstream.Id] = NotAvailable
		}
	}

	return &HealthChecker{
		serviceHealthMap: serviceHealthMap,
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
			u := upstream // Have to redeclare variable or goroutine wont work async

			go func(u *model.Upstream) {
				defer wg.Done()

				healthCheckPath := fmt.Sprintf("%s://%s:%d%s", service.Protocol, u.Host, u.Port, service.Health.Path)

				res, err := http.Get(healthCheckPath)
				if err != nil {
					hc.serviceHealthMap[service.Name][u.Id] = Unhealthy
					slog.Error("Service: " + service.Name + " upstream: " + u.Id + " is unhealthy. Something went wrong, unable to ping health path: " + healthCheckPath)
					return
				}

				if res.StatusCode != http.StatusOK {
					hc.serviceHealthMap[service.Name][u.Id] = Unhealthy
					slog.Error("Service: " + service.Name + " upstream: " + u.Id + " is unhealthy. Status not 200 OK for health path: " + healthCheckPath)
					return
				}

				hc.serviceHealthMap[service.Name][u.Id] = Healthy
				slog.Info("Service: " + service.Name + " upstream: " + u.Id + " is healthy. Status 200 OK for health path: " + healthCheckPath)
			}(&u)
		}
	}
	wg.Wait()
}
