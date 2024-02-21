package error

import (
	"encoding/json"
	"net/http"
)

// HttpError represents an error that returns a json message with code.
type HttpError struct {
	HttpCode int    `json:"httpCode"`
	Code     string `json:"code,omitempty"`
	Message  string `json:"message"`
}

// GenericError represents an error that happens in the application.
type GenericError struct {
	Code    string
	Message string
}

func NewHttpError(httpCode int, code string, message string) *HttpError {
	return &HttpError{
		HttpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}

func NewGenericError(code string, message string) *GenericError {
	return &GenericError{
		Code:    code,
		Message: message,
	}
}

// WriteJSONResponse writes the error as a JSON response to the client.
func (e *HttpError) WriteJSONResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(e.HttpCode) // Set the HTTP status code

	// Encode the PluginError as JSON and write to the response
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"error": e,
	})
	if err != nil {
		http.Error(w, "Failed to write error response", http.StatusInternalServerError)
	}
}
