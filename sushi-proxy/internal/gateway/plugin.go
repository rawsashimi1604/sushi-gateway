package gateway

import (
	"encoding/json"
	"net/http"
)

type PluginExecutor interface {
	Execute(next http.Handler) http.Handler
}

type Plugin struct {
	Name     string
	Priority uint
	Handler  PluginExecutor
}

// CreatePluginConfigJsonInput Convert the plugin input to how we would expect it to be in the gateway file
func CreatePluginConfigJsonInput(input map[string]interface{}) (map[string]interface{}, error) {
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