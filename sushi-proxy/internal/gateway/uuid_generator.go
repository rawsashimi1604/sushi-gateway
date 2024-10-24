package gateway

import (
	"github.com/google/uuid"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// Generates uuids for model objects.
type UUIDGenerator struct {
}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (u *UUIDGenerator) GenerateUUIDForService(service model.Service) {
	for i := range service.Upstreams {
		service.Upstreams[i].Id = uuid.New().String()
	}
	for i := range service.Plugins {
		service.Plugins[i].Id = uuid.New().String()
	}
	for i := range service.Routes {
		for j := range service.Routes[i].Plugins {
			service.Routes[i].Plugins[j].Id = uuid.New().String()
		}
	}
}

func (u *UUIDGenerator) GenerateUUIDForRoute(route model.Route) {
	for i := range route.Plugins {
		route.Plugins[i].Id = uuid.New().String()
	}
}
