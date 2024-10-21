package main

import (
	"crypto/tls"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/api"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"log"
	"log/slog"
	"net/http"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
)

func main() {
	gateway.GlobalAppConfig = gateway.LoadGlobalConfig()
	gateway.LoadProxyConfig(gateway.GlobalAppConfig.ConfigFilePath)
	go gateway.WatchConfigFile(gateway.GlobalAppConfig.ConfigFilePath)

	appRouter := gateway.NewRouter()

	// Setup http server
	go func() {
		slog.Info("Started sushi-proxy_pass http server on port: " + constant.PORT_HTTP)
		if err := http.ListenAndServe(":"+constant.PORT_HTTP, appRouter); err != nil {
			slog.Info("Failed to start HTTP server: %v", err)
			panic(err)
		}
	}()

	// Setup https server
	go func() {
		// Load global CA Cert Pool
		gateway.GlobalCaCertPool = gateway.LoadCertPool()

		cert, err := tls.LoadX509KeyPair(gateway.GlobalAppConfig.ServerCertPath, gateway.GlobalAppConfig.ServerKeyPath)
		if err != nil {
			log.Fatalf("server: loadkeys: %s", err)
		}

		// allow clients to send cert for mtls validation
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientCAs:    gateway.GlobalCaCertPool.Pool,
			ClientAuth:   tls.RequestClientCert,
		}

		server := &http.Server{
			Addr:      ":" + constant.PORT_HTTPS,
			Handler:   appRouter,
			TLSConfig: tlsConfig,
		}
		slog.Info("Started sushi-proxy_pass https server on port: " + constant.PORT_HTTPS)
		log.Fatal(server.ListenAndServeTLS("", "")) // Certs loaded from tls gateway.
	}()

	// Setup admin api
	go func() {
		adminApiRouter := api.NewAdminApiRouter()

		slog.Info("Started another API server on port: " + constant.PORT_ADMIN_API)
		if err := http.ListenAndServe(":"+constant.PORT_ADMIN_API, adminApiRouter); err != nil {
			slog.Info("Failed to start new API server: %v", err)
			log.Fatal(err)
		}
	}()

	db.ConnectDb()
	
	// Block forever
	select {}
}
