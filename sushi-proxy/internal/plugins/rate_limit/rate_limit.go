package rate_limit

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

// Global rate limit store
var globalRateLimitMinStore = RateLimitStore{
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
		defaultLimitPerMinute := 10

		// Async safe operation
		globalRateLimitMinStore.mu.Lock()
		count, exists := globalRateLimitMinStore.rates[clientIp]
		// If no ip hit, create a new entry.
		if !exists {
			globalRateLimitMinStore.rates[clientIp] = 1

			// Start the counter to reset after 1 minute.
			go func() {
				time.Sleep(1 * time.Minute)
				globalRateLimitMinStore.mu.Lock()
				delete(globalRateLimitMinStore.rates, clientIp)
				globalRateLimitMinStore.mu.Unlock()
			}()
		}
		globalRateLimitMinStore.rates[clientIp] = count + 1
		globalRateLimitMinStore.mu.Unlock()

		if count > defaultLimitPerMinute {
			err := errors.NewHttpError(http.StatusTooManyRequests,
				"RATE_LIMIT_MINUTE_EXCEEDED",
				"Rate limit exceeded.")
			err.WriteJSONResponse(w)
			return
		}

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
