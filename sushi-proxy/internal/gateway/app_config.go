package gateway

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ConfigFilePath string
	ServerCertPath string
	ServerKeyPath  string
	CACertPath     string
	AdminUser      string
	AdminPassword  string
}

var GlobalAppConfig *AppConfig

func LoadGlobalConfig() *AppConfig {
	slog.Info("Loading global app gateway")
	godotenv.Load()

	errors := make([]string, 0)

	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		errors = append(errors, "CONFIG_FILE_PATH is required.")
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

	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		errors = append(errors, "ADMIN_USER is required.")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		errors = append(errors, "ADMIN_PASSWORD is required.")
	}

	config := &AppConfig{
		ConfigFilePath: configFilePath,
		ServerCertPath: serverCertPath,
		ServerKeyPath:  serverKeyPath,
		CACertPath:     caCertPath,
		AdminUser:      adminUser,
		AdminPassword:  adminPassword,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
