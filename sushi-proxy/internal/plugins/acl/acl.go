package acl

import (
	"fmt"
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

		// Note: In a real-world scenario,
		// you might need to parse X-Forwarded-For header instead
		clientIP := r.RemoteAddr

		// Check if the IP is in the whitelist
		if isWhitelisted(clientIP) {
			next.ServeHTTP(w, r)
			return
		}

		// Check if the IP is in the blacklist
		if isBlacklisted(clientIP) {
			slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
			http.Error(w, "Access Denied", http.StatusForbidden)
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
