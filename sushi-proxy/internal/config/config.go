package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ReverseProxyHttpUrl  string
	ReverseProxyHttpsUrl string
}

var Config *AppConfig

func LoadConfig() *AppConfig {
	slog.Info("Loading configurations from environment")
	godotenv.Load()

	errors := make([]string, 0)
	revProxyHttpUrl := os.Getenv("REVERSE_PROXY_HTTP_URL")
	if revProxyHttpUrl == "" {
		errors = append(errors, "REVERSE_PROXY_HTTP_URL is required.")
	}

	revProxyHttpsUrl := os.Getenv("REVERSE_PROXY_HTTPS_URL")
	if revProxyHttpsUrl == "" {
		errors = append(errors, "REVERSE_PROXY_HTTPS_URL is required.")
	}

	config := &AppConfig{
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
