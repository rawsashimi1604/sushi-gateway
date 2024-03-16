package http_log

import (
	"bytes"
	"encoding/json"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/config"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"time"
)

type HttpLogPlugin struct {
	config map[string]interface{}
}

// TODO: parse log config
// TODO: send logs in batches (batch processing).
// TODO: log the response as well.
type HttpLogConfig struct {
	httpEndpoint  string
	method        string
	contentType   string
	timeout       uint
	retryCount    int
	queueSize     int
	flushInterval int
	headers       map[string]interface{}
}

func NewHttpLogPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
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

		log, err := plugin.createLogBody(r)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		err = plugin.sendLog(log)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}
	})
}

func (plugin HttpLogPlugin) parseConfig(config map[string]interface{}) (*HttpLogConfig, *errors.HttpError) {
	// TODO: read log config
	return nil, nil
}

func (plugin HttpLogPlugin) createLogBody(r *http.Request) (map[string]interface{}, *errors.HttpError) {

	// Get the service and route from the request
	service, route, err := util.GetServiceAndRouteFromRequest(&config.GlobalProxyConfig, r)
	if err != nil {
		return nil, errors.NewHttpError(500, "ERR_PARSING_SERVICE_ROUTE",
			"Error parsing service and route from request")
	}

	// Map the service and route to log
	log := map[string]interface{}{
		// TODO: for now use 1st upstream, later use round robin alg
		"service": map[string]interface{}{
			"name":     service.Name,
			"protocol": service.Protocol,
			"host":     service.Upstreams[0].Host,
			"port":     service.Upstreams[0].Port,
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
			"size":     util.ParseContentLength(r.Header.Get("Content-Length")),
			"headers":  r.Header,
		},
		"client_ip":  r.RemoteAddr,
		"started_at": time.Now(),
	}
	return log, nil
}

func (plugin HttpLogPlugin) sendLog(log map[string]interface{}) *errors.HttpError {

	// Convert the payload to JSON
	body, err := json.Marshal(log)
	if err != nil {
		return errors.NewHttpError(http.StatusBadGateway, "ERR_PARSING_LOG", "Error sending log when serializing log to JSON")
	}

	// Create a new request with POST method, URL, and payload
	req, err := http.NewRequest("POST", "http://localhost:8003/v1/log", bytes.NewBuffer(body))
	if err != nil {
		return errors.NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log when creating http request")
	}

	// Set request headers (optional)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.NewHttpError(http.StatusBadGateway, "ERR_SENDING_LOG", "Error sending log")
	}
	defer resp.Body.Close()
	return nil
}
