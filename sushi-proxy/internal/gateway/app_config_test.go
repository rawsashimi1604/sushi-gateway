package gateway

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTestEnv(t *testing.T) func() {
	// Save original env vars
	originalEnv := make(map[string]string)
	envVars := []string{
		"SERVER_CERT_PATH",
		"SERVER_KEY_PATH",
		"CA_CERT_PATH",
		"ADMIN_USER",
		"ADMIN_PASSWORD",
		"ADMIN_CORS_ORIGIN",
		"CONFIG_FILE_PATH",
	}

	for _, env := range envVars {
		originalEnv[env] = os.Getenv(env)
		os.Unsetenv(env)
	}

	// Return cleanup function
	return func() {
		for key, value := range originalEnv {
			if value != "" {
				os.Setenv(key, value)
			} else {
				os.Unsetenv(key)
			}
		}
	}
}

func TestLoadGlobalConfig(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	// Set required environment variables
	os.Setenv("ADMIN_USER", "admin")
	os.Setenv("ADMIN_PASSWORD", "password")
	os.Setenv("CONFIG_FILE_PATH", "config.json")

	config, err := LoadGlobalConfig()

	assert.NotNil(t, config)
	assert.NoError(t, err)
	assert.Equal(t, "admin", config.AdminUser)
	assert.Equal(t, "password", config.AdminPassword)
	assert.Equal(t, "config.json", config.ConfigFilePath)
	assert.Equal(t, filepath.Join(".", "server.crt"), config.ServerCertPath)
	assert.Equal(t, filepath.Join(".", "server.key"), config.ServerKeyPath)
}

func TestLoadGlobalConfig_WithCustomCerts(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	// Set required environment variables including custom certs
	os.Setenv("ADMIN_USER", "admin")
	os.Setenv("ADMIN_PASSWORD", "password")
	os.Setenv("CONFIG_FILE_PATH", "config.json")
	os.Setenv("SERVER_CERT_PATH", "/custom/cert.pem")
	os.Setenv("SERVER_KEY_PATH", "/custom/key.pem")
	os.Setenv("CA_CERT_PATH", "/custom/ca.pem")

	config, err := LoadGlobalConfig()

	assert.NotNil(t, config)
	assert.NoError(t, err)
	assert.Equal(t, "/custom/cert.pem", config.ServerCertPath)
	assert.Equal(t, "/custom/key.pem", config.ServerKeyPath)
	assert.Equal(t, "/custom/ca.pem", config.CACertPath)
}

func TestLoadGlobalConfig_FailingCustomCerts(t *testing.T) {
	cleanup := setupTestEnv(t)
	defer cleanup()

	// Set required environment variables but with incomplete cert configuration
	os.Setenv("ADMIN_USER", "admin")
	os.Setenv("ADMIN_PASSWORD", "password")
	os.Setenv("CONFIG_FILE_PATH", "config.json")
	os.Setenv("SERVER_CERT_PATH", "/custom/cert.pem")
	// Deliberately omit SERVER_KEY_PATH to trigger error

	_, err := LoadGlobalConfig()
	assert.Error(t, err)

	// Reset and try the opposite case
	cleanup()
	os.Setenv("ADMIN_USER", "admin")
	os.Setenv("ADMIN_PASSWORD", "password")
	os.Setenv("CONFIG_FILE_PATH", "config.json")
	os.Setenv("SERVER_KEY_PATH", "/custom/key.pem")
	// Deliberately omit SERVER_CERT_PATH to trigger error

	_, err = LoadGlobalConfig()
	assert.Error(t, err)
}
