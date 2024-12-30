package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
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
		Priority: 2000,
		Handler: CorsPlugin{
			config: config,
		},
		Validator: CorsPlugin{
			config: config,
		},
	}
}

func (plugin CorsPlugin) Validate() error {
	// Validate allow_origins
	origins, ok := plugin.config["allow_origins"].([]interface{})
	if !ok {
		return fmt.Errorf("allow_origins must be an array of strings")
	}
	if len(origins) == 0 {
		return fmt.Errorf("allow_origins cannot be empty")
	}
	for _, origin := range origins {
		if _, ok := origin.(string); !ok {
			return fmt.Errorf("allow_origins must contain only strings")
		}
	}

	// Validate allow_methods if present
	if methods, exists := plugin.config["allow_methods"].([]interface{}); exists {
		if len(methods) == 0 {
			return fmt.Errorf("allow_methods cannot be empty if specified")
		}
		validMethods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"}
		for _, method := range methods {
			methodStr, ok := method.(string)
			if !ok {
				return fmt.Errorf("allow_methods must contain only strings")
			}
			if !util.SliceContainsString(validMethods, strings.ToUpper(methodStr)) {
				return fmt.Errorf("invalid HTTP method: %s", methodStr)
			}
		}
	}

	// Validate allow_headers if present
	if headers, exists := plugin.config["allow_headers"].([]interface{}); exists {
		if len(headers) == 0 {
			return fmt.Errorf("allow_headers cannot be empty if specified")
		}
		for _, header := range headers {
			if _, ok := header.(string); !ok {
				return fmt.Errorf("allow_headers must contain only strings")
			}
		}
	}

	// Validate expose_headers if present
	if exposeHeaders, exists := plugin.config["expose_headers"].([]interface{}); exists {
		if len(exposeHeaders) == 0 {
			return fmt.Errorf("expose_headers cannot be empty if specified")
		}
		for _, header := range exposeHeaders {
			if _, ok := header.(string); !ok {
				return fmt.Errorf("expose_headers must contain only strings")
			}
		}
	}

	// Validate allow_credentials if present
	if _, exists := plugin.config["allow_credentials"].(bool); !exists && plugin.config["allow_credentials"] != nil {
		return fmt.Errorf("allow_credentials must be a boolean")
	}

	// Validate allow_private_network if present
	if _, exists := plugin.config["allow_private_network"].(bool); !exists && plugin.config["allow_private_network"] != nil {
		return fmt.Errorf("allow_private_network must be a boolean")
	}

	// Validate preflight_continue if present
	if _, exists := plugin.config["preflight_continue"].(bool); !exists && plugin.config["preflight_continue"] != nil {
		return fmt.Errorf("preflight_continue must be a boolean")
	}

	// Validate max_age if present
	if maxAge, exists := plugin.config["max_age"].(float64); exists {
		if maxAge < 0 {
			return fmt.Errorf("max_age cannot be negative")
		}
	} else if plugin.config["max_age"] != nil {
		return fmt.Errorf("max_age must be a number")
	}

	return nil
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
