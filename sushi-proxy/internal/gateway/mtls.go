package gateway

import (
	"crypto/x509"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"log/slog"
	"net/http"
)

type MtlsPlugin struct {
	config map[string]interface{}
}

func NewMtlsPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_MTLS,
		Priority: 12,
		Handler:  MtlsPlugin{config: config},
	}
}

func (plugin MtlsPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing mtls function...")

		if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
			err := model.NewHttpError(http.StatusUnauthorized,
				"MISSING_CERTIFICATE", "Missing client certificate in request.")
			err.WriteJSONResponse(w)
			return
		}

		// Verify the client certificate
		opts := x509.VerifyOptions{
			Roots: GlobalCaCertPool.Pool,
		}

		if _, err := r.TLS.PeerCertificates[0].Verify(opts); err != nil {
			err := model.NewHttpError(http.StatusUnauthorized,
				"INVALID_CLIENT_CERTIFICATE", "Invalid client cert")
			err.WriteJSONResponse(w)
			return
		}

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
