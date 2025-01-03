package gateway

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
)

type JwtPlugin struct {
	config map[string]interface{}
}

type JwtCredentials struct {
	alg    string `json:"alg"`
	iss    string `json:"iss"`
	secret string `json:"secret"`
}

func NewJwtPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_JWT,
		Priority: 1450,
		Handler: JwtPlugin{
			config: config,
		},
		Validator: JwtPlugin{
			config: config,
		},
	}
}

func (plugin JwtPlugin) Validate() error {
	alg, ok := plugin.config["alg"].(string)
	if !ok || alg == "" {
		return fmt.Errorf("alg must be a non-empty string")
	}

	// Only HS_256 is supported for now
	if !util.SliceContainsString([]string{constant.HS_256}, alg) {
		return fmt.Errorf("alg must be one of: HS256")
	}

	iss, ok := plugin.config["iss"].(string)
	if !ok || iss == "" {
		return fmt.Errorf("iss (issuer) must be a non-empty string")
	}

	secret, ok := plugin.config["secret"].(string)
	if !ok || secret == "" {
		return fmt.Errorf("secret must be a non-empty string")
	}

	return nil
}

func (plugin JwtPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing jwt auth function...")

		tokenString, err := verifyAndParseAuthHeaderJwt(r)
		if err != nil {
			writeWWWAuthenticateHeaderJwt(w)
			err.WriteJSONResponse(w)
			return
		}

		_, err = plugin.validateToken(tokenString)
		if err != nil {
			writeWWWAuthenticateHeaderJwt(w)
			err.WriteJSONResponse(w)
			return
		}

		// Strip Authorization header
		r.Header.Del("Authorization")

		next.ServeHTTP(w, r)
	})

}

func writeWWWAuthenticateHeaderJwt(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate",
		fmt.Sprintf("Bearer realm=\"%s\", "+
			"charset=%s",
			"Access to sushi gateway", constant.UTF_8))
}

func verifyAndParseAuthHeaderJwt(req *http.Request) (string, *model.HttpError) {
	authHeader := req.Header.Get("Authorization")
	bits := strings.Split(authHeader, " ")

	// valid format : Bearer token
	isValidAuthFormat := authHeader != "" && len(bits) == 2
	if !isValidAuthFormat {
		slog.Info("Invalid jwt auth format passed in.")
		return "", model.NewHttpError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid auth format passed in.")
	}

	return bits[1], nil
}

func (plugin JwtPlugin) validateToken(token string) (*jwt.Token, *model.HttpError) {
	tokenInvalidErr := model.NewHttpError(http.StatusUnauthorized, "INVALID_TOKEN", "The token is not valid.")

	config := plugin.config
	credentials := JwtCredentials{
		alg:    config["alg"].(string),
		iss:    config["iss"].(string),
		secret: config["secret"].(string),
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// TODO: do for other alg types (RSA 256)
		if credentials.alg == constant.HS_256 {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
		}
		return []byte(credentials.secret), nil
	})

	if err != nil {
		slog.Info("Error parsing token: " + err.Error())
		return nil, tokenInvalidErr
	}

	if !jwtToken.Valid {
		return nil, tokenInvalidErr
	}

	// Check claims if iss is valid from the token
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if ok {
		if iss, ok := claims["iss"].(string); ok {
			if iss == credentials.iss {
				return jwtToken, nil
			} else {
				slog.Info(fmt.Sprintf("Invalid JWT issuer: %s", iss))
				return nil, tokenInvalidErr
			}
		}
	}

	return nil, tokenInvalidErr
}
