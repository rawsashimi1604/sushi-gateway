package request_size_limit

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type RequestSizeLimitPlugin struct{}

var Plugin = NewRequestSizeLimitPlugin()

// In bytes
const maxRequestSize = 10

func NewRequestSizeLimitPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "request_size_limit",
		Priority: 10,
		Handler:  RequestSizeLimitPlugin{},
	}
}

func (plugin RequestSizeLimitPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing request_size_limit function...")

		err := checkRequestLength(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkRequestLength(r *http.Request) *errors.HttpError {
	if r.ContentLength > maxRequestSize {
		slog.Info(fmt.Sprintf("Request size too large: %vB", r.ContentLength))
		return errors.NewHttpError(http.StatusRequestEntityTooLarge,
			"REQUEST_TOO_LARGE", "Request size too large.")
	}

	return nil
}
