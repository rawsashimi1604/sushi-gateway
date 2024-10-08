package gateway

type ProxyConfig struct {
	Global struct {
		Name    string         `json:"name"`
		Plugins []PluginConfig `json:"plugins"` // Adjusted to use the Plugin struct
	} `json:"global"`
	Services []Service `json:"services"`
}

type PluginConfig map[string]interface{}

type Upstream struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Route struct {
	Name    string         `json:"name"`
	Path    string         `json:"path"`
	Methods []string       `json:"methods"`
	Plugins []PluginConfig `json:"plugins"` // Adjusted to use the Plugin struct
}

type Service struct {
	Name                  string                 `json:"name"`
	BasePath              string                 `json:"base_path"`
	Protocol              string                 `json:"protocol"`
	LoadBalancingStrategy LoadBalancingAlgorithm `json:"load_balancing_strategy"`
	Upstreams             []Upstream             `json:"upstreams"`
	Plugins               []PluginConfig         `json:"plugins"` // Adjusted to use the Plugin struct
	Routes                []Route                `json:"routes"`
}

type Credentials struct {
	Plugin  string                 `json:"plugin"`
	Enabled bool                   `json:"enabled"`
	Data    map[string]interface{} `json:"data"`
}
