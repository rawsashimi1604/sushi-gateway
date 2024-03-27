package rate_limit

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// Global rate limit stores
var globalRateLimitSecStore = RateLimitStore{
	mu:    sync.Mutex{},
	rates: make(map[string]int),
}

var globalRateLimitMinStore = RateLimitStore{
	mu:    sync.Mutex{},
	rates: make(map[string]int),
}

var globalRateLimitHourStore = RateLimitStore{
	mu:    sync.Mutex{},
	rates: make(map[string]int),
}

// Safe way to store rate limit values, no race condition.
type RateLimitStore struct {
	mu    sync.Mutex
	rates map[string]int
}

type RateLimitPlugin struct {
	config map[string]interface{}
}

func NewRateLimitPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_RATE_LIMIT,
		Priority: 10,
		Handler: RateLimitPlugin{
			config: config,
		},
	}
}

func (plugin RateLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing rate limit function...")

		clientIp := r.RemoteAddr

		// Enable globally...
		defaultLimitPerSecond := 1
		defaultLimitPerMinute := 10
		defaultLimitPerHour := 20

		// Async safe operation
		globalRateLimitSecStore.mu.Lock()
		globalRateLimitMinStore.mu.Lock()
		globalRateLimitHourStore.mu.Lock()

		secCount, secExists := globalRateLimitSecStore.rates[clientIp]
		if !secExists {
			// If no ip hit, create a new entry.
			globalRateLimitSecStore.rates[clientIp] = 1

			// Sec counter.
			go func() {
				time.Sleep(1 * time.Second)
				globalRateLimitSecStore.mu.Lock()
				delete(globalRateLimitSecStore.rates, clientIp)
				globalRateLimitSecStore.mu.Unlock()
			}()
		}

		minCount, minExists := globalRateLimitMinStore.rates[clientIp]
		if !minExists {
			// If no ip hit, create a new entry.
			globalRateLimitMinStore.rates[clientIp] = 1

			// Min counter.
			go func() {
				time.Sleep(1 * time.Minute)
				globalRateLimitMinStore.mu.Lock()
				delete(globalRateLimitMinStore.rates, clientIp)
				globalRateLimitMinStore.mu.Unlock()
			}()
		}

		hourCount, hourExists := globalRateLimitHourStore.rates[clientIp]
		if !hourExists {
			// If no ip hit, create a new entry.
			globalRateLimitHourStore.rates[clientIp] = 1

			// Hour counter.
			go func() {
				time.Sleep(1 * time.Hour)
				globalRateLimitHourStore.mu.Lock()
				delete(globalRateLimitHourStore.rates, clientIp)
				globalRateLimitHourStore.mu.Unlock()
			}()
		}

		globalRateLimitSecStore.rates[clientIp] = secCount + 1
		globalRateLimitMinStore.rates[clientIp] = minCount + 1
		globalRateLimitHourStore.rates[clientIp] = hourCount + 1
		globalRateLimitSecStore.mu.Unlock()
		globalRateLimitMinStore.mu.Unlock()
		globalRateLimitHourStore.mu.Unlock()

		slog.Info(fmt.Sprintf("secCount: %v", secCount))
		slog.Info(fmt.Sprintf("minCount: %v", minCount))
		slog.Info(fmt.Sprintf("hourCount: %v", hourCount))

		if secCount > defaultLimitPerSecond {
			err := errors.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_SECOND_EXCEEDED",
				"Rate limit exceeded.")
			err.WriteJSONResponse(w)
			return
		}

		if minCount > defaultLimitPerMinute {
			err := errors.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_MINUTE_EXCEEDED",
				"Rate limit exceeded.")
			err.WriteJSONResponse(w)
			return
		}

		if hourCount > defaultLimitPerHour {
			err := errors.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_HOUR_EXCEEDED",
				"Rate limit exceeded.")
			err.WriteJSONResponse(w)
			return
		}

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
