package gateway

import (
	"database/sql"
	"github.com/fsnotify/fsnotify"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"log/slog"
	"os"
	"sync"
)

// TODO: cron job to sync config in the case of db configuration.
// TODO: add method to sync config.

// Reads from gateway.json file from root directory...
var GlobalProxyConfig model.ProxyConfig
var configLock = &sync.RWMutex{}

// TODO: load proxy config from database
func LoadProxyConfigFromDb(database *sql.DB) {

	slog.Info("Refreshing global proxy gateway configurations from database.")
	serviceRepo := db.NewServiceRepository(database)
	services, err := serviceRepo.GetAllServices()
	if err != nil {
		slog.Info("Error reading services from database during proxy config sync", err)
		// We don't terminate here as we can still run with the existing cached config.
	}

	pluginRepo := db.NewPluginRepository(database)
	globalPlugins, err := pluginRepo.GetPlugins("global", "mock")
	if err != nil {
		slog.Info("Error reading plugins from database during proxy config sync", err)
		// We don't terminate here as we can still run with the existing cached config.
	}

	// Update the global proxy config
	GlobalProxyConfig = model.ProxyConfig{
		Global: model.Global{
			Name:    "someRandomName",
			Plugins: globalPlugins,
		},
		Services: services,
	}
}

func LoadProxyConfigFromConfigFile(filePath string) {

	slog.Info("Loading proxy_pass gateway from config file.")
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		slog.Info("Error reading gateway file", err)
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
		slog.Info(err.Error())
		panic("Error validating gateway file")
	}

	slog.Info("Config file loaded successfully")
	// Validations passed
	GlobalProxyConfig = *config

	// Reset load balancer caches
	Reset()

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
					LoadProxyConfigFromConfigFile(filePath)
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
	slog.Info("Started watching gateway file: " + filePath)

	<-done
}
