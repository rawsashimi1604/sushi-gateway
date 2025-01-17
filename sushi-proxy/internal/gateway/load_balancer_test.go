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
				result := lb.GetNextUpstream(service, "")
				assert.Equal(t, model.NoUpstreamsAvailable, result,
					"Expected NoUpstreamsAvailable when all upstreams are unhealthy")
			} else {
				// Test the sequence of returned indexes
				for _, expectedIdx := range tt.expectedIndexes {
					result := lb.GetNextUpstream(service, "")
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
			result := lb.GetNextUpstream(service, "")
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
	assert.Equal(t, 0, lb.GetNextUpstream(service, ""))
	assert.Equal(t, 1, lb.GetNextUpstream(service, ""))

	// Mark upstream1 as unhealthy
	healthChecker.serviceHealthMap["test-service"]["upstream1"] = Unhealthy

	// Should only return upstream2
	assert.Equal(t, 1, lb.GetNextUpstream(service, ""))
	assert.Equal(t, 1, lb.GetNextUpstream(service, ""))

	// Mark upstream1 as healthy again
	healthChecker.serviceHealthMap["test-service"]["upstream1"] = Healthy

	// Should resume round robin from last position
	assert.Equal(t, 0, lb.GetNextUpstream(service, ""))
	assert.Equal(t, 1, lb.GetNextUpstream(service, ""))

	// Mark both as unhealthy
	healthChecker.serviceHealthMap["test-service"]["upstream1"] = Unhealthy
	healthChecker.serviceHealthMap["test-service"]["upstream2"] = Unhealthy

	// Should return no upstreams available
	assert.Equal(t, model.NoUpstreamsAvailable, lb.GetNextUpstream(service, ""))
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
				result := lb.GetNextUpstream(service, "")
				assert.Equal(t, expectedIdx, result,
					"Expected upstream index %d but got %d when health check disabled", expectedIdx, result)
			}
		})
	}
}

func TestLoadBalancer_IPHash_ConsistentHashing(t *testing.T) {
	// Create test service with multiple upstreams
	service := model.Service{
		Name: "test-service",
		Health: model.Health{
			Enabled: false,
			Path:    "/mock",
		},
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
			{Id: "upstream3", Host: "localhost", Port: 8083},
		},
		LoadBalancingStrategy: model.IPHash,
	}

	tests := []struct {
		name     string
		clientIP string
		runCount int // Number of times to run the test to verify consistency
	}{
		{
			name:     "Same IP should always map to same upstream",
			clientIP: "192.168.1.1",
			runCount: 10,
		},
		{
			name:     "Different IP should potentially map to different upstream",
			clientIP: "192.168.1.2",
			runCount: 10,
		},
		{
			name:     "IPv6 address handling",
			clientIP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			runCount: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResetLoadBalancers()

			healthChecker := &HealthChecker{
				serviceHealthMap: map[string]map[string]HealthStatus{},
			}

			lb := NewLoadBalancer(healthChecker)

			// First run to get initial mapping
			firstIndex := lb.handleIPHash(service, tt.clientIP)

			// Verify the index is valid
			assert.GreaterOrEqual(t, firstIndex, 0)
			assert.Less(t, firstIndex, len(service.Upstreams))

			// Run multiple times to verify consistency
			for i := 0; i < tt.runCount; i++ {
				index := lb.handleIPHash(service, tt.clientIP)
				// Same IP should always map to the same upstream
				assert.Equal(t, firstIndex, index,
					"Same IP should map to same upstream on multiple calls")
			}

			// If we have a different test case, verify it maps to a potentially different upstream
			if len(tests) > 1 {
				differentIP := "192.168.1.100"
				if tt.clientIP == differentIP {
					differentIP = "192.168.1.200"
				}
				differentIndex := lb.handleIPHash(service, differentIP)
				// Note: There's a small chance this could fail if the hash happens to map to the same upstream
				// This is expected and acceptable in a real-world scenario
				if differentIndex == firstIndex {
					t.Logf("Different IP mapped to same upstream (this is possible but rare)")
				}
			}
		})
	}
}

