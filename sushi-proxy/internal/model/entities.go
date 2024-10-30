package model

// Domain Model of objects in Sushi Gateway
type ProxyConfig struct {
	Global   Global    `json:"global"`
	Services []Service `json:"services"`
}
type Global struct {
	Name    string         `json:"name"`
	Plugins []PluginConfig `json:"plugins"`
}

type PluginConfig struct {
	Id      string                 `json:"id"`
	Name    string                 `json:"name"`
	Config  map[string]interface{} `json:"config"`
	Enabled bool                   `json:"enabled"`
}

type Upstream struct {
	Id   string `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Route struct {
	Name    string         `json:"name"`
	Path    string         `json:"path"`
	Methods []string       `json:"methods"`
	Plugins []PluginConfig `json:"plugins"`
}

type Service struct {
	Name                  string                 `json:"name"`
	BasePath              string                 `json:"base_path"`
	Protocol              string                 `json:"protocol"`
	LoadBalancingStrategy LoadBalancingAlgorithm `json:"load_balancing_strategy"`
	Upstreams             []Upstream             `json:"upstreams"`
	Plugins               []PluginConfig         `json:"plugins"`
	Routes                []Route                `json:"routes"`
}
