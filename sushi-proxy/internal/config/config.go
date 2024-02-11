package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ProxyPort            string
	ReverseProxyHttpUrl  string
	ReverseProxyHttpsUrl string
}

var GlobalAppConfig *AppConfig

func LoadConfig() *AppConfig {
	slog.Info("Loading configurations from environment")
	godotenv.Load()

	errors := make([]string, 0)
	proxyPort := os.Getenv("PROXY_PORT")
	if proxyPort == "" {
		errors = append(errors, "PROXY_PORT is required.")
	}

	revProxyHttpUrl := os.Getenv("REVERSE_PROXY_HTTP_URL")
	if revProxyHttpUrl == "" {
		errors = append(errors, "REVERSE_PROXY_HTTP_URL is required.")
	}

	revProxyHttpsUrl := os.Getenv("REVERSE_PROXY_HTTPS_URL")
	if revProxyHttpsUrl == "" {
		errors = append(errors, "REVERSE_PROXY_HTTPS_URL is required.")
	}

	config := &AppConfig{
		ProxyPort:            proxyPort,
		ReverseProxyHttpUrl:  revProxyHttpUrl,
		ReverseProxyHttpsUrl: revProxyHttpsUrl,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
