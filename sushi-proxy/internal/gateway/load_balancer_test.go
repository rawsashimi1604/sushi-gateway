package gateway

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"testing"
)

// Write test to test load balancer given a service with multiple upstreams

func TestRoundRobin(t *testing.T) {
	Reset()
	lb := NewLoadBalancer()
	service := model.Service{
		Name: "test",
		Upstreams: []model.Upstream{
			{
				Host: "localhost",
				Port: 8080,
			},
			{
				Host: "localhost",
				Port: 8081,
			},
			{
				Host: "localhost",
				Port: 8082,
			},
		},
	}
	firstReq := lb.handleRoundRobin(service)
	secondReq := lb.handleRoundRobin(service)
	thirdReq := lb.handleRoundRobin(service)

	if firstReq != 0 {
		t.Errorf("Expected 0, got %d", firstReq)
	}
	if secondReq != 1 {
		t.Errorf("Expected 1, got %d", secondReq)
	}
	if thirdReq != 2 {
		t.Errorf("Expected 2, got %d", thirdReq)
	}

	// Try the load balancer again it should round robin to first
	fourthReq := lb.handleRoundRobin(service)
	if fourthReq != 0 {
		t.Errorf("Expected 0, got %d", fourthReq)
	}
	Reset()
}
