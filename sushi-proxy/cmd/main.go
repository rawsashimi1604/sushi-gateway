package main

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/api"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"golang.org/x/sync/errgroup"
)

func main() {

	// Load gateway environment config
	loadedConfig, err := gateway.LoadGlobalConfig()
	if err != nil {
		os.Exit(1)
	}
	gateway.GlobalAppConfig = loadedConfig

	// Setup error group with cancellation context
	errGrpCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errGroup, errGrpCtx := errgroup.WithContext(errGrpCtx)

	// Initialize all servers and routers first
	appRouter := gateway.NewRouter()

	// Initialize HTTP server
	httpServer := &http.Server{
		Addr:    ":" + constant.PORT_HTTP,
		Handler: appRouter,
	}

	// Initialize HTTPS server
	cert, err := tls.LoadX509KeyPair(gateway.GlobalAppConfig.ServerCertPath, gateway.GlobalAppConfig.ServerKeyPath)
	if err != nil {
		slog.Error("Failed to load TLS keys", "error", err)
		log.Fatal(err)
	}

	// Load global CA Cert Pool, allowing clients to send CA certificate for authentication via MTLS
	gateway.GlobalCaCertPool = gateway.LoadCertPool()

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    gateway.GlobalCaCertPool.Pool,
		ClientAuth:   tls.RequestClientCert,
	}

	httpsServer := &http.Server{
		Addr:      ":" + constant.PORT_HTTPS,
		Handler:   appRouter,
		TLSConfig: tlsConfig,
	}

	// Initialize Admin API server
	var adminApiRouter http.Handler

	if gateway.GlobalAppConfig.PersistenceConfig == constant.DBLESS_MODE {
		// Init Dbless mode gateway
		slog.Info("PersistenceConfig:: Starting gateway in DB-less mode.")
		adminApiRouter = api.NewAdminApiRouter(nil)
	}

	adminServer := &http.Server{
		Addr:    ":" + constant.PORT_ADMIN_API,
		Handler: adminApiRouter,
	}

	// Initialize config file watcher
	// Do this on gateway startup, load the config from config file
	if err := gateway.LoadProxyConfigFromConfigFile(gateway.GlobalAppConfig.ConfigFilePath); err != nil {
		slog.Error("Failed to load initial config file", "error", err)
		os.Exit(1)
	}

	// Start the file watcher
	errGroup.Go(func() error {
		return gateway.WatchConfigFile(errGrpCtx, gateway.GlobalAppConfig.ConfigFilePath)
	})

	// Start health checker, we start the health checker before the servers start, so that we can verify the health of the services before they are proxied.
	// We also add it to the error group, so that it can be stopped gracefully when the gateway is shutdown.
	gateway.GlobalHealthChecker.Initialize()
	gateway.GlobalHealthChecker.CheckHealthForAllServices() // Initial health check, run it once before starting the ticker
	errGroup.Go(func() error {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		// Periodic health checks
		for {
			select {
			case <-errGrpCtx.Done():
				slog.Info("Stopping health checker...")
				return nil // Stop the health checker by exiting the infinite loop
			case <-ticker.C:
				gateway.GlobalHealthChecker.CheckHealthForAllServices()
			}
		}
	})

	// Start all servers concurrently
	// Start HTTP server
	errGroup.Go(func() error {
		slog.Info("Started sushi-proxy_pass http server on port: " + constant.PORT_HTTP)

		// Graceful shutdown on context cancellation
		go func() {
			<-errGrpCtx.Done()
			httpServer.Shutdown(context.Background())
			slog.Info("Gracefully shutdown HTTP Server....")
		}()

		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("HTTP server failed", "error", err)
			return err
		}
		return nil
	})

	// Start HTTPS server
	errGroup.Go(func() error {
		slog.Info("Started sushi-proxy_pass https server on port: " + constant.PORT_HTTPS)

		// Graceful shutdown on context cancellation
		go func() {
			<-errGrpCtx.Done()
			httpsServer.Shutdown(context.Background())
			slog.Info("Gracefully shutdown HTTPS Server....")
		}()

		if err := httpsServer.ListenAndServeTLS("", ""); err != http.ErrServerClosed {
			slog.Error("HTTPS server failed", "error", err)
			return err
		}
		return nil
	})

	// Start Admin API server
	errGroup.Go(func() error {
		slog.Info("Started admin API server on port: " + constant.PORT_ADMIN_API)

		// Graceful shutdown on context cancellation
		go func() {
			<-errGrpCtx.Done()
			adminServer.Shutdown(context.Background())
			slog.Info("Gracefully shutdown Admin API Server....")
		}()

		if err := adminServer.ListenAndServe(); err != http.ErrServerClosed {
			slog.Error("Admin API server failed", "error", err)
			return err
		}
		return nil
	})

	// Wait for all servers and handle errors
	if err := errGroup.Wait(); err != nil {
		slog.Error("Server error detected, shutting down...", "error", err)
		log.Fatal(err)
	}
}
