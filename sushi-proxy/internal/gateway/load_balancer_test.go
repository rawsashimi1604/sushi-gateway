package gateway

import (
	"testing"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

func TestRoundRobin(t *testing.T) {
	// Reset the load balancer state
	ResetLoadBalancers()
	lb := NewLoadBalancer()

	// Create a test service with round robin strategy
	service := model.Service{
		Name:                  "test-service",
		Protocol:              "http",
		BasePath:              "/test",
		LoadBalancingStrategy: model.RoundRobin,
		Upstreams: []model.Upstream{
			{
				Id:   "upstream1",
				Host: "localhost",
				Port: 8080,
			},
			{
				Id:   "upstream2",
				Host: "localhost",
				Port: 8081,
			},
			{
				Id:   "upstream3",
				Host: "localhost",
				Port: 8082,
			},
		},
	}

	// First round of requests should cycle through all upstreams
	firstIndex := lb.GetNextUpstream(service)
	if firstIndex != 0 {
		t.Errorf("First request: expected upstream index 0, got %d", firstIndex)
	}

	secondIndex := lb.GetNextUpstream(service)
	if secondIndex != 1 {
		t.Errorf("Second request: expected upstream index 1, got %d", secondIndex)
	}

	thirdIndex := lb.GetNextUpstream(service)
	if thirdIndex != 2 {
		t.Errorf("Third request: expected upstream index 2, got %d", thirdIndex)
	}

	// Fourth request should cycle back to the first upstream
	fourthIndex := lb.GetNextUpstream(service)
	if fourthIndex != 0 {
		t.Errorf("Fourth request: expected upstream index 0, got %d", fourthIndex)
	}

	ResetLoadBalancers()
}
