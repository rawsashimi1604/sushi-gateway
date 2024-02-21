package main

import (
	"github.com/rawsashimi1604/sushi-gateway/admin-api/internal/router"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("cmd:: Hello from admin api.")
	appRouter := router.NewRouter()
	slog.Info("Started admin-api service on port: 8081")
	log.Fatal(http.ListenAndServe(":8081", appRouter))
}
