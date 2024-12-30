package gateway

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

type RequestSizeLimitPlugin struct {
	config map[string]interface{}
}

func NewRequestSizeLimitPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_REQUEST_SIZE_LIMIT,
		Priority: 951,
		Handler: RequestSizeLimitPlugin{
			config: config,
		},
		Validator: RequestSizeLimitPlugin{
			config: config,
		},
	}
}

func (plugin RequestSizeLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing request_size_limit function...")

		err := plugin.checkRequestLength(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (plugin RequestSizeLimitPlugin) checkRequestLength(r *http.Request) *model.HttpError {
	config := plugin.config
	maxPayloadSize := config["max_payload_size"].(float64)

	if r.ContentLength > int64(maxPayloadSize) {
		slog.Info(fmt.Sprintf("Request size too large: %vB", r.ContentLength))
		return model.NewHttpError(http.StatusRequestEntityTooLarge,
			"REQUEST_TOO_LARGE", "Request size too large.")
	}

	return nil
}

func (plugin RequestSizeLimitPlugin) Validate() error {
	maxSize, ok := plugin.config["max_payload_size"].(float64)
	if !ok {
		return fmt.Errorf("max_payload_size must be a number")
	}

	if maxSize <= 0 {
		return fmt.Errorf("max_payload_size must be greater than 0")
	}

	return nil
}
