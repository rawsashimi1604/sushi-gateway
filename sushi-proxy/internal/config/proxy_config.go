package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/models"
	"log/slog"
	"os"
	"sync"
)

// Reads from config.json file from root directory...
var GlobalProxyConfig models.ProxyConfig
var configLock = &sync.RWMutex{}

func LoadProxyConfig(filePath string) {
	slog.Info("Loading proxy_pass config")
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		slog.Info("Error reading config file", err)
		panic("Error reading config file")
	}

	// Validate the config file, if valid -> assign to GlobalProxyConfig
	configLock.Lock()
	config, err := ValidateAndParseSchema(configFile)
	if err != nil {
		panic("Error parsing config file")
	}

	err = ValidateConfig(config)
	if err != nil {
		slog.Info(err.Error())
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
