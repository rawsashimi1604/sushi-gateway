package gateway

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
)

type AclPlugin struct {
	config map[string]interface{}
}

func NewAclPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_ACL,
		Priority: 10,
		Handler: AclPlugin{
			config: config,
		},
	}
}

func (plugin AclPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing acl function...")

		// Check both forwarded ip and client ip
		clientIP := r.RemoteAddr
		forwardedIP := r.Header.Get(constant.X_FORWARDED_FOR)

		// Check if the IP is in the whitelist
		if plugin.isWhitelisted(clientIP) || plugin.isWhitelisted(forwardedIP) {
			next.ServeHTTP(w, r)
			return
		}

		if plugin.isBlacklisted(clientIP) {
			slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
			err := NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		if plugin.isBlacklisted(forwardedIP) {
			slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
			err := NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ValidateACLPlugin(config map[string]interface{}) {
	// TODO: complete
}

func (plugin AclPlugin) isWhitelisted(ip string) bool {
	// TODO: add validation for this plugin in the gateway file
	config := plugin.config
	whitelist := util.ToStringSlice(config["whitelist"].([]interface{}))

	for _, whitelistedIP := range whitelist {
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}

func (plugin AclPlugin) isBlacklisted(ip string) bool {
	// TODO: add validation for this plugin in the gateway file
	config := plugin.config
	blacklist := util.ToStringSlice(config["blacklist"].([]interface{}))

	for _, blacklistedIP := range blacklist {
		if ip == blacklistedIP {
			return true
		}
	}
	return false
}
