package bot_protection

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"strings"
)

type BotProtectionPlugin struct {
	config map[string]interface{}
}

func NewBotProtectionPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_BOT_PROTECTION,
		Priority: 10,
		Handler: BotProtectionPlugin{
			config: config,
		},
	}
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

func (plugin BotProtectionPlugin) verifyIsBot(userAgent string) *errors.HttpError {
	// TODO: add validation for this plugin in the config file
	data := plugin.config["data"].(map[string]interface{})
	blacklist := util.ToStringSlice(data["blacklist"].([]interface{}))

	for _, bot := range blacklist {
		if strings.Contains(userAgent, bot) {
			slog.Info("Bot detected: " + userAgent)
			return errors.NewHttpError(http.StatusForbidden, "BOT_DETECTED", "Bot detected.")
		}
	}
	return nil
}
