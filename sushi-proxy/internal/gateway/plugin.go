package gateway

import (
	"encoding/json"
	"net/http"
)

type PluginPhase string

const (
	AccessPhase PluginPhase = "access"
	LogPhase    PluginPhase = "log"
)

type PluginExecutor interface {
	Execute(next http.Handler) http.Handler
}

type PluginValidation interface {
	Validate() error
}

type Plugin struct {
	Name      string
	Priority  uint
	Phase     PluginPhase
	Handler   PluginExecutor
	Validator PluginValidation
}

// CreatePluginConfigInput Convert the plugin input to how we would expect it to be in the gateway file
func CreatePluginConfigInput(input map[string]interface{}) (map[string]interface{}, error) {
	converted, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	config := make(map[string]interface{})
	err = json.Unmarshal(converted, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
