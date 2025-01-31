package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

type HttpLogPlugin struct {
	config map[string]interface{}
}

// TODO: send logs in batches (batch processing).
// TODO: log the response as well.
type HttpLogConfig struct {
	httpEndpoint string
	method       string
	contentType  string
}

func NewHttpLogPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_HTTP_LOG,
		Priority: 12,
		Phase:    LogPhase,
		Handler: HttpLogPlugin{
			config: config,
		},
		Validator: HttpLogPlugin{
			config: config,
		},
	}
}

func (plugin HttpLogPlugin) Validate() error {
	httpEndpoint, ok := plugin.config["http_endpoint"].(string)
	if !ok || httpEndpoint == "" {
		return fmt.Errorf("http_endpoint must be a non-empty string")
	}

	method, ok := plugin.config["method"].(string)
	if !ok || method == "" {
		return fmt.Errorf("method must be a non-empty string")
	}
	method = strings.ToUpper(method)
	if !util.SliceContainsString([]string{"GET", "POST", "PUT", "PATCH"}, method) {
		return fmt.Errorf("method must be one of: GET, POST, PUT, PATCH")
	}

	contentType, ok := plugin.config["content_type"].(string)
	if !ok || contentType == "" {
		return fmt.Errorf("content_type must be a non-empty string")
	}
	if !strings.HasPrefix(strings.ToLower(contentType), "application/") {
		return fmt.Errorf("content_type must be an application type (e.g., application/json)")
	}

	return nil
}

func (plugin HttpLogPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing httplog function...")
		next.ServeHTTP(w, r)

		// For http logs, any errors should not stop the request, so we will log the error and continue
		config := plugin.parseConfig()

		log, err := plugin.createLogBody(r)
		if err != nil {
			err.WriteLogMessage()
		}

		err = plugin.sendLog(log, config)
		if err != nil {
			err.WriteLogMessage()
		}
	})
}

func (plugin HttpLogPlugin) parseConfig() *HttpLogConfig {
	config := plugin.config

	return &HttpLogConfig{
		httpEndpoint: config["http_endpoint"].(string),
		method:       config["method"].(string),
		contentType:  config["content_type"].(string),
	}
}

func (plugin HttpLogPlugin) createLogBody(r *http.Request) (map[string]interface{}, *model.HttpError) {

	// Get the service and route from the request
	service, route, err := util.GetServiceAndRouteFromRequest(&GlobalProxyConfig, r)
	if err != nil {
		return nil, model.NewHttpError(500, "ERR_PARSING_SERVICE_ROUTE",
			"Error parsing service and route from request")
	}

	lb := NewLoadBalancer(GlobalHealthChecker)

	clientIp, err := util.GetHostIp(r.RemoteAddr)
	if err != nil {
		slog.Error("Error getting client ip", "error", err)
		err.WriteLogMessage()
	}

	upstreamIndexToRoute := lb.GetCurrentUpstream(*service, clientIp)

	// Map the service and route to log
	log := map[string]interface{}{
		"service": map[string]interface{}{
			"name":     service.Name,
			"protocol": service.Protocol,
			"host":     service.Upstreams[upstreamIndexToRoute].Host,
			"port":     service.Upstreams[upstreamIndexToRoute].Port,
		},
		"route": map[string]interface{}{
			"path": route.Path,
		},
		"request": map[string]interface{}{
			"protocol": r.Proto,
			"tls":      r.TLS != nil,
			"method":   r.Method,
			"path":     r.URL.Path,
			"url":      r.URL.String(),
			"uri":      r.RequestURI,
			"size":     util.GetContentLength(r.Header.Get("Content-Length")),
			"headers":  r.Header,
		},
		"client_ip":  clientIp,
		"started_at": time.Now(),
	}
	return log, nil
}

func (plugin HttpLogPlugin) sendLog(log map[string]interface{}, config *HttpLogConfig) *model.HttpError {

	// Convert the payload to JSON
	body, err := json.Marshal(log)
	if err != nil {
		return model.NewHttpError(http.StatusBadGateway, "ERR_PARSING_LOG", "Error sending log when serializing log to JSON")
	}

	// Create a new request with POST method, URL, and payload
	req, err := http.NewRequest(config.method, config.httpEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return model.NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log when creating http request to "+config.method+" "+config.httpEndpoint)
	}

	// Set request headers (optional)
	req.Header.Set("Content-Type", config.contentType)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log to "+config.method+" "+config.httpEndpoint)
	}
	defer resp.Body.Close()

	slog.Info("Successfully sent log to " + config.method + " " + config.httpEndpoint)

	return nil
}
