package egress

import (
	"bytes"
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
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

func (s *EgressService) HandleProxyPass(w http.ResponseWriter, req *http.Request) ([]byte, int, *errors.HttpError) {

	path, convertErr := s.convertPathToProxyPassUrl(req)
	if convertErr != nil {
		return nil, 0, convertErr
	}

	slog.Info("path: " + path)
	target, err := url.Parse(path)
	if err != nil {
		return nil, 0, &errors.HttpError{
			Code:     "ERROR_PARSING_PROXY_URL",
			Message:  "Error parsing URL when creating proxy_pass",
			HttpCode: http.StatusInternalServerError,
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	captureWriter := newCaptureResponseWriter(w)

	// Customize the Director to modify request before forwarding
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		// Call the original Director to preserve other behaviors
		originalDirector(req)

		// Now adjust req.URL.Path here as needed
		req.URL.Path = target.Path
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Header.Set(constant.X_FORWARDED_HOST, req.Header.Get("Host"))
		req.Header.Set(constant.X_FORWARDED_FOR, req.RemoteAddr)
		req.Host = target.Host
	}

	// Serve the proxy and capture the response
	proxy.ServeHTTP(captureWriter, req)

	return captureWriter.body.Bytes(), captureWriter.statusCode, nil
}

func (s *EgressService) convertPathToProxyPassUrl(req *http.Request) (string, *errors.HttpError) {
	matchedService, matchedRoute, err := util.GetServiceAndRouteFromRequest(req)
	if err != nil {
		return "", err
	}

	// Assuming the use of the first upstream for simplicity
	if len(matchedService.Upstreams) == 0 {
		return "", &errors.HttpError{
			Code:     "NO_UPSTREAMS",
			Message:  "No upstreams found for the matched service",
			HttpCode: http.StatusInternalServerError,
		}
	}

	// Use first upstream for now, configure load balancing next time.
	upstream := matchedService.Upstreams[0]
	proxyURL := fmt.Sprintf("%s://%s:%d%s", matchedService.Protocol, upstream.Host, upstream.Port, matchedRoute.Path)
	return proxyURL, nil
}
