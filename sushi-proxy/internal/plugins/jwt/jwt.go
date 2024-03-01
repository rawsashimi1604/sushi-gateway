package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/plugins"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"strings"
)

type JwtPlugin struct{}

type JwtCredentials struct {
	alg    string `json:"alg"`
	iss    string `json:"iss"`
	secret string `json:"secret"`
}

var Plugin = NewJwtPlugin()

func (plugin JwtPlugin) Execute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Executing jwt auth function...")

		creds, err := loadCredentialsFromReq(r)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		tokenString, err := verifyAndParseAuthHeader(r)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		_, err = validateToken(creds, tokenString)
		if err != nil {
			writeWWWAuthenticateHeader(w)
			err.WriteJSONResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewJwtPlugin() *plugins.Plugin {
	return &plugins.Plugin{
		Name:     "jwt",
		Priority: 15,
		Handler:  JwtPlugin{},
	}
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

func loadCredentialsFromReq(req *http.Request) (*JwtCredentials, *errors.HttpError) {
	service, _, err := util.GetServiceAndRouteFromRequest(req)
	if err != nil {
		return nil, err
	}

	for _, cred := range service.Credentials {
		if cred.Plugin == "jwt" {
			// Get from cred map
			alg, ok := cred.Data["alg"].(string)
			iss, ok := cred.Data["iss"].(string)
			secret, ok := cred.Data["secret"].(string)
			if !ok {
				return nil, errors.NewHttpError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "invalid credentials, please try again.")
			}
			return &JwtCredentials{
				alg:    alg,
				iss:    iss,
				secret: secret,
			}, nil
		}
	}

	return nil, errors.NewHttpError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "invalid credentials, please try again.")
}

func validateToken(credentials *JwtCredentials, token string) (*jwt.Token, *errors.HttpError) {
	tokenInvalidErr := errors.NewHttpError(http.StatusUnauthorized, "INVALID_TOKEN", "The token is not valid.")

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
		slog.Info("Error parsing token: ", err.Error())
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
