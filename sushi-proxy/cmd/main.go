package main

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/router"
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

	// Setup http server
	go func() {
		slog.Info("Started sushi-proxy_pass http server on port: " + config.GlobalAppConfig.ProxyPortHttp)
		if err := http.ListenAndServe(":"+config.GlobalAppConfig.ProxyPortHttp, appRouter); err != nil {
			slog.Info("Failed to start HTTP server: %v", err)
			panic(err)
		}
	}()

	// Setup https server
	go func() {
		slog.Info("Started sushi-proxy_pass https server on port: " + config.GlobalAppConfig.ProxyPortHttps)
		if err := http.ListenAndServeTLS(":"+config.GlobalAppConfig.ProxyPortHttps,
			config.GlobalAppConfig.TLSCertPath, config.GlobalAppConfig.TLSKeyPath, appRouter); err != nil {
			slog.Info("Failed to start HTTPS server: %v", err)
			panic(err)
		}
	}()

	// Block forever
	select {}
}
