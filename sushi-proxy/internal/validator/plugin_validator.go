package validator

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

// TODO: move validation logic to model.

type PluginValidator struct {
}

func NewPluginValidator() *PluginValidator {
	return &PluginValidator{}
}

func (pv *PluginValidator) ValidatePlugin(plugin model.PluginConfig) error {

	// TODO: validate each plugin data/gateway schema
	if err := pv.validateNameExists(plugin); err != nil {
		return err
	}
	if err := pv.validateAvailablePlugin(plugin); err != nil {
		return err
	}

	// Pass validation
	return nil
}

func (pv *PluginValidator) validateNameExists(plugin model.PluginConfig) error {
	if plugin.Name == "" {
		return fmt.Errorf("plugin name is required")
	}

	return nil
}

func (pv *PluginValidator) validateAvailablePlugin(plugin model.PluginConfig) error {
	if !util.SliceContainsString(constant.AVAILABLE_PLUGINS, plugin.Name) {
		return fmt.Errorf("plugin name is invalid. "+
			"Available plugins: %v", constant.AVAILABLE_PLUGINS)
	}
	return nil
}
