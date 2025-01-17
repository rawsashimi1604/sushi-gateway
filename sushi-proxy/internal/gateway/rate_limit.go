package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"golang.org/x/time/rate"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

// IPRateLimiter holds the rate limiter for an IP address
type IPRateLimiter struct {
	mu         sync.RWMutex
	limiterSec *rate.Limiter
	limiterMin *rate.Limiter
	limiterHr  *rate.Limiter
}

// RateLimitStore stores rate limiters for different scopes, that points to ip
type RateLimitStore struct {
	limits map[string]map[string]*IPRateLimiter // scope -> ip -> limiter
}

// Global rate limit stores
var globalRateLimitStore = &RateLimitStore{
	limits: make(map[string]map[string]*IPRateLimiter),
}

type RateLimitPlugin struct {
	config      map[string]interface{}
	proxyConfig *model.ProxyConfig
}

// getLimiter retrieves or creates a rate limiter for an IP address
func (s *RateLimitStore) getLimiter(scope, ip string, secLimit, minLimit, hrLimit float64) *IPRateLimiter {
	if s.limits[scope] == nil {
		s.limits[scope] = make(map[string]*IPRateLimiter)
	}

	if limiter, exists := s.limits[scope][ip]; exists {
		return limiter
	}

	limiter := &IPRateLimiter{
		limiterSec: rate.NewLimiter(rate.Limit(secLimit), 1),                // per second
		limiterMin: rate.NewLimiter(rate.Limit(minLimit/60), int(minLimit)), // per minute
		limiterHr:  rate.NewLimiter(rate.Limit(hrLimit/3600), int(hrLimit)), // per hour
	}
	s.limits[scope][ip] = limiter
	return limiter
}

func NewRateLimitPlugin(config map[string]interface{}, proxyConfig *model.ProxyConfig) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_RATE_LIMIT,
		Priority: 910,
		Handler: RateLimitPlugin{
			config:      config,
			proxyConfig: proxyConfig,
		},
		Validator: RateLimitPlugin{
			config: config,
		},
	}
}

func (plugin RateLimitPlugin) Validate() error {
	limitSec, ok := plugin.config["limit_second"].(float64)
	if !ok {
		return fmt.Errorf("limit_second must be a number")
	}
	if limitSec <= 0 {
		return fmt.Errorf("limit_second must be greater than 0")
	}

	limitMin, ok := plugin.config["limit_min"].(float64)
	if !ok {
		return fmt.Errorf("limit_min must be a number")
	}
	if limitMin <= 0 {
		return fmt.Errorf("limit_min must be greater than 0")
	}

	limitHour, ok := plugin.config["limit_hour"].(float64)
	if !ok {
		return fmt.Errorf("limit_hour must be a number")
	}
	if limitHour <= 0 {
		return fmt.Errorf("limit_hour must be greater than 0")
	}

	return nil
}

func (plugin RateLimitPlugin) detectRateLimitOperationLevel(service *model.Service, route *model.Route) string {
	// Check whether global, service or route level rate limit.
	for _, servicePlugin := range service.Plugins {
		name := servicePlugin.Name
		if name == constant.PLUGIN_RATE_LIMIT {
			return "Service"
		}
	}

	for _, routePlugin := range route.Plugins {
		name := routePlugin.Name
		if name == constant.PLUGIN_RATE_LIMIT {
			return "Route"
		}
	}

	return "Global"
}

func (plugin RateLimitPlugin) getMapKeyEntry(configLevel string, service *model.Service, route *model.Route) string {
	if configLevel == "Global" {
		return "Global"
	} else if configLevel == "Service" {
		return fmt.Sprintf("Service::%s", service.Name)
	} else {
		return fmt.Sprintf("Route::%s", route.Name)
	}

}

// Execute implementation
func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")

		service, route, err := util.GetServiceAndRouteFromRequest(plugin.proxyConfig, r)
		if err != nil {
			err.WriteLogMessage()
			err.WriteJSONResponse(w)
			return
		}

		rateLimitOperationLevel := plugin.detectRateLimitOperationLevel(service, route)
		clientIp, err := util.GetHostIp(r.RemoteAddr)
		if err != nil {
			err.WriteLogMessage()
			err.WriteJSONResponse(w)
			return
		}

		// Get rate limits from config
		limitSec := plugin.config["limit_second"].(float64)
		limitMin := plugin.config["limit_min"].(float64)
		limitHour := plugin.config["limit_hour"].(float64)

		// Get scope key
		scope := plugin.getMapKeyEntry(rateLimitOperationLevel, service, route)

		// Get or create limiter for this IP
		limiter := globalRateLimitStore.getLimiter(scope, clientIp, limitSec, limitMin, limitHour)

		// Guard against race conditions
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		slog.Info("Rate limiting for IP: " + clientIp)
		slog.Info("Remaining seconds: " + fmt.Sprintf("%f", limiter.limiterSec.Tokens()))
		slog.Info("Remaining minutes: " + fmt.Sprintf("%f", limiter.limiterMin.Tokens()))
		slog.Info("Remaining hours: " + fmt.Sprintf("%f", limiter.limiterHr.Tokens()))

		// Check all limits
		if !limiter.limiterSec.Allow() {
			err := model.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_SECOND_EXCEEDED",
				fmt.Sprintf("Rate limit exceeded for %s (per second)", scope))
			err.WriteLogMessage()
			err.WriteJSONResponse(w)
			return
		}

		if !limiter.limiterMin.Allow() {
			err := model.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_MINUTE_EXCEEDED",
				fmt.Sprintf("Rate limit exceeded for %s (per minute)", scope))
			err.WriteLogMessage()
			err.WriteJSONResponse(w)
			return
		}

		if !limiter.limiterHr.Allow() {
			err := model.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_HOUR_EXCEEDED",
				fmt.Sprintf("Rate limit exceeded for %s (per hour)", scope))
			err.WriteLogMessage()
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
