package gateway

import (
	"crypto/x509"
	"log"
	"os"
)

var GlobalCaCertPool *CertPool

type CertPool struct {
	Pool *x509.CertPool
}

func LoadCertPool() *CertPool {
	// Load CA certificate to create a CA pool
	caCert, err := os.ReadFile(GlobalAppConfig.CACertPath)
	if err != nil {
		log.Fatalf("server: read ca: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return &CertPool{
		Pool: caCertPool,
	}
}
