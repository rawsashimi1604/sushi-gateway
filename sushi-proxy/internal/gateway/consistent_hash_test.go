package gateway

import (
	"testing"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestConsistentHashRing_GetUpstream(t *testing.T) {
	// Create a test service with multiple upstreams
	service := model.Service{
		Name: "test-service",
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
			{Id: "upstream3", Host: "localhost", Port: 8083},
		},
	}

	// Create a new consistent hash ring
	chr := NewConsistentHashRing(service)

	// Test consistent hashing for multiple IPs
	tests := []struct {
		name     string
		hostIP   string
		runCount int // Number of times to run the test to verify consistency
	}{
		{
			name:     "Same IP should always map to same upstream",
			hostIP:   "192.168.1.1",
			runCount: 10,
		},
		{
			name:     "Different IP should potentially map to different upstream",
			hostIP:   "192.168.1.2",
			runCount: 10,
		},
		{
			name:     "IPv6 address handling",
			hostIP:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			runCount: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First run to get initial mapping
			firstUpstream := chr.GetUpstream(tt.hostIP)

			// Verify the upstream is valid
			assert.NotEmpty(t, firstUpstream.Id)
			assert.Contains(t, []string{"upstream1", "upstream2", "upstream3"}, firstUpstream.Id)

			// Run multiple times to verify consistency
			for i := 0; i < tt.runCount; i++ {
				upstream := chr.GetUpstream(tt.hostIP)
				// Same IP should always map to the same upstream
				assert.Equal(t, firstUpstream.Id, upstream.Id,
					"Same IP should map to same upstream on multiple calls")
			}
		})
	}
}

func TestConsistentHashRing_AddNewUpstream(t *testing.T) {
	// Create initial service with upstreams
	service := model.Service{
		Name: "test-service",
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
		},
	}

	chr := NewConsistentHashRing(service)

	// Test IPs and their initial mappings
	testIPs := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
		"192.168.1.4",
	}

	// Store initial mappings of client ip addresses to upstream Ids
	initialMappings := make(map[string]string)
	for _, ip := range testIPs {
		upstream := chr.GetUpstream(ip)
		initialMappings[ip] = upstream.Id
	}

	// Add new upstream
	newUpstream := model.Upstream{
		Id:   "upstream3",
		Host: "localhost",
		Port: 8083,
	}
	chr.AddNewUpstream(newUpstream)

	// Check new mappings
	remappedCount := 0
	for _, ip := range testIPs {
		newUpstream := chr.GetUpstream(ip)
		if newUpstream.Id != initialMappings[ip] {
			remappedCount++
		}
	}

	// Some IPs should be remapped, but not all
	// This verifies that adding a new node doesn't cause complete rehashing
	assert.True(t, remappedCount > 0 && remappedCount < len(testIPs),
		"Adding new upstream should cause some but not all IPs to be remapped")
}

func TestConsistentHashRing_RemoveUpstream(t *testing.T) {
	// Create initial service with upstreams
	service := model.Service{
		Name: "test-service",
		Upstreams: []model.Upstream{
			{Id: "upstream1", Host: "localhost", Port: 8081},
			{Id: "upstream2", Host: "localhost", Port: 8082},
			{Id: "upstream3", Host: "localhost", Port: 8083},
		},
	}

	chr := NewConsistentHashRing(service)

	// Test IPs
	testIPs := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
		"192.168.1.4",
	}

	// Store initial mappings
	initialMappings := make(map[string]string)
	for _, ip := range testIPs {
		upstream := chr.GetUpstream(ip)
		initialMappings[ip] = upstream.Id
	}

	// Remove an upstream
	chr.RemoveUpstream("upstream2")

	// Check new mappings
	remappedCount := 0
	for _, ip := range testIPs {
		newUpstream := chr.GetUpstream(ip)
		if initialMappings[ip] == "upstream2" {
			// IPs that were mapped to the removed upstream must be remapped
			assert.NotEqual(t, "upstream2", newUpstream.Id)
			remappedCount++
		} else if newUpstream.Id != initialMappings[ip] {
			// Some other IPs might be remapped too
			remappedCount++
		}
		// Verify we never get the removed upstream
		assert.NotEqual(t, "upstream2", newUpstream.Id)
	}

	// Verify that some remapping occurred
	assert.True(t, remappedCount > 0,
		"Removing an upstream should cause some IPs to be remapped")
}
