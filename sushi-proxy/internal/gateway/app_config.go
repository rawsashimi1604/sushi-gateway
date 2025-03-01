package gateway

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerCertPath          string
	ServerKeyPath           string
	CACertPath              string
	AdminUser               string
	AdminPassword           string
	PersistenceConfig       string
	PersistenceSyncInterval int
	ConfigFilePath          string
	DbConnectionHost        string
	DbConnectionName        string
	DbConnectionUser        string
	DbConnectionPass        string
	DbConnectionPort        string
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

	// Admin User and Password is used for ADMIN API credentials
	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		errors = append(errors, "ADMIN_USER is required.")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		errors = append(errors, "ADMIN_PASSWORD is required.")
	}

	// Persistence Config must be "db" or "dbless", it specifies the persistence mode for the gateway.
	persistenceConfig := os.Getenv("PERSISTENCE_CONFIG")
	if persistenceConfig == "" {
		errors = append(errors, "PERSISTENCE_CONFIG is required.")
	}
	if strings.ToLower(persistenceConfig) != constant.DB_MODE &&
		strings.ToLower(persistenceConfig) != constant.DBLESS_MODE {
		errors = append(errors,
			"PERSISTENCE_CONFIG must be \"db\" or \"dbless\".")
	}

	// Sync Interval defines how often we sync with the database in seconds. Only required for db mode
	var syncIntervalInteger int
	persistenceSyncInterval := os.Getenv("PERSISTENCE_SYNC_INTERVAL")
	if persistenceConfig == constant.DB_MODE {
		if persistenceSyncInterval == "" {
			errors = append(errors, "PERSISTENCE_SYNC_INTERVAL is required.")
		}
		if val, err := strconv.Atoi(persistenceSyncInterval); err != nil {
			syncIntervalInteger = 0
			errors = append(errors, "PERSISTENCE_SYNC_INTERVAL must be a valid integer.")
		} else {
			syncIntervalInteger = val
		}
	}

	// Defines database connection variables, only needed in db mode.
	dbConnectionHost := os.Getenv("DB_CONNECTION_HOST")
	dbConnectionName := os.Getenv("DB_CONNECTION_NAME")
	dbConnectionUser := os.Getenv("DB_CONNECTION_USER")
	dbConnectionPass := os.Getenv("DB_CONNECTION_PASS")
	dbConnectionPort := os.Getenv("DB_CONNECTION_PORT")

	if persistenceConfig == constant.DB_MODE {
		if dbConnectionHost == "" {
			errors = append(errors, "DB_CONNECTION_HOST is required.")
		}
		if dbConnectionName == "" {
			errors = append(errors, "DB_CONNECTION_NAME is required.")
		}
		if dbConnectionUser == "" {
			errors = append(errors, "DB_CONNECTION_USER is required.")
		}
		if dbConnectionPass == "" {
			errors = append(errors, "DB_CONNECTION_PASS is required.")
		}
		if dbConnectionPort == "" {
			errors = append(errors, "DB_CONNECTION_PORT is required.")
		}
	}

	// Defines the path to our declarative configuration file, only required in dbless mode.
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if persistenceConfig == constant.DBLESS_MODE && configFilePath == "" {
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
		ServerCertPath: serverCertPath,
		ServerKeyPath:  serverKeyPath,
		// CACertPath:              caCertPath,
		AdminUser:               adminUser,
		AdminPassword:           adminPassword,
		PersistenceConfig:       persistenceConfig,
		PersistenceSyncInterval: syncIntervalInteger,
		ConfigFilePath:          configFilePath,
		DbConnectionHost:        dbConnectionHost,
		DbConnectionName:        dbConnectionName,
		DbConnectionUser:        dbConnectionUser,
		DbConnectionPass:        dbConnectionPass,
		DbConnectionPort:        dbConnectionPort,
	}

	return config, nil
}
