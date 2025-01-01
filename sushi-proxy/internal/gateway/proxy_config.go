package gateway

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// TODO: cron job to sync config in the case of db configuration.
// TODO: add method to sync config.

// Reads from gateway.json file from root directory...
var GlobalProxyConfig model.ProxyConfig
var configLock = &sync.RWMutex{}

func LoadProxyConfigFromDb(database *sql.DB) {

	slog.Info("Refreshing global proxy gateway configurations from database.")
	serviceRepo := db.NewServiceRepository(database)
	services, err := serviceRepo.GetAllServices()
	if err != nil {
		slog.Info("Error reading services from database during proxy config sync", "error", err)
		// We don't terminate here as we can still run with the existing cached config.
	}

	gatewayRepo := db.NewGatewayRepository(database)
	gatewayGlobalConfig, err := gatewayRepo.GetGatewayInfo()
	if err != nil {
		slog.Info("Error reading gateway global config from database during proxy config sync", "error", err)
		// We don't terminate here as we can still run with the existing cached config.
	}

	// Update the global proxy config
	GlobalProxyConfig = model.ProxyConfig{
		Global:   gatewayGlobalConfig,
		Services: services,
	}
}

func LoadProxyConfigFromConfigFile(filePath string) {

	slog.Info("Loading proxy_pass gateway from config file.")
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		slog.Info("Error reading gateway file", "error", err)
		panic("Error reading gateway file")
	}

	// Validate the gateway file, if valid -> assign to GlobalProxyConfig
	configLock.Lock()
	config, err := ValidateAndParseSchema(configFile)
	if err != nil {
		panic("Error parsing gateway file")
	}

	err = ValidateConfig(config)
	if err != nil {
		slog.Info(err.Error(), "error", err)
		panic("Error validating gateway file")
	}

	slog.Info("Config file loaded successfully")
	// Validations passed
	GlobalProxyConfig = *config

	// Reset load balancer caches
	Reset()

	configLock.Unlock()

}

func StartProxyConfigCronJob(database *sql.DB, interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				LoadProxyConfigFromDb(database)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func WatchConfigFile(filePath string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		slog.Error("Error creating watcher", "error", err)
		return err
	}
	defer watcher.Close()

	if err := watcher.Add(filePath); err != nil {
		slog.Error("Error adding file to watcher", "error", err)
		return err
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if event.Op&fsnotify.Write != 0 {
				LoadProxyConfigFromConfigFile(filePath)
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
