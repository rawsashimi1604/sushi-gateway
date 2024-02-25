package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ProxyPort string
}

var GlobalAppConfig *AppConfig

func LoadGlobalConfig() *AppConfig {
	slog.Info("Loading global app config")
	godotenv.Load()

	errors := make([]string, 0)
	proxyPort := os.Getenv("PROXY_PORT")
	if proxyPort == "" {
		errors = append(errors, "PROXY_PORT is required.")
	}

	config := &AppConfig{
		ProxyPort: proxyPort,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
