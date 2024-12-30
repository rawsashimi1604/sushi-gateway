package gateway

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateSelfSignedCerts(t *testing.T) {
	// Create temp directory for test certificates
	tempDir := t.TempDir()

	// Generate certificates
	err := GenerateSelfSignedCerts(tempDir)
	if err != nil {
		t.Fatalf("Failed to generate certificates: %v", err)
	}

	// Verify server certificate exists and is valid
	certPath := filepath.Join(tempDir, "server.crt")
	certBytes, err := os.ReadFile(certPath)
	if err != nil {
		t.Fatalf("Failed to read server certificate: %v", err)
	}

	// Parse certificate
	block, _ := pem.Decode(certBytes)
	if block == nil {
		t.Fatal("Failed to decode PEM block from server certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse server certificate: %v", err)
	}

	// Verify certificate properties
	if cert.Subject.CommonName != "sushi.gateway.local" {
		t.Errorf("Expected CommonName to be 'sushi.gateway.local', got %s", cert.Subject.CommonName)
	}

	if len(cert.DNSNames) != 2 {
		t.Errorf("Expected 2 DNS names, got %d", len(cert.DNSNames))
	}

	// Verify server key exists
	keyPath := filepath.Join(tempDir, "server.key")
	keyBytes, err := os.ReadFile(keyPath)
	if err != nil {
		t.Fatalf("Failed to read server key: %v", err)
	}

	// Parse private key
	block, _ = pem.Decode(keyBytes)
	if block == nil {
		t.Fatal("Failed to decode PEM block from server key")
	}

	if block.Type != "RSA PRIVATE KEY" {
		t.Errorf("Expected key type to be 'RSA PRIVATE KEY', got %s", block.Type)
	}
}
