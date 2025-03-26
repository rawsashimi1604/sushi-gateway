package gateway

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerCertPath  string
	ServerKeyPath   string
	CACertPath      string
	AdminUser       string
	AdminPassword   string
	AdminCorsOrigin string
	ConfigFilePath  string
}

var GlobalAppConfig *AppConfig

func LoadGlobalConfig() (*AppConfig, error) {
	slog.Info("Loading Global application config for sushi gateway from environment variables...")
	godotenv.Load()

	errors := make([]string, 0)

	// Get certificate paths from environment
	serverCertPath := os.Getenv("SERVER_CERT_PATH")
	serverKeyPath := os.Getenv("SERVER_KEY_PATH")

	// Check if server cert paths are provided
	hasServerCert := serverCertPath != ""
	hasServerKey := serverKeyPath != ""

	// If any server cert path is provided, both must be provided
	if hasServerCert || hasServerKey {
		// When both server cert and key is provided, it is loaded into our configuration,
		// else it is an user configuration error as they did not pass either cert or key
		if !hasServerCert {
			errors = append(errors, "if you want to use your own certificates, SERVER_CERT_PATH is required when SERVER_KEY_PATH is provided. for auto generating the certificates, leave both SERVER_CERT_PATH and SERVER_KEY_PATH empty")
		}
		if !hasServerKey {
			errors = append(errors, "if you want to use your own certificates, SERVER_KEY_PATH is required when SERVER_CERT_PATH is provided. for auto generating the certificates, leave both SERVER_CERT_PATH and SERVER_KEY_PATH empty")
		}
	} else {
		// If user did not spsecify both the server cert or key, we generate the self signed cert in the gateway on load.
		slog.Info("Since no certs were found, auto generating self signed certs for the TLS server...")
		if err := GenerateSelfSignedCerts("."); err != nil {
			slog.Error("Failed to generate self-signed certificates", "error", err)
			errors = append(errors, "Failed to generate self-signed certificates")
		} else {
			// Set certificate paths
			serverCertPath = filepath.Join(".", "server.crt")
			serverKeyPath = filepath.Join(".", "server.key")
		}
	}

	// Optional, we only need CA Certs for MTLS communications
	// caCertPath := os.Getenv("CA_CERT_PATH")

	// CORS configurations

	// Admin User and Password is used for ADMIN API credentials
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		errors = append(errors, "ADMIN_USER is required.")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		errors = append(errors, "ADMIN_PASSWORD is required.")
	}

	// Admin API Cors configurations, optional, if not provided, set to default localhost sushi manager (localhost:5173)
	adminCorsOrigin := os.Getenv("ADMIN_CORS_ORIGIN")

	// Defines the path to our declarative configuration file to load configurations for the gateway.
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		errors = append(errors, "CONFIG_FILE_PATH is required.")
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		slog.Error("Errors detected when loading environment configuration exiting...")
		return nil, fmt.Errorf("failed to load environment configuration")
	}

	config := &AppConfig{
		ServerCertPath:  serverCertPath,
		ServerKeyPath:   serverKeyPath,
		CACertPath:      caCertPath,
		AdminUser:       adminUser,
		AdminPassword:   adminPassword,
		AdminCorsOrigin: adminCorsOrigin,
		ConfigFilePath:  configFilePath,
	}

	return config, nil
}
