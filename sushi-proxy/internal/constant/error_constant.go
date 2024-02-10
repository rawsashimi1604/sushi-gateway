package constant

// Error constants in the app.
const (
	CREATE_HAPROXY_REQUEST_ERROR_CODE     = "CREATE_HAPROXY_REQUEST_ERROR"
	CREATE_HAPROXY_REQUEST_ERROR          = "Error creating request to HAProxy: %s"
	FORWARD_HAPROXY_REQUEST_ERROR_CODE    = "FORWARD_HAPROXY_REQUEST_ERROR"
	FORWARD_HAPROXY_REQUEST_ERROR         = "Error forwarding request to HAProxy: %s"
	READ_HAPROXY_RESPONSE_BODY_ERROR_CODE = "READ_HAPROXY_RESPONSE_BODY_ERROR"
	READ_HAPROXY_RESPONSE_BODY_ERROR      = "Error reading response body from HAProxy: %s"
)
