package gateway

import (
	"sync"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// Contains all logic related to getting the upstream for load balancing based on the load balancing strategy.
type LoadBalancer struct {
	mu sync.Mutex
}

// Create a round robin cache based on service name
// Stores the counter of upstream to map to
var roundRobinCache = make(map[string]int)

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
		return roundRobinCache[service.Name]
	default:
		return 0
	}
}

// ResetLoadBalancers the load balancer caches
func ResetLoadBalancers() {
	roundRobinCache = make(map[string]int)
}

func (lb *LoadBalancer) handleRoundRobin(service model.Service) int {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if len(service.Upstreams) == 1 {
		return 0
	}

	_, exists := roundRobinCache[service.Name]
	if !exists {
		roundRobinCache[service.Name] = 0
		return 0
	}

	roundRobinCache[service.Name] = (roundRobinCache[service.Name] + 1) % len(service.Upstreams)
	return roundRobinCache[service.Name]
}
