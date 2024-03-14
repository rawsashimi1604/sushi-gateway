package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ProxyPortHttp  string
	ProxyPortHttps string
	ServerCertPath string
	ServerKeyPath  string
	CACertPath     string
}

var GlobalAppConfig *AppConfig

func LoadGlobalConfig() *AppConfig {
	slog.Info("Loading global app config")
	godotenv.Load()

	errors := make([]string, 0)
	proxyPortHttp := os.Getenv("PROXY_PORT_HTTP")
	if proxyPortHttp == "" {
		errors = append(errors, "PROXY_PORT_HTTP is required.")
	}

	proxyPortHttps := os.Getenv("PROXY_PORT_HTTPS")
	if proxyPortHttps == "" {
		errors = append(errors, "PROXY_PORT_HTTPS is required.")
	}

	serverCertPath := os.Getenv("SERVER_CERT_PATH")
	if serverCertPath == "" {
		errors = append(errors, "SERVER_CERT_PATH is required.")
	}

	serverKeyPath := os.Getenv("SERVER_KEY_PATH")
	if serverKeyPath == "" {
		errors = append(errors, "SERVER_KEY_PATH is required.")
	}

	caCertPath := os.Getenv("CA_CERT_PATH")
	if caCertPath == "" {
		errors = append(errors, "CA_CERT_PATH is required.")
	}

	config := &AppConfig{
		ProxyPortHttp:  proxyPortHttp,
		ProxyPortHttps: proxyPortHttps,
		ServerCertPath: serverCertPath,
		ServerKeyPath:  serverKeyPath,
		CACertPath:     caCertPath,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
