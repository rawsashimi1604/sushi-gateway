package config

import (
	"encoding/json"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"log/slog"
	"sync"
)

// TODO: refactor errors
// TODO: add global param for config file path
// Reads from config.json file from root directory...
type ProxyConfig struct {
	Global struct {
		Name    string        `json:"name"`
		Plugins []interface{} `json:"plugins"` // If you have a specific plugin struct, replace interface{} with that
	} `json:"global"`
	Upstreams []struct {
		Name    string        `json:"name"`
		Host    string        `json:"host"`
		Port    int           `json:"port"`
		Plugins []interface{} `json:"plugins"` // If you have a specific plugin struct, replace interface{} with that
		Routes  []struct {
			Path     string        `json:"path"`
			Upstream string        `json:"upstream"`
			Plugins  []interface{} `json:"plugins"` // If you have a specific plugin struct, replace interface{} with that
		} `json:"routes"`
	} `json:"upstreams"`
}

var GlobalProxyConfig ProxyConfig
var configLock = &sync.RWMutex{}

func LoadProxyConfig(filePath string) {
	slog.Info("Loading proxy config")
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		slog.Info("Error reading config file", err)
		panic("Error reading config file")
	}

	configLock.Lock()
	err = json.Unmarshal(configFile, &GlobalProxyConfig)
	configLock.Unlock()

	if err != nil {
		slog.Info("Error parsing config file", err)
		panic("Error parsing config file")
	}
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
				slog.Info("Filesystem watcher error: %v", err)
				panic("Filesystem watcher error")
			}
		}
	}()

	err = watcher.Add(filePath)
	if err != nil {
		slog.Info("Error adding watcher to file: %v", err)
		panic("Error adding watcher to file")
	}
	slog.Info("Started watching config file:", filePath)

	<-done
}
