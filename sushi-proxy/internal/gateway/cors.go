package gateway

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type CORSConfig struct {
	AllowOrigins        []string
	AllowMethods        []string
	AllowHeaders        []string
	ExposeHeaders       []string
	AllowCredentials    bool
	AllowPrivateNetwork bool
	PreflightContinue   bool
	MaxAge              int64
}

type CorsPlugin struct {
	config map[string]interface{}
}

func NewCorsPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_CORS,
		Priority: 20,
		Handler:  CorsPlugin{config: config},
	}
}

func (plugin CorsPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Convert gateway map to CORSConfig struct. Assume gateway has been validated.
		corsConfig := plugin.parseCORSConfig()

		slog.Info("Executing CORS function...")

		// Set CORS headers
		if len(corsConfig.AllowOrigins) > 0 {
			w.Header().Set("Access-Control-Allow-Origin", strings.Join(corsConfig.AllowOrigins, ","))
		}
		if len(corsConfig.AllowMethods) > 0 {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowMethods, ","))
		}
		if len(corsConfig.AllowHeaders) > 0 {
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowHeaders, ","))
		}
		if len(corsConfig.ExposeHeaders) > 0 {
			w.Header().Set("Access-Control-Expose-Headers", strings.Join(corsConfig.ExposeHeaders, ","))
		}
		if corsConfig.AllowCredentials {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		if corsConfig.MaxAge > 0 {
			w.Header().Set("Access-Control-Max-Age", strconv.FormatInt(corsConfig.MaxAge, 10))
		}

		// Handle preflight (OPTIONS) requests, don't proxy the request.
		if r.Method == http.MethodOptions && !corsConfig.PreflightContinue {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (plugin CorsPlugin) parseCORSConfig() CORSConfig {

	config := plugin.config

	corsConfig := CORSConfig{
		AllowOrigins:        util.ToStringSlice(config["allow_origins"].([]interface{})),
		AllowMethods:        util.ToStringSlice(config["allow_methods"].([]interface{})),
		AllowHeaders:        util.ToStringSlice(config["allow_headers"].([]interface{})),
		ExposeHeaders:       util.ToStringSlice(config["expose_headers"].([]interface{})),
		AllowCredentials:    config["allow_credentials"].(bool),
		AllowPrivateNetwork: config["allow_private_network"].(bool),
		PreflightContinue:   config["preflight_continue"].(bool),
		MaxAge:              int64(config["max_age"].(float64)),
	}

	return corsConfig
}
