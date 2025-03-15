package api

import (
	"context"
	"encoding/base64"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
)

// TODO: externalise this
var jwtKey = []byte("secret-jwt-key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (c *AuthController) RegisterRoutes(router *mux.Router) {
	router.Path("/login").Methods("POST").HandlerFunc(c.Login())
	router.Path("/logout").Methods("DELETE").HandlerFunc(c.Logout())
}

func (c *AuthController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("AuthController:: Admin API - Logging in")
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Authorization header missing").WriteJSONResponse(w)
			return
		}

		if !strings.HasPrefix(authHeader, "Basic ") {
			model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid Authorization scheme").WriteJSONResponse(w)
			return
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid Base64 encoding").WriteJSONResponse(w)
			return
		}

		credentials := string(decodedBytes)
		parts := strings.SplitN(credentials, ":", 2)
		if len(parts) != 2 {
			model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid credentials format").WriteJSONResponse(w)
			return
		}

		username, password := parts[0], parts[1]
		if !validate(username, password) {
			model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH",
				"Invalid credentials").WriteJSONResponse(w)
			return
		}

		tokenString, err := generateJWT("user")
		if err != nil {
			model.NewHttpError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR",
				"Error generating JWT token").WriteJSONResponse(w)
			return
		}

		// Set JWT in cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		slog.Info("Login success:: " + username)
		w.WriteHeader(http.StatusOK)
	}
}

func (c *AuthController) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		slog.Info("AuthController:: Admin API - Logging out")
		// To log out, we invalidate the token by setting a past expiration date
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})
		w.WriteHeader(http.StatusOK)
	}
}

func validate(username string, password string) bool {
	return username == gateway.GlobalAppConfig.AdminUser &&
		password == gateway.GlobalAppConfig.AdminPassword
}

func generateJWT(username string) (string, error) {
	// Set expiration time to 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "sushi-gateway-admin-api",
			Audience:  []string{"sushi-gateway-manager"},
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateJWT(tokenString string) (*Claims, *model.HttpError) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH", "Invalid token")
	}
	if !token.Valid {
		return nil, model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH", "Invalid token")
	}

	return claims, nil
}

func ProtectRouteUsingJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("token")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				model.NewHttpError(http.StatusUnauthorized, "UNAUTHORIZED_AUTH", "Invalid token").WriteJSONResponse(w)
				return
			}
			model.NewHttpError(http.StatusBadRequest, "BAD_REQUEST", "Bad Request")
			return
		}

		claims, httperr := validateJWT(cookie.Value)
		if httperr != nil {
			httperr.WriteJSONResponse(w)
			return
		}

		// Store the claims in the request context for use in handlers
		ctx := req.Context()
		ctx = context.WithValue(ctx, "username", claims.Username)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
