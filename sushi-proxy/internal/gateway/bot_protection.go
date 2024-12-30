package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

type BotProtectionPlugin struct {
	config map[string]interface{}
}

func NewBotProtectionPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_BOT_PROTECTION,
		Priority: 2500,
		Handler: BotProtectionPlugin{
			config: config,
		},
		Validator: BotProtectionPlugin{
			config: config,
		},
	}
}

func (plugin BotProtectionPlugin) Validate() error {
	blacklist, exists := plugin.config["blacklist"]
	if !exists {
		return fmt.Errorf("blacklist configuration is required")
	}

	blacklistArr, ok := blacklist.([]interface{})
	if !ok {
		return fmt.Errorf("blacklist must be an array of strings")
	}

	if len(blacklistArr) == 0 {
		return fmt.Errorf("blacklist cannot be empty")
	}

	for _, bot := range blacklistArr {
		if _, ok := bot.(string); !ok {
			return fmt.Errorf("blacklist entries must be strings")
		}
		if bot.(string) == "" {
			return fmt.Errorf("blacklist entries cannot be empty strings")
		}
	}

	return nil
}

func (plugin BotProtectionPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing bot protection function...")

		userAgent := r.Header.Get("User-Agent")
		if userAgent == "" {
			next.ServeHTTP(w, r)
		}

		err := plugin.verifyIsBot(userAgent)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (plugin BotProtectionPlugin) verifyIsBot(userAgent string) *model.HttpError {
	// TODO: add validation for this plugin in the gateway file
	config := plugin.config
	blacklist := util.ToStringSlice(config["blacklist"].([]interface{}))

	for _, bot := range blacklist {
		if strings.Contains(userAgent, bot) {
			slog.Info("Bot detected: " + userAgent)
			return model.NewHttpError(http.StatusForbidden, "BOT_DETECTED", "Bot detected.")
		}
	}
	return nil
}
