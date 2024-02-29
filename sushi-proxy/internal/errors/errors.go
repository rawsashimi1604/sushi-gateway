package errors

type HttpError struct {
	Code     string // Internal error code
	Message  string // Human-readable error message
	HttpCode int    // Http error code if applicable.
}

func (e *HttpError) Error() string {
	return e.Message
}
