package main

import (
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/api"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
)

func main() {
	gateway.GlobalAppConfig = gateway.LoadGlobalConfig()

	// DB MODE run a thread to sync the config file from db.
	if gateway.GlobalAppConfig.PersistenceConfig == constant.DB_MODE {
		database, err := db.ConnectDb()
		if err != nil {
			panic("unable to connect to database.")
		}

		// Load the initial config file.
		gateway.LoadProxyConfigFromDb(database)

		// On first load, if the config file is empty, means that the config is not in sync...
		// Means we could not get the config from the db
		// We should therefore terminate the application.
		if gateway.GlobalProxyConfig.Global.Name == "" {
			slog.Info("Unable to sync config from database on first load. Terminating...")
			panic("unable to sync config from database.")
		}

		// Thereafter we should run a cron job to sync the config from the db.
		gateway.StartProxyConfigCronJob(database,
			gateway.GlobalAppConfig.PersistenceSyncInterval)
	}

	// DB LESS MODE run a thread to monitor the config file for changes and do an initial boot up...
	if gateway.GlobalAppConfig.PersistenceConfig == constant.DBLESS_MODE {
		gateway.LoadProxyConfigFromConfigFile(gateway.GlobalAppConfig.ConfigFilePath)
		go gateway.WatchConfigFile(gateway.GlobalAppConfig.ConfigFilePath)
	}

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
		var adminApiRouter http.Handler
		if gateway.GlobalAppConfig.PersistenceConfig == constant.DB_MODE {
			database, err := db.ConnectDb()
			if err != nil {
				panic("unable to connect to database.")
			}
			slog.Info("PersistenceConfig:: Starting gateway in DB mode.")
			adminApiRouter = api.NewAdminApiRouter(database)
		} else {
			slog.Info("PersistenceConfig:: Starting gateway in DB-less mode.")
			adminApiRouter = api.NewAdminApiRouter(nil)
		}

		slog.Info("Started another API server on port: " + constant.PORT_ADMIN_API)
		if err := http.ListenAndServe(":"+constant.PORT_ADMIN_API, adminApiRouter); err != nil {
			slog.Info("Failed to start new API server: %v", err)
			log.Fatal(err)
		}
	}()

	// Block forever
	select {}
}
