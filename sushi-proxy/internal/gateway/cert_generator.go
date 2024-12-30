package gateway

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

// GenerateSelfSignedCerts generates a self-signed server certificate and key
func GenerateSelfSignedCerts(certDir string) error {
	slog.Info("Generating self-signed certificates...")

	// Generate paths
	serverCertPath := filepath.Join(certDir, "server.crt")
	serverKeyPath := filepath.Join(certDir, "server.key")

	// Generate server key pair
	serverPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate server private key: %v", err)
	}

	// Create server certificate template
	serverTemplate := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Sushi Gateway"},
			CommonName:   "sushi.gateway.local",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 365),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost", "sushi.gateway.local"},
	}

	// Create self-signed server certificate
	serverCertBytes, err := x509.CreateCertificate(rand.Reader, &serverTemplate, &serverTemplate, &serverPrivKey.PublicKey, serverPrivKey)
	if err != nil {
		return fmt.Errorf("failed to create server certificate: %v", err)
	}

	// Save server certificate and private key
	if err := saveCertAndKey(serverCertPath, serverKeyPath, serverCertBytes, serverPrivKey); err != nil {
		return fmt.Errorf("failed to save server cert and key: %v", err)
	}

	slog.Info("Successfully generated self-signed certificates",
		"server_cert", serverCertPath,
		"server_key", serverKeyPath)

	return nil
}

// saveCertAndKey saves a certificate and private key to files
func saveCertAndKey(certPath, keyPath string, certBytes []byte, privKey *rsa.PrivateKey) error {
	// Save certificate
	certFile, err := os.Create(certPath)
	if err != nil {
		return fmt.Errorf("failed to create cert file: %v", err)
	}
	defer certFile.Close()

	if err := pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}); err != nil {
		return fmt.Errorf("failed to encode certificate: %v", err)
	}

	// Save private key
	keyFile, err := os.Create(keyPath)
	if err != nil {
		return fmt.Errorf("failed to create key file: %v", err)
	}
	defer keyFile.Close()

	privKeyBytes := x509.MarshalPKCS1PrivateKey(privKey)
	if err := pem.Encode(keyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privKeyBytes}); err != nil {
		return fmt.Errorf("failed to encode private key: %v", err)
	}

	return nil
}
