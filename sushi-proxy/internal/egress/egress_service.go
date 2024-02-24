package egress

import (
	"bytes"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type EgressService struct {
}

func NewEgressService() *EgressService {
	return &EgressService{}
}

// captureResponseWriter is used to capture the HTTP response
type captureResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func newCaptureResponseWriter(w http.ResponseWriter) *captureResponseWriter {
	// Default the status code to 200 in case WriteHeader is not called
	return &captureResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (crw *captureResponseWriter) WriteHeader(code int) {
	crw.statusCode = code
	crw.ResponseWriter.WriteHeader(code)
}

func (crw *captureResponseWriter) Write(data []byte) (int, error) {
	crw.body.Write(data)
	return crw.ResponseWriter.Write(data)
}

func (s *EgressService) HandleProxyPass(w http.ResponseWriter, req *http.Request) ([]byte, int, *EgressError) {
	// TODO: shift url parsing logic to a separate function
	// TODO: add path detection and parsing
	target, err := url.Parse("http://localhost:8001")
	if err != nil {
		return nil, 0, &EgressError{
			Code:     "ERROR_CREATING_PROXY",
			Message:  "Error parsing URL when creating proxy_pass",
			HttpCode: http.StatusInternalServerError,
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	captureWriter := newCaptureResponseWriter(w)

	// Modify the request as needed
	req.URL.Host = target.Host
	req.URL.Scheme = target.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = target.Host

	// Serve the proxy and capture the response
	proxy.ServeHTTP(captureWriter, req)

	// After the proxy serves, you can access the captured response.
	return captureWriter.body.Bytes(), captureWriter.statusCode, nil
}
