package gateway

import (
	"sync"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// Contains all logic related to getting the upstream for load balancing based on the load balancing strategy.
type LoadBalancer struct{}

// Create a round robin cache based on service name
// Stores the counter of upstream to map to
var roundRobinCache sync.Map

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{}
}

// Gets the index of upstream to forward the request to based on the load balancing algorithm
func (lb *LoadBalancer) GetNextUpstream(service model.Service) int {
	switch service.LoadBalancingStrategy {
	case model.RoundRobin:
		return lb.handleRoundRobin(service)
	default:
		return 0
	}
}

// Get the current upstream request is routed to.
func (lb *LoadBalancer) GetCurrentUpstream(service model.Service) int {
	switch service.LoadBalancingStrategy {
	case model.RoundRobin:
		if val, ok := roundRobinCache.Load(service.Name); ok {
			return val.(int)
		}
		return 0
	default:
		return 0
	}
}

// ResetLoadBalancers the load balancer caches
func ResetLoadBalancers() {
	roundRobinCache = sync.Map{}
}

func (lb *LoadBalancer) handleRoundRobin(service model.Service) int {
	if len(service.Upstreams) == 1 {
		return 0
	}

	// Get or initialize the current index
	currentVal, _ := roundRobinCache.LoadOrStore(service.Name, 0)
	currentIndex := currentVal.(int)

	// Calculate and store next index
	nextIndex := (currentIndex + 1) % len(service.Upstreams)
	roundRobinCache.Store(service.Name, nextIndex)

	return currentIndex
}
