package gateway

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HttpError struct {
	Code     string // Internal error code
	Message  string // Human-readable error message
	HttpCode int    // Http error code if applicable.
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(httpCode int, code string, message string) *HttpError {
	return &HttpError{
		HttpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}

// WriteJSONResponse writes the error as a JSON response to the client.
func (e *HttpError) WriteJSONResponse(w http.ResponseWriter) {
	w.WriteHeader(e.HttpCode)
	bytes, err := json.Marshal(map[string]interface{}{
		"error": e,
	})

	if err != nil {
		http.Error(w, "Failed to write error response", http.StatusInternalServerError)
	}
	w.Write(bytes)
}

func (e *HttpError) WriteLogMessage() {
	slog.Info("Error code: " + e.Code + ", message: " + e.Message)
}
