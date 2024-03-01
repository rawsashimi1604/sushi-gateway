package bot_protection

import (
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
	"strings"
)

type BotProtectionPlugin struct{}

var Plugin = NewBotProtectionPlugin()

// TODO: externalize this list to config file
var blacklist = []string{"googlebot", "yahoobot", "bingbot"}

func NewBotProtectionPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "bot_protection",
		Priority: 10,
		Handler:  BotProtectionPlugin{},
	}
}

func (plugin BotProtectionPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing bot protection function...")

		userAgent := r.Header.Get("User-Agent")
		err := verifyIsBot(userAgent)
		if err != nil {
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func verifyIsBot(userAgent string) *errors.HttpError {
	for _, bot := range blacklist {
		if strings.Contains(userAgent, bot) {
			slog.Info("Bot detected: ", userAgent)
			return errors.NewHttpError(http.StatusForbidden, "BOT_DETECTED", "Bot detected.")
		}
	}
	return nil
}
