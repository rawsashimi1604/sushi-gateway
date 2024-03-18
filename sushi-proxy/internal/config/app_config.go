package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerCertPath string
	ServerKeyPath  string
	CACertPath     string
}

var GlobalAppConfig *AppConfig

func LoadGlobalConfig() *AppConfig {
	slog.Info("Loading global app config")
	godotenv.Load()

	errors := make([]string, 0)

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
