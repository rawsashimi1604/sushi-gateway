package plugins

import (
	"encoding/json"
	"net/http"
)

// PluginError represents an error that can occur within a plugin.
type PluginError struct {
	HttpCode int    `json:"httpCode"`
	Code     string `json:"code,omitempty"`
	Message  string `json:"message"`
}

// NewPluginError creates a new instance of PluginError.
func NewPluginError(httpCode int, code string, message string) *PluginError {
	return &PluginError{
		HttpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}

// WriteJSONResponse writes the error as a JSON response to the client.
func (e *PluginError) WriteJSONResponse(w http.ResponseWriter) {
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
