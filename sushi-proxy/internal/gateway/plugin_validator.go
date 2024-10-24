package gateway

import "github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"

// Contains all logic related to validating plugin configurations.
type PluginValidator struct {
}

func NewPluginValidator() *PluginValidator {
	return &PluginValidator{}
}

func (pv *PluginValidator) ValidatePluginConfig(config model.PluginConfig) bool {
	// TODO: complete plugin validation architecture logic....
	return true
}
