package config

import (
	"encoding/json"
	"log/slog"
)

func validateAndParseSchema(raw []byte) (*ProxyConfig, error) {
	var config ProxyConfig
	err := json.Unmarshal(raw, &config)
	if err != nil {
		slog.Info("Error parsing config file", err)
		return nil, err
	}

	return &config, nil
}

func validateProxyConfig(config *ProxyConfig) error {

	return nil
}
