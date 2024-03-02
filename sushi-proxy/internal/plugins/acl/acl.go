package acl

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

type AclPlugin struct{}

var Plugin = NewAclPlugin()

// TODO: externalize this to config file
var whitelist = []string{"127.0.0.1"}
var blacklist = []string{"192.168.1.1"}

func NewAclPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "acl",
		Priority: 10,
		Handler:  AclPlugin{},
	}
}

func (plugin AclPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing acl function...")

		// Check both forwarded ip and client ip
		clientIP := r.RemoteAddr
		forwardedIP := r.Header.Get(constant.X_FORWARDED_FOR)
		slog.Info("clientIP: " + clientIP)
		slog.Info("forwardedIP: " + forwardedIP)

		// Check if the IP is in the whitelist
		if isWhitelisted(clientIP) || isWhitelisted(forwardedIP) {
			next.ServeHTTP(w, r)
			return
		}

		if isBlacklisted(clientIP) {
			slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
			err := errors.NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		if isBlacklisted(forwardedIP) {
			slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
			err := errors.NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isWhitelisted(ip string) bool {
	for _, whitelistedIP := range whitelist {
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}

func isBlacklisted(ip string) bool {
	for _, blacklistedIP := range blacklist {
		if ip == blacklistedIP {
			return true
		}
	}
	return false
}
