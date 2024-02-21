package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	PostgresUrl string
}

var GlobalAppConfig *AppConfig

func LoadConfig() *AppConfig {
	slog.Info("Loading configurations from environment")
	godotenv.Load()

	errors := make([]string, 0)
	postgresUrl := os.Getenv("POSTGRES_URL")
	if postgresUrl == "" {
		errors = append(errors, "POSTGRES_URL is required.")
	}

	config := &AppConfig{
		PostgresUrl: postgresUrl,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
