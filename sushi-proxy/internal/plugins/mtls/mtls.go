package mtls

import (
	"crypto/x509"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/cert"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type MtlsPlugin struct{}

func NewMtlsPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_MTLS,
		Priority: 12,
		Handler:  MtlsPlugin{},
	}
}

func (plugin MtlsPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing mtls function...")

		if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
			err := errors.NewHttpError(http.StatusUnauthorized,
				"MISSING_CERTIFICATE", "Missing client certificate in request.")
			err.WriteJSONResponse(w)
			return
		}

		// Verify the client certificate
		opts := x509.VerifyOptions{
			Roots: cert.GlobalCaCertPool.Pool,
			// TODO: set more options for verification
		}

		if _, err := r.TLS.PeerCertificates[0].Verify(opts); err != nil {
			err := errors.NewHttpError(http.StatusUnauthorized,
				"INVALID_CLIENT_CERTIFICATE", "Invalid client cert")
			err.WriteJSONResponse(w)
			return
		}

		// call the next plugin.
		next.ServeHTTP(w, r)
	})
}
