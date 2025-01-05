package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

const (
	testHS256Secret = "test-secret-key"
	testIssuer      = "test-issuer"
	// Test RSA keys (these are test keys, never use in production)
	testRSAPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCQCDqfAqoSsLupIJSTLBO5vfCB
GehwP7R+uw926TZzkOAbi3Cp/YILG/OhCiSuYpqjC4yP9EL/Gt2bnBej/wLzzXT8
ofn5wJtxc+60JSjK6bOpiStFcCe8jBpaQNN9zumUlTU7dBWHF2BEHYBX7wjjjwtJ
+5Aehp4df5XyhHvJbwIDAQAB
-----END PUBLIC KEY-----`
	testRSAPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCQCDqfAqoSsLupIJSTLBO5vfCBGehwP7R+uw926TZzkOAbi3Cp
/YILG/OhCiSuYpqjC4yP9EL/Gt2bnBej/wLzzXT8ofn5wJtxc+60JSjK6bOpiStF
cCe8jBpaQNN9zumUlTU7dBWHF2BEHYBX7wjjjwtJ+5Aehp4df5XyhHvJbwIDAQAB
AoGAD3wzAUGCE3xY0LgmASSuAjw/jUHB0b+nojBuWzR7nDUpQwzc0gnlX1hj+x5i
DzWetoTZWejDAzZaOQ6xI/qY0H/xrQCjhiDptsxSjd8xXwC2M03gQUlbzj9uwgO8
QH6jqlDvKOTb0P6b7YHkY0CXoRtII0sm9fNV/R0G0t+U1sECQQDVsyPnqofPlFOc
rNF/hjo8nHdI7Z0axp/bnQlG6j7VERCfxtjTc+W/FdRUQbiLfnKP0PyzxudP8jbG
4+ohaCffAkEArIrN9j14B49NSjzKrTg/prmaThAzEND1VRzWOf4FGhp1/DVGy8TO
PJJfAcg1jW/Ivl1zdJx8hohsm1dHAoPQcQJAdRVwb6Z0QJww71+UbP1a/QhxJqjh
ceEvhsDUa2E+SbjO1eu5sqkGUJqiOgPEG9GM7RUAz3MEGz5HGtOW3PTXGwJBAKdQ
Rl70xnMWPA20G5mThO2o53+xZ8NzzaMGPpqnv8zLQgQaqZcpNhA4o9Z3ja6kalZn
CnFW2c4fdqnAHYTLy5ECQQDQGfm/n3Iv7m556GgLlMVaDtb9cLjIaqbmbSpDs2I8
LnUKSTg8AGSYqP2tABaqQbrO0RzcFcTsaPvgOXbq1Ttz
-----END RSA PRIVATE KEY-----`
)

func createTestToken(signingMethod jwt.SigningMethod, key interface{}, issuer string) (string, error) {
	claims := jwt.MapClaims{
		"iss": issuer,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(signingMethod, claims)
	return token.SignedString(key)
}

func TestJwtPlugin_Validate(t *testing.T) {
	tests := []struct {
		name        string
		config      map[string]interface{}
		expectError bool
	}{
		{
			name: "Valid HS256 config",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			expectError: false,
		},
		{
			name: "Valid RS256 config",
			config: map[string]interface{}{
				"alg":       "RS256",
				"iss":       testIssuer,
				"publicKey": testRSAPublicKey,
			},
			expectError: false,
		},
		{
			name: "Missing algorithm",
			config: map[string]interface{}{
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			expectError: true,
		},
		{
			name: "Invalid algorithm",
			config: map[string]interface{}{
				"alg":    "INVALID",
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			expectError: true,
		},
		{
			name: "Missing secret for HS256",
			config: map[string]interface{}{
				"alg": "HS256",
				"iss": testIssuer,
			},
			expectError: true,
		},
		{
			name: "Missing public key for RS256",
			config: map[string]interface{}{
				"alg": "RS256",
				"iss": testIssuer,
			},
			expectError: true,
		},
		{
			name: "Invalid RSA public key",
			config: map[string]interface{}{
				"alg":       "RS256",
				"iss":       testIssuer,
				"publicKey": "invalid-key",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin := JwtPlugin{config: tt.config}
			err := plugin.Validate()
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestJwtPlugin_Execute(t *testing.T) {
	tests := []struct {
		name           string
		config         map[string]interface{}
		setupToken     func() string
		expectStatus   int
		expectResponse string
	}{
		{
			name: "Valid HS256 token",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			setupToken: func() string {
				token, _ := createTestToken(jwt.SigningMethodHS256, []byte(testHS256Secret), testIssuer)
				return token
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "Valid RS256 token",
			config: map[string]interface{}{
				"alg":       "RS256",
				"iss":       testIssuer,
				"publicKey": testRSAPublicKey,
			},
			setupToken: func() string {
				privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(testRSAPrivateKey))
				token, _ := createTestToken(jwt.SigningMethodRS256, privateKey, testIssuer)
				return token
			},
			expectStatus: http.StatusOK,
		},
		{
			name: "Invalid token signature",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			setupToken: func() string {
				token, _ := createTestToken(jwt.SigningMethodHS256, []byte("wrong-secret"), testIssuer)
				return token
			},
			expectStatus: http.StatusUnauthorized,
		},
		{
			name: "Missing token",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			setupToken: func() string {
				return ""
			},
			expectStatus: http.StatusUnauthorized,
		},
		{
			name: "Wrong issuer",
			config: map[string]interface{}{
				"alg":    "HS256",
				"iss":    testIssuer,
				"secret": testHS256Secret,
			},
			setupToken: func() string {
				token, _ := createTestToken(jwt.SigningMethodHS256, []byte(testHS256Secret), "wrong-issuer")
				return token
			},
			expectStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin := NewJwtPlugin(tt.config)
			handler := plugin.Handler.(JwtPlugin)

			// Create test request
			req := httptest.NewRequest("GET", "/test", nil)
			if token := tt.setupToken(); token != "" {
				req.Header.Set("Authorization", "Bearer "+token)
			}

			// Create response recorder
			rr := httptest.NewRecorder()

			// Create next handler that always returns 200 OK
			nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// Execute plugin
			handler.Execute(nextHandler).ServeHTTP(rr, req)

			assert.Equal(t, tt.expectStatus, rr.Code)
		})
	}
}

func TestVerifyAndParseAuthHeaderJwt(t *testing.T) {
	tests := []struct {
		name           string
		authHeader     string
		expectError    bool
		expectedToken  string
		expectedStatus int
	}{
		{
			name:           "Valid auth header",
			authHeader:     "Bearer token123",
			expectError:    false,
			expectedToken:  "token123",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Missing auth header",
			authHeader:     "",
			expectError:    true,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid format - no Bearer",
			authHeader:     "token123",
			expectError:    true,
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "Invalid format - extra parts",
			authHeader:     "Bearer token123 extra",
			expectError:    true,
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			token, err := verifyAndParseAuthHeaderJwt(req)
			if tt.expectError {
				assert.NotNil(t, err)
				assert.Equal(t, tt.expectedStatus, err.HttpCode)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expectedToken, token)
			}
		})
	}
}
