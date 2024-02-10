package main

import (
	"github.com/rawsashimi1604/sushi-gateway/internal"
	"log"
	"log/slog"
	"net/http"
)

func main() {

	// Run checks to see if api req is present in egress
	// Run through auth
	// Delegate to plugin runner
	// Run through other plugins
	internal.LoadConfig()
	router := internal.NewRouter()
	slog.Info("Started sushi-proxy service!")
	log.Fatal(http.ListenAndServe(":8008", router))
}
