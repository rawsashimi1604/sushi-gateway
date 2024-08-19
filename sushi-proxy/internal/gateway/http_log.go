package gateway

import (
	"bytes"
	"encoding/json"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
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
		Priority: 10000,
		Handler: HttpLogPlugin{
			config: config,
		},
	}
}

func (plugin HttpLogPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing httplog function...")
		next.ServeHTTP(w, r)

		config := plugin.parseConfig()

		log, err := plugin.createLogBody(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		err = plugin.sendLog(log, config)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}
	})
}

func (plugin HttpLogPlugin) parseConfig() *HttpLogConfig {
	data := plugin.config["data"].(map[string]interface{})

	return &HttpLogConfig{
		httpEndpoint: data["http_endpoint"].(string),
		method:       data["method"].(string),
		contentType:  data["content_type"].(string),
	}
}

func (plugin HttpLogPlugin) createLogBody(r *http.Request) (map[string]interface{}, *HttpError) {

	// Get the service and route from the request
	service, route, err := GetServiceAndRouteFromRequest(&GlobalProxyConfig, r)
	if err != nil {
		return nil, NewHttpError(500, "ERR_PARSING_SERVICE_ROUTE",
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
			"size":     GetContentLength(r.Header.Get("Content-Length")),
			"headers":  r.Header,
		},
		"client_ip":  r.RemoteAddr,
		"started_at": time.Now(),
	}
	return log, nil
}

func (plugin HttpLogPlugin) sendLog(log map[string]interface{}, config *HttpLogConfig) *HttpError {

	// Convert the payload to JSON
	body, err := json.Marshal(log)
	if err != nil {
		return NewHttpError(http.StatusBadGateway, "ERR_PARSING_LOG", "Error sending log when serializing log to JSON")
	}

	// Create a new request with POST method, URL, and payload
	req, err := http.NewRequest(config.method, config.httpEndpoint, bytes.NewBuffer(body))
	if err != nil {
		return NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log when creating http request")
	}

	// Set request headers (optional)
	req.Header.Set("Content-Type", config.contentType)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log")
	}
	defer resp.Body.Close()
	return nil
}
