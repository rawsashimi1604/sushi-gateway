package util

import "encoding/json"

// CreatePluginConfigJsonInput Convert the plugin input to how we would expect it to be in the config file
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
