package main

import (
	"crypto/tls"
	certificate "github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/cert"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
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
		certificate.GlobalCaCertPool = certificate.LoadCertPool()

		cert, err := tls.LoadX509KeyPair(config.GlobalAppConfig.ServerCertPath, config.GlobalAppConfig.ServerKeyPath)
		if err != nil {
			log.Fatalf("server: loadkeys: %s", err)
		}

		// allow clients to send cert for mtls validation
		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cert},
			ClientCAs:    certificate.GlobalCaCertPool.Pool,
			ClientAuth:   tls.RequestClientCert,
		}

		server := &http.Server{
			Addr:      ":" + constant.PORT_HTTPS,
			Handler:   appRouter,
			TLSConfig: tlsConfig,
		}
		slog.Info("Started sushi-proxy_pass https server on port: " + constant.PORT_HTTPS)
		log.Fatal(server.ListenAndServeTLS("", "")) // Certs loaded from tls config.
	}()

	// Block forever
	select {}
}
