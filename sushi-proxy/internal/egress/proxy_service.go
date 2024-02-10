package egress

import (
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type ProxyService struct {
	HAProxyURL string
}

func NewProxyService(haProxyURL string) *ProxyService {
	return &ProxyService{
		HAProxyURL: haProxyURL,
	}
}

func (p *ProxyService) ForwardRequest(req *http.Request) ([]byte, int, error) {
	client := &http.Client{}

	// Parse the HAProxy URL
	parsedURL, err := url.Parse(p.HAProxyURL)
	if err != nil {
		slog.Error("Error parsing HAProxy URL: ", err)
		return nil, 0, err
	}

	// Create a new request to send to HAProxy
	proxyReq, err := http.NewRequest(req.Method, parsedURL.String(), req.Body)
	if err != nil {
		slog.Error("Error creating request to HAProxy: ", err)
		return nil, 0, err
	}

	// Copy headers from the original request
	proxyReq.Header = req.Header

	// Forward the request to HAProxy
	resp, err := client.Do(proxyReq)
	if err != nil {
		slog.Error("Error forwarding request to HAProxy: ", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body from HAProxy: ", err)
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
