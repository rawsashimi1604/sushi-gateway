package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ProxyPortHttp  string
	ProxyPortHttps string
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

	config := &AppConfig{
		ProxyPortHttp:  proxyPortHttp,
		ProxyPortHttps: proxyPortHttps,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
