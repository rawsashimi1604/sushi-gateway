package internal

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	REVERSE_PROXY_HTTP_URL  string
	REVERSE_PROXY_HTTPS_URL string
}

func LoadConfig() (*AppConfig, error) {
	godotenv.Load()

	revProxyHttpUrl := os.Getenv("REVERSE_PROXY_HTTP_URL")
	if revProxyHttpUrl == "" {
		return nil, errors.New("REVERSE_PROXY_HTTP_URL is required")
	}

	revProxyHttpsUrl := os.Getenv("REVERSE_PROXY_HTTPS_URL")
	if revProxyHttpsUrl == "" {
		return nil, errors.New("REVERSE_PROXY_HTTPS_URL is required")
	}

	config := &AppConfig{
		REVERSE_PROXY_HTTP_URL:  revProxyHttpUrl,
		REVERSE_PROXY_HTTPS_URL: revProxyHttpsUrl,
	}

	return config, nil
}
