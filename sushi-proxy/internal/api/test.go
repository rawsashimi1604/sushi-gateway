package api

import (
	"os"
	"testing"
)

func IntegrationTestGuard(t *testing.T) {
	if os.Getenv("SKIP_INTEGRATION_TEST") == "true" {
		t.Skip("Skipping test because SKIP_INTEGRATION_TEST is set")
	}
}
