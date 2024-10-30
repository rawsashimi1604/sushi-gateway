package gateway

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"log/slog"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerCertPath    string
	ServerKeyPath     string
	CACertPath        string
	AdminUser         string
	AdminPassword     string
	PersistenceConfig string
	ConfigFilePath    string
	DbConnectionHost  string
	DbConnectionName  string
	DbConnectionUser  string
	DbConnectionPass  string
	DbConnectionPort  string
}

var GlobalAppConfig *AppConfig

func LoadGlobalConfig() *AppConfig {
	slog.Info("Loading global app gateway")
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

	adminUser := os.Getenv("ADMIN_USER")
	if adminUser == "" {
		errors = append(errors, "ADMIN_USER is required.")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		errors = append(errors, "ADMIN_PASSWORD is required.")
	}

	persistenceConfig := os.Getenv("PERSISTENCE_CONFIG")
	if persistenceConfig == "" {
		errors = append(errors, "PERSISTENCE_CONFIG is required.")
	}
	if strings.ToLower(persistenceConfig) != constant.DB_MODE &&
		strings.ToLower(persistenceConfig) != constant.DBLESS_MODE {
		errors = append(errors,
			"PERSISTENCE_CONFIG must be \"db\" or \"dbless\".")
	}

	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if persistenceConfig == constant.DBLESS_MODE && configFilePath == "" {
		errors = append(errors, "CONFIG_FILE_PATH is required.")
	}

	dbConnectionHost := os.Getenv("DB_CONNECTION_HOST")
	dbConnectionName := os.Getenv("DB_CONNECTION_NAME")
	dbConnectionUser := os.Getenv("DB_CONNECTION_USER")
	dbConnectionPass := os.Getenv("DB_CONNECTION_PASS")
	dbConnectionPort := os.Getenv("DB_CONNECTION_PORT")

	if persistenceConfig != "" {
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

	config := &AppConfig{
		ServerCertPath:    serverCertPath,
		ServerKeyPath:     serverKeyPath,
		CACertPath:        caCertPath,
		AdminUser:         adminUser,
		AdminPassword:     adminPassword,
		PersistenceConfig: persistenceConfig,
		ConfigFilePath:    configFilePath,
		DbConnectionHost:  dbConnectionHost,
		DbConnectionName:  dbConnectionName,
		DbConnectionUser:  dbConnectionUser,
		DbConnectionPass:  dbConnectionPass,
		DbConnectionPort:  dbConnectionPort,
	}

	if len(errors) > 0 {
		for _, err := range errors {
			slog.Error(err)
		}
		panic("Errors detected when loading environment configuration...")
	}

	return config
}
