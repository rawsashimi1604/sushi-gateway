package load_balancer

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"sync"
)

type LoadBalancer struct {
	mu sync.Mutex
}

type LoadBalancingAlgorithm string

const (
	RoundRobin LoadBalancingAlgorithm = "round_robin"
	Weighted   LoadBalancingAlgorithm = "weighted"
	IPHash     LoadBalancingAlgorithm = "ip_hash"
)

// Create a round robin cache based on service name
// Stores the counter of upstream to map to
var roundRobinCache = make(map[string]int)

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{}
}

// Gets the index of upstream to forward the request to based on the load balancing algorithm
func (lb *LoadBalancer) GetNextUpstream(alg LoadBalancingAlgorithm, service models.Service) int {
	switch alg {
	case RoundRobin:
		return lb.handleRoundRobin(service)
	default:
		return 0
	}
}

// Reset the load balancer caches
func Reset() {
	roundRobinCache = make(map[string]int)
}

func (lb *LoadBalancer) handleRoundRobin(service models.Service) int {
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
