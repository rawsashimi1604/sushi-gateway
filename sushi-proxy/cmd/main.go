package main

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/router"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	// Load configs
	configPath := "./config.json"
	config.GlobalAppConfig = config.LoadGlobalConfig()
	config.LoadProxyConfig(configPath)

	go config.WatchConfigFile(configPath)

	appRouter := router.NewRouter()
	slog.Info("Started sushi-proxy_pass service on port: " + config.GlobalAppConfig.ProxyPort)
	log.Fatal(http.ListenAndServe(":"+config.GlobalAppConfig.ProxyPort, appRouter))
}
