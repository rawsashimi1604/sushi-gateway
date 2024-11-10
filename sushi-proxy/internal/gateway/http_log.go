package gateway

import (
	"bytes"
	"encoding/json"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"time"
)

type HttpLogPlugin struct {
	config map[string]interface{}
}

// TODO: parse log gateway
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
		Handler: HttpLogPlugin{
			config: config,
		},
	}
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

	lb := NewLoadBalancer()
	upstreamIndexToRoute := lb.GetCurrentUpstream(*service)

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
		"client_ip":  r.RemoteAddr,
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
		return model.NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log when creating http request")
	}

	// Set request headers (optional)
	req.Header.Set("Content-Type", config.contentType)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log")
	}
	defer resp.Body.Close()
	return nil
}
