package gateway

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// Global rate limit stores
var globalRateLimitSecStore = RateLimitStore{
	mu:    sync.Mutex{},
	rates: make(map[string]map[string]int),
}

var globalRateLimitMinStore = RateLimitStore{
	mu:    sync.Mutex{},
	rates: make(map[string]map[string]int),
}

var globalRateLimitHourStore = RateLimitStore{
	mu:    sync.Mutex{},
	rates: make(map[string]map[string]int),
}

type RateLimitStore struct {
	mu    sync.Mutex
	rates map[string]map[string]int
}

type RateLimitPlugin struct {
	config      map[string]interface{}
	proxyConfig *ProxyConfig
}

func NewRateLimitPlugin(config map[string]interface{}, proxyConfig *ProxyConfig) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_RATE_LIMIT,
		Priority: 10,
		Handler: RateLimitPlugin{
			config:      config,
			proxyConfig: proxyConfig,
		},
	}
}

func (plugin RateLimitPlugin) detectRateLimitOperationLevel(service *Service, route *Route, r *http.Request) (string, *HttpError) {
	// Check whether global, service or route level rate limit.
	for _, servicePlugin := range service.Plugins {
		name := servicePlugin.Name
		if name == constant.PLUGIN_RATE_LIMIT {
			return "service", nil
		}
	}

	for _, routePlugin := range route.Plugins {
		name := routePlugin.Name
		if name == constant.PLUGIN_RATE_LIMIT {
			return "route", nil
		}
	}

	return "global", nil
}

func (plugin RateLimitPlugin) getMapKeyEntry(configLevel string, service *Service, route *Route) string {
	if configLevel == "global" {
		return "global"
	}

	if configLevel == "service" {
		return service.Name
	}

	if configLevel == "route" {
		return fmt.Sprintf("%s_%s", service.Name, route.Name)
	}

	return "global"
}

func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")

		service, route, err := GetServiceAndRouteFromRequest(plugin.proxyConfig, r)
		rateLimitOperationLevel, err := plugin.detectRateLimitOperationLevel(service, route, r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		clientIp := r.RemoteAddr

		// Get proxy configs
		config := plugin.config
		limitSec := int64(config["limit_second"].(float64))
		limitMin := int64(config["limit_min"].(float64))
		limitHour := int64(config["limit_hour"].(float64))

		// Async safe operation
		globalRateLimitSecStore.mu.Lock()
		globalRateLimitMinStore.mu.Lock()
		globalRateLimitHourStore.mu.Lock()

		mapEntry := plugin.getMapKeyEntry(rateLimitOperationLevel, service, route)
		secCount, secExists := globalRateLimitSecStore.rates[mapEntry][clientIp]
		if !secExists {
			// If no ip hit, create a new entry, depending on the plugin configuration level. )
			globalRateLimitSecStore.rates[mapEntry] = make(map[string]int)
			globalRateLimitSecStore.rates[mapEntry][clientIp] = 1

			// Sec counter.
			go func() {
				time.Sleep(1 * time.Second)
				globalRateLimitSecStore.mu.Lock()
				delete(globalRateLimitSecStore.rates[mapEntry], clientIp)
				globalRateLimitSecStore.mu.Unlock()
			}()
		}

		minCount, minExists := globalRateLimitMinStore.rates[mapEntry][clientIp]
		if !minExists {
			// If no ip hit, create a new entry.
			globalRateLimitMinStore.rates[mapEntry] = make(map[string]int)
			globalRateLimitMinStore.rates[mapEntry][clientIp] = 1

			// Min counter.
			go func() {
				time.Sleep(1 * time.Minute)
				globalRateLimitMinStore.mu.Lock()
				delete(globalRateLimitMinStore.rates[mapEntry], clientIp)
				globalRateLimitMinStore.mu.Unlock()
			}()
		}

		hourCount, hourExists := globalRateLimitHourStore.rates[mapEntry][clientIp]
		if !hourExists {
			// If no ip hit, create a new entry.
			globalRateLimitHourStore.rates[mapEntry] = make(map[string]int)
			globalRateLimitHourStore.rates[mapEntry][clientIp] = 1

			// Hour counter.
			go func() {
				time.Sleep(1 * time.Hour)
				globalRateLimitHourStore.mu.Lock()
				delete(globalRateLimitHourStore.rates[mapEntry], clientIp)
				globalRateLimitHourStore.mu.Unlock()
			}()
		}

		globalRateLimitSecStore.rates[mapEntry][clientIp] = secCount + 1
		globalRateLimitMinStore.rates[mapEntry][clientIp] = minCount + 1
		globalRateLimitHourStore.rates[mapEntry][clientIp] = hourCount + 1
		globalRateLimitSecStore.mu.Unlock()
		globalRateLimitMinStore.mu.Unlock()
		globalRateLimitHourStore.mu.Unlock()

		if int64(secCount) > limitSec {
			err := NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_SECOND_EXCEEDED",
				"Rate limit exceeded for "+mapEntry)
			err.WriteJSONResponse(w)
			return
		}

		if int64(minCount) > limitMin {
			err := NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_MINUTE_EXCEEDED",
				"Rate limit exceeded for "+mapEntry)
			err.WriteJSONResponse(w)
			return
		}

		if int64(hourCount) > limitHour {
			err := NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_HOUR_EXCEEDED",
				"Rate limit exceeded for "+mapEntry)
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
