package gateway

import (
	"crypto/x509"
	"log"
	"log/slog"
	"os"
)

var GlobalCaCertPool *CertPool

type CertPool struct {
	Pool *x509.CertPool
}

func LoadCertPool() *CertPool {
	// Load CA certificate to create a CA pool
	// We didnt find any CA Certs provided to the gateway, returning an empty cert pool.
	if GlobalAppConfig.CACertPath == "" {
		slog.Info("CA_CERT_PATH not defined, no certs found. Skip loading client certificates.")
		return &CertPool{}
	}

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
