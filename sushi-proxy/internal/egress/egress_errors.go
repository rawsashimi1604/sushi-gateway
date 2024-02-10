package egress

type EgressError struct {
	Code     string // Internal error code
	Message  string // Human-readable error message
	HttpCode int    // Http error code if applicable.
}

func (e *EgressError) Error() string {
	return e.Message
}
