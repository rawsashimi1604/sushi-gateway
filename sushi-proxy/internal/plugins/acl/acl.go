package acl

import (
	"fmt"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
)

// TODO: add allow and deny mechanism, now both are enabled
type AclPlugin struct {
	config map[string]interface{}
}

func NewAclPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
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
			err := errors.NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		if plugin.isBlacklisted(forwardedIP) {
			slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
			err := errors.NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (plugin AclPlugin) isWhitelisted(ip string) bool {
	// TODO: add validation for this plugin in the config file
	data := plugin.config["data"].(map[string]interface{})
	whitelistInterface := data["whitelist"].([]interface{}) // Assert to []interface{} first

	var whitelist []string
	for _, w := range whitelistInterface {
		ipStr := w.(string)
		whitelist = append(whitelist, ipStr)
	}

	for _, whitelistedIP := range whitelist {
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}

func (plugin AclPlugin) isBlacklisted(ip string) bool {
	// TODO: add validation for this plugin in the config file
	data := plugin.config["data"].(map[string]interface{})
	blacklistInterface := data["blacklist"].([]interface{}) // Assert to []interface{} first

	var blacklist []string
	for _, w := range blacklistInterface {
		ipStr := w.(string)
		blacklist = append(blacklist, ipStr)
	}

	for _, blacklistedIP := range blacklist {
		if ip == blacklistedIP {
			return true
		}
	}
	return false
}
