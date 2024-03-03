package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"log/slog"
	"net/http"
	"strings"
)

type JwtPlugin struct {
	config map[string]interface{}
}

type JwtCredentials struct {
	alg    string `json:"alg"`
	iss    string `json:"iss"`
	secret string `json:"secret"`
}

func NewJwtPlugin(config map[string]interface{}) *plugins.Plugin {
	return &plugins.Plugin{
		Name:     constant.PLUGIN_JWT,
		Priority: 200,
		Handler: JwtPlugin{
			config: config,
		},
	}
}

func (plugin JwtPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing jwt auth function...")

		tokenString, err := verifyAndParseAuthHeader(r)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		_, err = plugin.validateToken(tokenString)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func writeWWWAuthenticateHeader(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate",
		fmt.Sprintf("Bearer realm=\"%s\", "+
			"charset=%s",
			"Access to sushi gateway", constant.UTF_8))
}

func verifyAndParseAuthHeader(req *http.Request) (string, *errors.HttpError) {
	authHeader := req.Header.Get("Authorization")
	bits := strings.Split(authHeader, " ")

	// valid format : Bearer token
	isValidAuthFormat := authHeader != "" && len(bits) == 2
	if !isValidAuthFormat {
		slog.Info("Invalid jwt auth format passed in.")
		return "", errors.NewHttpError(http.StatusUnauthorized,
			"MALFORMED_AUTH_HEADER", "Invalid auth format passed in.")
	}

	return bits[1], nil
}

func (plugin JwtPlugin) validateToken(token string) (*jwt.Token, *errors.HttpError) {
	tokenInvalidErr := errors.NewHttpError(http.StatusUnauthorized, "INVALID_TOKEN", "The token is not valid.")

	data := plugin.config["data"].(map[string]interface{})
	credentials := JwtCredentials{
		alg:    data["alg"].(string),
		iss:    data["iss"].(string),
		secret: data["secret"].(string),
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
