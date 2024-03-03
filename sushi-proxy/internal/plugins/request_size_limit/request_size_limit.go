package request_size_limit

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type RequestSizeLimitPlugin struct {
	config map[string]interface{}
}

func NewRequestSizeLimitPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_REQUEST_SIZE_LIMIT,
		Priority: 10,
		Handler: RequestSizeLimitPlugin{
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

func (plugin RequestSizeLimitPlugin) checkRequestLength(r *http.Request) *errors.HttpError {
	data := plugin.config["data"].(map[string]interface{})
	maxPayloadSize := data["max_payload_size"].(float64)

	if r.ContentLength > int64(maxPayloadSize) {
		slog.Info(fmt.Sprintf("Request size too large: %vB", r.ContentLength))
		return errors.NewHttpError(http.StatusRequestEntityTooLarge,
			"REQUEST_TOO_LARGE", "Request size too large.")
	}

	return nil
}