func TestLoadBalancer_IPHash_HealthCheck(t *testing.T) {
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
		LoadBalancingStrategy: model.IPHash,
	}

	tests := []struct {
		name             string
		clientIP         string
		healthStatuses   map[string]map[string]HealthStatus
		expectNoUpstream bool
	}{
		{
			name:     "All upstreams healthy",
			clientIP: "192.168.1.1",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": Healthy,
					"upstream2": Healthy,
					"upstream3": Healthy,
				},
			},
			expectNoUpstream: false,
		},
		{
			name:     "Some upstreams unhealthy",
			clientIP: "192.168.1.2",
			healthStatuses: map[string]map[string]HealthStatus{
				"test-service": {
					"upstream1": Unhealthy,
					"upstream2": Healthy,
					"upstream3": Healthy,
				},
			},
			expectNoUpstream: false,
		},
		{
			name:     "All upstreams unhealthy",
			clientIP: "192.168.1.3",
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
			ResetLoadBalancers()

			healthChecker := &HealthChecker{
				serviceHealthMap: tt.healthStatuses,
			}

			lb := NewLoadBalancer(healthChecker)
			result := lb.handleIPHash(service, tt.clientIP)

			if tt.expectNoUpstream {
				assert.Equal(t, model.NoUpstreamsAvailable, result,
					"Expected no available upstreams")
			} else {
				assert.GreaterOrEqual(t, result, 0)
				assert.Less(t, result, len(service.Upstreams))

				// Verify the selected upstream is healthy
				upstream := service.Upstreams[result]
				status := tt.healthStatuses["test-service"][upstream.Id]
				assert.Equal(t, Healthy, status,
					"Selected upstream should be healthy")
			}
		})
	}
}

func TestLoadBalancer_IPHash_SingleUpstream(t *testing.T) {
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
		LoadBalancingStrategy: model.IPHash,
	}

	tests := []struct {
		name           string
		clientIP       string
		healthStatus   HealthStatus
		expectedResult int
	}{
		{
			name:           "Single healthy upstream",
			clientIP:       "192.168.1.1",
			healthStatus:   Healthy,
			expectedResult: 0,
		},
		{
			name:           "Single unhealthy upstream",
			clientIP:       "192.168.1.2",
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
			result := lb.handleIPHash(service, tt.clientIP)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestLoadBalancer_GetCurrentUpstream_IPHash(t *testing.T) {
	// Create test service with multiple upstreams
	service := model.Service{
		Name: "test-service",
		Health: model.Health{
			Enabled: false,
			Path:    "/mock",
		},
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
			{Id: "upstream3", Host: "localhost", Port: 8083},
		},
		LoadBalancingStrategy: model.IPHash,
	}

	tests := []struct {
		name     string
		clientIP string
	}{
		{
			name:     "Should return same upstream for same IP",
			clientIP: "192.168.1.1",
		},
		{
			name:     "Should handle IPv6",
			clientIP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ResetLoadBalancers()

			healthChecker := &HealthChecker{
				serviceHealthMap: map[string]map[string]HealthStatus{},
			}

			lb := NewLoadBalancer(healthChecker)

			// Get initial upstream
			initialUpstream := lb.GetCurrentUpstream(service, tt.clientIP)

			// Verify the index is valid
			assert.GreaterOrEqual(t, initialUpstream, 0)
			assert.Less(t, initialUpstream, len(service.Upstreams))

			// Multiple calls should return the same upstream for the same IP
			for i := 0; i < 5; i++ {
				currentUpstream := lb.GetCurrentUpstream(service, tt.clientIP)
				assert.Equal(t, initialUpstream, currentUpstream,
					"Same IP should map to same upstream on multiple GetCurrentUpstream calls")
			}

			// Different IP should potentially map to different upstream
			differentIP := "192.168.1.100"
			if tt.clientIP == differentIP {
				differentIP = "192.168.1.200"
			}
			differentUpstream := lb.GetCurrentUpstream(service, differentIP)

			// Note: There's a small chance this could be equal due to hash collision
			if differentUpstream == initialUpstream {
				t.Logf("Different IP mapped to same upstream (this is possible but rare)")
			}
		})
	}
}
