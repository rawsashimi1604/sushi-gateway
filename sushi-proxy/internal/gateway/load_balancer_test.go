package gateway

import (
	"testing"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestLoadBalancer_RoundRobin_UnhealthyUpstreams(t *testing.T) {
	// Create test service with multiple upstreams
	service := model.Service{
		Name: "test-service",
		Health: model.Health{
			Enabled: true,
			Path:    "/mock",
		},
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
			{Id: "upstream3", Host: "localhost", Port: 8083},
		},
		LoadBalancingStrategy: model.RoundRobin,
	}

	tests := []struct {
		name             string
		healthStatuses   map[string]map[string]HealthStatus // service -> upstream -> status
		expectedIndexes  []int                              // sequence of expected indexes
		expectNoUpstream bool
	}{
		{
			name: "All upstreams healthy",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": Healthy,
					"upstream2": Healthy,
					"upstream3": Healthy,
				},
			},
			expectedIndexes: []int{0, 1, 2, 0}, // Should round robin through all
		},
		{
			name: "One upstream unhealthy",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": Healthy,
					"upstream2": Unhealthy,
					"upstream3": Healthy,
				},
			},
			expectedIndexes: []int{0, 2, 0, 2}, // Should skip index 1
		},
		{
			name: "Two upstreams unhealthy",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": Unhealthy,
					"upstream2": Unhealthy,
					"upstream3": Healthy,
				},
			},
			expectedIndexes: []int{2, 2, 2, 2}, // Should always return index 2
		},
		{
			name: "All upstreams unhealthy",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": Unhealthy,
					"upstream2": Unhealthy,
					"upstream3": Unhealthy,
				},
			},
			expectNoUpstream: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset the round robin cache before each test
			ResetLoadBalancers()

			// Create health checker with mock statuses
			healthChecker := &HealthChecker{
				serviceHealthMap: tt.healthStatuses,
			}

			lb := NewLoadBalancer(healthChecker)

			if tt.expectNoUpstream {
				// Test that we get NoUpstreamsAvailable when all are unhealthy
				result := lb.GetNextUpstream(service)
				assert.Equal(t, model.NoUpstreamsAvailable, result,
					"Expected NoUpstreamsAvailable when all upstreams are unhealthy")
			} else {
				// Test the sequence of returned indexes
				for _, expectedIdx := range tt.expectedIndexes {
					result := lb.GetNextUpstream(service)
					assert.Equal(t, expectedIdx, result,
						"Expected upstream index %d but got %d", expectedIdx, result)
				}
			}
		})
	}
}

func TestLoadBalancer_RoundRobin_SingleUpstream(t *testing.T) {
	// Test service with single upstream
	service := model.Service{
		Name: "test-service",
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
		},
		Health: model.Health{
			Enabled: true,
			Path:    "/mock",
		},
		LoadBalancingStrategy: model.RoundRobin,
	}

	tests := []struct {
		name           string
		healthStatus   HealthStatus
		expectedResult int
	}{
		{
			name:           "Single healthy upstream",
			healthStatus:   Healthy,
			expectedResult: 0,
		},
		{
			name:           "Single unhealthy upstream",
			healthStatus:   Unhealthy,
			expectedResult: model.NoUpstreamsAvailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResetLoadBalancers()

			healthChecker := &HealthChecker{
				serviceHealthMap: map[string]map[string]HealthStatus{
					"test-service": {
						"upstream1": tt.healthStatus,
					},
				},
			}

			lb := NewLoadBalancer(healthChecker)
			result := lb.GetNextUpstream(service)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestLoadBalancer_RoundRobin_HealthStateTransitions(t *testing.T) {
	service := model.Service{
		Name: "test-service",
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
		},
		Health: model.Health{
			Enabled: true,
			Path:    "/mock",
		},
		LoadBalancingStrategy: model.RoundRobin,
	}

	healthChecker := &HealthChecker{
		serviceHealthMap: map[string]map[string]HealthStatus{
			"test-service": {
				"upstream1": Healthy,
				"upstream2": Healthy,
			},
		},
	}

	lb := NewLoadBalancer(healthChecker)
	ResetLoadBalancers()

	// Initially both healthy, should round robin
	assert.Equal(t, 0, lb.GetNextUpstream(service))
	assert.Equal(t, 1, lb.GetNextUpstream(service))

	// Mark upstream1 as unhealthy
	healthChecker.serviceHealthMap["test-service"]["upstream1"] = Unhealthy

	// Should only return upstream2
	assert.Equal(t, 1, lb.GetNextUpstream(service))
	assert.Equal(t, 1, lb.GetNextUpstream(service))

	// Mark upstream1 as healthy again
	healthChecker.serviceHealthMap["test-service"]["upstream1"] = Healthy

	// Should resume round robin from last position
	assert.Equal(t, 0, lb.GetNextUpstream(service))
	assert.Equal(t, 1, lb.GetNextUpstream(service))

	// Mark both as unhealthy
	healthChecker.serviceHealthMap["test-service"]["upstream1"] = Unhealthy
	healthChecker.serviceHealthMap["test-service"]["upstream2"] = Unhealthy

	// Should return no upstreams available
	assert.Equal(t, model.NoUpstreamsAvailable, lb.GetNextUpstream(service))
}

func TestLoadBalancer_RoundRobin_HealthCheckDisabled(t *testing.T) {
	// Create test service with health check disabled
	service := model.Service{
		Name: "test-service",
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
			{Id: "upstream3", Host: "localhost", Port: 8083},
		},
		Health: model.Health{
			Enabled: false, // Health check disabled
			Path:    "/health",
		},
		LoadBalancingStrategy: model.RoundRobin,
	}

	tests := []struct {
		name            string
		healthStatuses  map[string]map[string]HealthStatus
		expectedIndexes []int
	}{
		{
			name: "Should round-robin if health check is not turned on",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": NotAvailable,
					"upstream2": NotAvailable,
					"upstream3": NotAvailable,
				},
			},
			expectedIndexes: []int{0, 1, 2, 0}, // Should round robin through all despite health status
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResetLoadBalancers()

			// Create health checker with mock statuses
			healthChecker := &HealthChecker{
				serviceHealthMap: tt.healthStatuses,
			}

			lb := NewLoadBalancer(healthChecker)

			// Test the sequence of returned indexes
			for _, expectedIdx := range tt.expectedIndexes {
				result := lb.GetNextUpstream(service)
				assert.Equal(t, expectedIdx, result,
					"Expected upstream index %d but got %d when health check disabled", expectedIdx, result)
			}
		})
	}
}
