package gateway

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
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

func (s *EgressService) HandleProxyPass(w http.ResponseWriter, req *http.Request) *HttpError {

	path, convertErr := s.convertPathToProxyPassUrl(req)
	if convertErr != nil {
		return convertErr
	}

	slog.Info("path: " + path)
	target, err := url.Parse(path)
	if err != nil {
		return &HttpError{
			Code:     "ERROR_PARSING_PROXY_URL",
			Message:  "Error parsing URL when creating proxy_pass",
			HttpCode: http.StatusInternalServerError,
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Customize the Director to modify request before forwarding
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		// Call the original Director to preserve other behaviors
		originalDirector(req)

		req.URL.Path = target.Path
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Header.Set(constant.X_FORWARDED_HOST, req.Header.Get("Host"))
		req.Header.Set(constant.X_FORWARDED_FOR, req.RemoteAddr)
		req.Host = target.Host
	}
	proxy.ServeHTTP(w, req)
	return nil
}

func (s *EgressService) convertPathToProxyPassUrl(req *http.Request) (string, *HttpError) {
	matchedService, matchedRoute, err := GetServiceAndRouteFromRequest(&GlobalProxyConfig, req)
	if err != nil {
		return "", err
	}

	// Assuming the use of the first upstream for simplicity
	if len(matchedService.Upstreams) == 0 {
		return "", &HttpError{
			Code:     "NO_UPSTREAMS",
			Message:  "No upstreams found for the matched service",
			HttpCode: http.StatusInternalServerError,
		}
	}

	loadBalancer := NewLoadBalancer()
	// TODO: check which algorithm for routing. For now, use round robin
	upstreamIndex := loadBalancer.GetNextUpstream(RoundRobin, *matchedService)
	upstream := matchedService.Upstreams[upstreamIndex]
	proxyURL := fmt.Sprintf("%s://%s:%d%s", matchedService.Protocol, upstream.Host, upstream.Port, matchedRoute.Path)
	return proxyURL, nil
}
