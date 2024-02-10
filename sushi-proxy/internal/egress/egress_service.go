package egress

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/internal/constant"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type EgressService struct {
	HAProxyURL string
}

func NewEgressService(haProxyURL string) *EgressService {
	return &EgressService{
		HAProxyURL: haProxyURL,
	}
}

func (s *EgressService) ForwardRequest(req *http.Request) ([]byte, int, *EgressError) {
	client := &http.Client{}

	// Parse the HAProxy URL
	parsedURL, err := url.Parse(s.HAProxyURL)
	if err != nil {
		return nil, 0, &EgressError{
			Code:    constant.PARSE_HAPROXY_URL_ERROR_CODE,
			Message: fmt.Sprint(constant.PARSE_HAPROXY_URL_ERROR, err),
		}
	}
	parsedURL = parsedURL.JoinPath(req.URL.Path)

	// Create a new request to send to HAProxy
	proxyReq, err := http.NewRequest(req.Method, parsedURL.String(), req.Body)
	if err != nil {
		return nil, 0, &EgressError{
			Code:    constant.CREATE_HAPROXY_REQUEST_ERROR_CODE,
			Message: fmt.Sprint(constant.CREATE_HAPROXY_REQUEST_ERROR, err),
		}
	}

	// Copy headers from the original request
	proxyReq.Header = req.Header

	// Forward the request to HAProxy
	resp, err := client.Do(proxyReq)
	if err != nil {
		return nil, 0, &EgressError{
			Code:    constant.FORWARD_HAPROXY_REQUEST_ERROR_CODE,
			Message: fmt.Sprint(constant.FORWARD_HAPROXY_REQUEST_ERROR, err),
		}
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body from HAProxy: ", err)
		return nil, resp.StatusCode, &EgressError{
			Code:    constant.READ_HAPROXY_RESPONSE_BODY_ERROR_CODE,
			Message: constant.READ_HAPROXY_RESPONSE_BODY_ERROR,
		}
	}

	return body, resp.StatusCode, nil
}
