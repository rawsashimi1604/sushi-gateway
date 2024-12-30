package gateway

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

type AclPlugin struct {
	config map[string]interface{}
}

func NewAclPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_ACL,
		Priority: 950,
		Handler: AclPlugin{
			config: config,
		},
		Validator: AclPlugin{
			config: config,
		},
	}
}

func (plugin AclPlugin) Validate() error {
	whitelist, hasWhitelist := plugin.config["whitelist"]
	blacklist, hasBlacklist := plugin.config["blacklist"]

	// Check that only one list type exists
	if hasWhitelist && hasBlacklist {
		return fmt.Errorf("only one of whitelist or blacklist can be configured at a time")
	}

	if !hasWhitelist && !hasBlacklist {
		return fmt.Errorf("either whitelist or blacklist must be configured")
	}

	// Validate the present list
	if hasWhitelist {
		whitelistArr, ok := whitelist.([]interface{})
		if !ok {
			return fmt.Errorf("whitelist must be an array of strings")
		}
		for _, ip := range whitelistArr {
			if _, ok := ip.(string); !ok {
				return fmt.Errorf("whitelist entries must be strings")
			}
		}
	}

	if hasBlacklist {
		blacklistArr, ok := blacklist.([]interface{})
		if !ok {
			return fmt.Errorf("blacklist must be an array of strings")
		}
		for _, ip := range blacklistArr {
			if _, ok := ip.(string); !ok {
				return fmt.Errorf("blacklist entries must be strings")
			}
		}
	}

	return nil
}

func (plugin AclPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing acl function...")

		// Check both forwarded ip and client ip
		clientIP := r.RemoteAddr
		forwardedIP := r.Header.Get(constant.X_FORWARDED_FOR)

		// If whitelist exists, only allow whitelisted IPs
		if _, exists := plugin.config["whitelist"]; exists {
			if plugin.isWhitelisted(clientIP) || plugin.isWhitelisted(forwardedIP) {
				next.ServeHTTP(w, r)
				return
			}
			err := model.NewHttpError(http.StatusForbidden,
				"ACCESS_DENIED", "Access Denied")
			err.WriteJSONResponse(w)
			return
		}

		// If blacklist exists, deny blacklisted IPs
		if _, exists := plugin.config["blacklist"]; exists {
			if plugin.isBlacklisted(clientIP) || plugin.isBlacklisted(forwardedIP) {
				slog.Info(fmt.Sprintf("IP %s is blacklisted", clientIP))
				err := model.NewHttpError(http.StatusForbidden,
					"ACCESS_DENIED", "Access Denied")
				err.WriteJSONResponse(w)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		// If neither list exists (shouldn't happen due to validation), allow the request
		next.ServeHTTP(w, r)
	})
}

func ValidateACLPlugin(config map[string]interface{}) {
	// TODO: complete
}

func (plugin AclPlugin) isWhitelisted(ip string) bool {
	if ip == "" {
		return false
	}

	config := plugin.config
	whitelist, exists := config["whitelist"]
	if !exists {
		return false
	}

	whitelistArr, ok := whitelist.([]interface{})
	if !ok {
		return false
	}

	for _, whitelistedIP := range util.ToStringSlice(whitelistArr) {
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}

func (plugin AclPlugin) isBlacklisted(ip string) bool {
	if ip == "" {
		return false
	}

	config := plugin.config
	blacklist, exists := config["blacklist"]
	if !exists {
		return false
	}

	blacklistArr, ok := blacklist.([]interface{})
	if !ok {
		return false
	}

	for _, blacklistedIP := range util.ToStringSlice(blacklistArr) {
		if ip == blacklistedIP {
			return true
		}
	}
	return false
}
