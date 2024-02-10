package main

import (
	"github.com/rawsashimi1604/sushi-gateway/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/internal/router"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	config.Config = config.LoadConfig()
	appRouter := router.NewRouter()
	slog.Info("Started sushi-proxy service!")
	log.Fatal(http.ListenAndServe(":8008", appRouter))
}
