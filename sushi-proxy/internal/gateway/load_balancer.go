package gateway

import (
	"sync"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// Contains all logic related to getting the upstream for load balancing based on the load balancing strategy.
type LoadBalancer struct {
	healthChecker *HealthChecker
}

// Create a round robin cache based on service name
// Stores the counter of upstream to map to
var roundRobinCache sync.Map

// Create a consistent hash cache based on service name
// Stores the consistent hash ring for each service
var consistentHashCache sync.Map

func NewLoadBalancer(healthChecker *HealthChecker) *LoadBalancer {
	return &LoadBalancer{healthChecker: healthChecker}
}

// Gets the index of upstream to forward the request to based on the load balancing algorithm
func (lb *LoadBalancer) GetNextUpstream(service model.Service, clientIP string) int {
	switch service.LoadBalancingStrategy {
	case model.RoundRobin:
		return lb.handleRoundRobin(service)
	case model.IPHash:
		return lb.handleIPHash(service, clientIP)
	default:
		return 0
	}
}

// Get the current upstream request is routed to.
func (lb *LoadBalancer) GetCurrentUpstream(service model.Service, clientIP string) int {
	switch service.LoadBalancingStrategy {
	case model.RoundRobin:
		if val, ok := roundRobinCache.Load(service.Name); ok {
			return val.(int)
		}
		return 0
	case model.IPHash:
		return lb.handleIPHash(service, clientIP)
	default:
		return 0
	}
}

// ResetLoadBalancers the load balancer caches
func ResetLoadBalancers() {
	roundRobinCache = sync.Map{}
	consistentHashCache = sync.Map{}
}

func (lb *LoadBalancer) handleIPHash(service model.Service, clientIP string) int {
	// Get or create consistent hash ring for this service
	ring, _ := consistentHashCache.LoadOrStore(service.Name, NewConsistentHashRing(service))
	consistentRing := ring.(*ConsistentHashRing)

	// If health check is not enabled, just use the consistent hash ring directly
	if !service.Health.Enabled {
		if len(service.Upstreams) == 1 {
			return 0
		}

		// Get the upstream from the ring using client IP
		upstream := consistentRing.GetUpstream(clientIP)

		// Find the index of the upstream in the service's upstreams
		for i, u := range service.Upstreams {
			if u.Id == upstream.Id {
				return i
			}
		}
		return 0
	}

	// Health check is enabled, so we need to get the healthy upstreams
	healthyUpstreams := lb.healthChecker.GetHealthyUpstreams(service)
	if len(healthyUpstreams) == 0 {
		return model.NoUpstreamsAvailable
	}

	if len(service.Upstreams) == 1 {
		return 0
	}

	// Get the upstream from the ring using client IP
	upstream := consistentRing.GetUpstream(clientIP)

	// Check if the selected upstream is healthy
	for i, u := range service.Upstreams {
		if u.Id == upstream.Id {
			if status, exists := lb.healthChecker.serviceHealthMap[service.Name][u.Id]; exists {
				if status == Healthy {
					return i
				}
			}
		}
	}

	// If the selected upstream is not healthy, find the first healthy one
	for i, u := range service.Upstreams {
		if status, exists := lb.healthChecker.serviceHealthMap[service.Name][u.Id]; exists {
			if status == Healthy {
				return i
			}
		}
	}

	return model.NoUpstreamsAvailable
}

func (lb *LoadBalancer) handleRoundRobin(service model.Service) int {

	// TODO: probably refactor this code, it's a bit messy
	// Health check is not enabled, so we just round robin through all upstreams
	if !service.Health.Enabled {
		if len(service.Upstreams) == 1 {
			return 0
		}

		currentVal, _ := roundRobinCache.LoadOrStore(service.Name, 0)
		currentIndex := currentVal.(int)

		nextIndex := (currentIndex + 1) % len(service.Upstreams)
		roundRobinCache.Store(service.Name, nextIndex)

		return currentIndex
	}

	// Health check is enabled, so we need to get the healthy upstreams to round robin through
	healthyUpstreams := lb.healthChecker.GetHealthyUpstreams(service)
	if len(healthyUpstreams) == 0 {
		return model.NoUpstreamsAvailable
	}

	if len(service.Upstreams) == 1 {
		return 0
	}

	// Get or initialize the current index
	currentVal, _ := roundRobinCache.LoadOrStore(service.Name, 0)
	currentIndex := currentVal.(int)

	numUpstreams := len(service.Upstreams)
	// Try to find the next healthy upstream
	for i := 0; i < numUpstreams; i++ {
		// Calculate next index with wraparound
		candidateIndex := (currentIndex + i) % numUpstreams
		upstream := service.Upstreams[candidateIndex]

		// Check if upstream is healthy
		if status, exists := lb.healthChecker.serviceHealthMap[service.Name][upstream.Id]; exists {
			if status == Healthy {
				// Store the next index for subsequent requests
				nextIndex := (candidateIndex + 1) % numUpstreams
				roundRobinCache.Store(service.Name, nextIndex)
				return candidateIndex
			}
		}
	}

	// No healthy upstreams found,
	return model.NoUpstreamsAvailable
}
