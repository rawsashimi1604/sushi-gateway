package main

import (
	"github.com/rawsashimi1604/sushi-gateway/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/internal/router"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	config.GlobalAppConfig = config.LoadConfig()
	appRouter := router.NewRouter()
	slog.Info("Started sushi-proxy service on port: " + config.GlobalAppConfig.ProxyPort)
	log.Fatal(http.ListenAndServe(":"+config.GlobalAppConfig.ProxyPort, appRouter))
}
