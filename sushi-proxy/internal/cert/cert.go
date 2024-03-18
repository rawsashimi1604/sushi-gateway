package cert

import (
	"crypto/x509"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"io/ioutil"
	"log"
	"os"
)

var GlobalCaCertPool *CertPool

type CertPool struct {
	Pool *x509.CertPool
}

func LoadCertPool() *CertPool {
	// Load CA certificate to create a CA pool
	caCert, err := os.ReadFile((config.GlobalAppConfig.CACertPath)
	if err != nil {
		log.Fatalf("server: read ca: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return &CertPool{
		Pool: caCertPool,
	}
}
