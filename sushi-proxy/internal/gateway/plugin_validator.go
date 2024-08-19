package gateway

// Contains all logic related to validating plugin configurations.
type PluginValidator struct {
}

func NewPluginValidator() *PluginValidator {
	return &PluginValidator{}
}

func (pv *PluginValidator) ValidatePluginConfig(config PluginConfig) bool {
	// TODO: complete plugin validation architecture logic....
	return true
}
