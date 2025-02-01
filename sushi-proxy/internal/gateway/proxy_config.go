package gateway

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// Reads from declarative config file
var GlobalProxyConfig model.ProxyConfig
var configLock = &sync.RWMutex{}

func LoadProxyConfigFromConfigFile(filePath string) error {
	slog.Info("Loading proxy_pass gateway from config file.")
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("Error reading gateway file", "error", err)
		return err
	}

	// Validate the gateway file, if valid -> assign to GlobalProxyConfig
	configLock.Lock()
	defer configLock.Unlock()

	config, err := ValidateAndParseSchema(configFile)
	if err != nil {
		slog.Error("Error parsing gateway file", "error", err)
		return err
	}

	err = ValidateConfig(config)
	if err != nil {
		slog.Error("Error validating gateway file", "error", err)
		return err
	}

	slog.Info("Config file loaded successfully")
	// Validations passed
	GlobalProxyConfig = *config

	// Reset load balancer caches
	ResetLoadBalancers()

	return nil
}

func WatchConfigFile(ctx context.Context, filePath string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		slog.Error("Error creating watcher for config file", "error", err)
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(filePath); err != nil {
		slog.Error("Error adding config file to watcher", "error", err)
		return err
	}

	for {
		select {
		case <-ctx.Done():
			slog.Info("Config file watcher shutting down...")
			return nil
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if event.Op&fsnotify.Write != 0 {
				if err := LoadProxyConfigFromConfigFile(filePath); err != nil {
					slog.Error("Failed to load config file", "error", err)
					return err
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			slog.Error("Filesystem watcher error", "error", err)
			return err
		}
	}
}
