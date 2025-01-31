package gateway

import (
	"encoding/json"
	"net/http"
)

/*
	We execute plugins in different phases.
	The access phase is executed before the request is proxied to the upstream API.
		- Plugins that need to ensure access to the upstream API such as authentication and authorization can be included here.
	The log phase is executed after the request is proxied to the upstream API.
		- Plugins that need to be executed regardless of the request outcome can be included here. This is useful for logging or metrics plugins.
	The response phase is executed after the response is received from the upstream API.
		- It is mainly used for logging the response metadata.
*/

type PluginPhase string

const (
	AccessPhase   PluginPhase = "access"
	ResponsePhase PluginPhase = "response"
	LogPhase      PluginPhase = "log"
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
