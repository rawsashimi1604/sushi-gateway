package gateway

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
)

// This plugin is not configurable
// Used specially to capture response metadata in the Response Phase
type ResponseHandlerPlugin struct {
	config map[string]interface{}
}

func NewResponseHandlerPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_RESPONSE_HANDLER,
		Priority: 10000,
		Phase:    ResponsePhase,
		Handler: ResponseHandlerPlugin{
			config: config,
		},
		Validator: ResponseHandlerPlugin{
			config: config,
		},
	}
}

func (plugin ResponseHandlerPlugin) Validate() error {
	return nil
}

func (plugin ResponseHandlerPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing response handler function...")
		next.ServeHTTP(w, r)

		// Get the capture writer injected from config
		captureWriter, _ := plugin.config["capture_writer"].(*captureResponseWriter)

		// After response is sent, add metadata to the request context
		// Add response metadata to request context after handling completes
		ctx := r.Context()
		ctx = context.WithValue(ctx, constant.CONTEXT_RESPONSE_HEADERS, captureWriter.Header())
		ctx = context.WithValue(ctx, constant.CONTEXT_RESPONSE_SIZE, captureWriter.size)
		ctx = context.WithValue(ctx, constant.CONTEXT_RESPONSE_STATUS, captureWriter.statusCode)
		ctx = context.WithValue(ctx, constant.CONTEXT_END_TIME, time.Now())
		*r = *r.WithContext(ctx)
	})
}
