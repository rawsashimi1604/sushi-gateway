package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"io/ioutil"
	"log/slog"
	"sync"
)

// TODO: refactor errors, add error codes.
// TODO: add global param for config file path
// TODO: add validation for config file
// Reads from config.json file from root directory...
type ProxyConfig struct {
	Global struct {
		Name    string                `json:"name"`
		Plugins []models.PluginConfig `json:"plugins"` // Adjusted to use the Plugin struct
	} `json:"global"`
	Services []models.Service `json:"services"`
}

var GlobalProxyConfig ProxyConfig
var configLock = &sync.RWMutex{}

func LoadProxyConfig(filePath string) {
	slog.Info("Loading proxy_pass config")
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		slog.Info("Error reading config file", err)
		panic("Error reading config file")
	}

	// Validate the config file, if valid -> assign to GlobalProxyConfig
	configLock.Lock()
	config, err := validateAndParseSchema(configFile)
	if err != nil {
		panic("Error parsing config file")
	}

	err = validateConfig(config)
	if err != nil {
		panic("Error validating config file")
	}

	slog.Info("Config file loaded successfully")
	// Validations passed
	GlobalProxyConfig = *config
	configLock.Unlock()

}

func WatchConfigFile(filePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		slog.Info("Error creating watcher: %v", err)
		panic("Error creating watcher")
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write != 0 {
					LoadProxyConfig(filePath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				slog.Info("Filesystem watcher error: " + err.Error())
				panic("Filesystem watcher error")
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		slog.Info("Error adding watcher to file: %v", err)
		panic("Error adding watcher to file")
	}
	slog.Info("Started watching config file: " + filePath)

	<-done
}
