package constant

// TODO: prevent collisions of context keys by using iota
// Context keys used throughout the application
const (
	CONTEXT_CAPTURE_WRITER = "capture_writer"

	// StartTime represents the start time of a request
	CONTEXT_START_TIME = "start_time"
	CONTEXT_END_TIME   = "end_time"

	// HTTP Response
	CONTEXT_RESPONSE_SIZE    = "response_size"
	CONTEXT_RESPONSE_STATUS  = "response_status"
	CONTEXT_RESPONSE_HEADERS = "response_headers"
)
