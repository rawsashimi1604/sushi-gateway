package gateway

import (
	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

type ConsistentHashRing struct {
	service model.Service
	ring    *consistent.Consistent
}

// Create hasher interface used to uniformly distribute nodes/virtual nodes and keys.
type hasher struct{}

func (h hasher) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

type hashMember struct {
	upstream model.Upstream
}

func (hm hashMember) String() string {
	return hm.upstream.Id
}

func NewConsistentHashRing(service model.Service) *ConsistentHashRing {
	// Create a new consistent interface
	config := consistent.Config{
		PartitionCount:    1024,
		ReplicationFactor: 50,
		Load:              1.25,
		Hasher:            hasher{},
	}

	ring := consistent.New(nil, config)

	// Add the service's upstreams to the ring
	for _, upstream := range service.Upstreams {
		node := hashMember{upstream: upstream}
		ring.Add(node)
	}

	return &ConsistentHashRing{
		service: service,
		ring:    ring,
	}
}

// Get the upstream id for a given client IP address hash
func (chr *ConsistentHashRing) GetUpstream(hostIp string) model.Upstream {
	// Converts to byte first...
	key := []byte(hostIp)
	upstreamId := chr.ring.LocateKey(key)

	for _, upstream := range chr.service.Upstreams {
		if upstream.Id == upstreamId.String() {
			return upstream
		}
	}

	// Won't happen, since the upstream id is guaranteed to be in the ring
	return model.Upstream{}
}

func (chr *ConsistentHashRing) AddNewUpstream(upstream model.Upstream) {
	chr.ring.Add(hashMember{upstream: upstream})
}

func (chr *ConsistentHashRing) RemoveUpstream(upstreamId string) {
	chr.ring.Remove(upstreamId)
}
