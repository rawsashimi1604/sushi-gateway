package api

import (
	"encoding/base64"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/errors"
	"log/slog"
	"net/http"
	"strings"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) RegisterRoutes(router *mux.Router) {
	router.PathPrefix("/login").Methods("POST").HandlerFunc(c.Login())
}

func (c *AuthController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			errors.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Authorization header missing").WriteJSONResponse(w)
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			errors.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid Authorization scheme").WriteJSONResponse(w)
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			errors.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid Base64 encoding").WriteJSONResponse(w)
			return
		}

		credentials := string(decodedBytes)
		parts := strings.SplitN(credentials, ":", 2)
		if len(parts) != 2 {
			errors.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid credentials format").WriteJSONResponse(w)
			return
		}

		username, password := parts[0], parts[1]
		if !validate(username, password) {
			errors.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid credentials").WriteJSONResponse(w)
			return
		}

		slog.Info("Login successful for user: " + username)
		return
	}
}

func validate(username string, password string) bool {
	return username == "admin" && password == "changeme"
}
