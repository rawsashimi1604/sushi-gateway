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
	alg       string
	iss       string
	secret    string
	publicKey string
}

func NewJwtPlugin(config map[string]interface{}) *Plugin {
	return &Plugin{
		Name:     constant.PLUGIN_JWT,
		Priority: 1450,
		Phase:    AccessPhase,
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

	// Only HS256 and RSA256 is supported for now
	supportedJwtSigningMethods := []string{constant.HS_256, constant.RSA_256}
	if !util.SliceContainsString(supportedJwtSigningMethods, alg) {
		return fmt.Errorf("alg must be one of: HS256")
	}

	iss, ok := plugin.config["iss"].(string)
	if !ok || iss == "" {
		return fmt.Errorf("iss (issuer) must be a non-empty string")
	}

	if alg == constant.HS_256 {
		secret, ok := plugin.config["secret"].(string)
		if !ok || secret == "" {
			return fmt.Errorf("secret must be a non-empty string")
		}
	}

	if alg == constant.RSA_256 {
		publicKey, ok := plugin.config["publicKey"].(string)
		if !ok || publicKey == "" {
			return fmt.Errorf("publicKey must be a non-empty string")
		}

		// Validate the RSA public key format and structure
		_, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		if err != nil {
			return fmt.Errorf("invalid RSA public key: %v", err)
		}
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

		err = plugin.validateToken(tokenString)
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

func (plugin JwtPlugin) validateToken(token string) *model.HttpError {

	config := plugin.config

	credentials := JwtCredentials{
		alg: config["alg"].(string),
		iss: config["iss"].(string),
	}

	if credentials.alg == constant.HS_256 {
		credentials.secret = config["secret"].(string)
		return plugin.validateHS256(credentials, token)
	} else {
		credentials.publicKey = config["publicKey"].(string)
		return plugin.validateRS256(credentials, token)
	}
}

func (plugin JwtPlugin) validateRS256(credentials JwtCredentials, token string) *model.HttpError {
	tokenInvalidErr := model.NewHttpError(http.StatusUnauthorized, "INVALID_TOKEN", "The token is not valid.")

	// Parse public key
	rsaPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(credentials.publicKey))
	if err != nil {
		slog.Info("Failed to parse RSA public key: " + err.Error())
		return model.NewHttpError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Something went wrong")
	}

	// Parse and validate the token
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})

	if err != nil {
		slog.Info("Error parsing token: " + err.Error())
		return tokenInvalidErr
	}

	if !jwtToken.Valid {
		return tokenInvalidErr
	}

	// Check claims if iss is valid from the token
	if !plugin.isClaimValid(jwtToken, credentials.iss) {
		return tokenInvalidErr
	}

	return nil
}

func (plugin JwtPlugin) validateHS256(credentials JwtCredentials, token string) *model.HttpError {

	tokenInvalidErr := model.NewHttpError(http.StatusUnauthorized, "INVALID_TOKEN", "The token is not valid.")

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if credentials.alg == constant.HS_256 {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
		}
		return []byte(credentials.secret), nil
	})

	if err != nil {
		slog.Info("Error parsing token: " + err.Error())
		return tokenInvalidErr
	}

	if !jwtToken.Valid {
		return tokenInvalidErr
	}

	// Check claims if iss is valid from the token
	if !plugin.isClaimValid(jwtToken, credentials.iss) {
		return tokenInvalidErr
	}

	return nil
}

func (plugin JwtPlugin) isClaimValid(token *jwt.Token, issuer string) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		if iss, ok := claims["iss"].(string); ok {
			if iss == issuer {
				return true
			} else {
				slog.Info(fmt.Sprintf("Invalid JWT issuer: %s", iss))
				return false
			}
		}
	}

	slog.Error("Invalid JWT claims")
	return false
}
