package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

// Auth validates JWT Bearer tokens or API Keys.
// On success it sets "tenant_id" and "user_id" in the request context.
// This is fail-closed by default. Local development may set ANI_AUTH_MODE=dev
// and pass X-Dev-Tenant-ID to exercise routes before auth-service exists.
func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if isPublicPath(string(c.Path())) {
			c.Next(ctx)
			return
		}

		if os.Getenv("ANI_AUTH_MODE") == "dev" {
			tenantID := string(c.GetHeader("X-Dev-Tenant-ID"))
			if tenantID == "" {
				tenantID = "00000000-0000-0000-0000-000000000001"
			}
			userID := string(c.GetHeader("X-Dev-User-ID"))
			if userID == "" {
				userID = "00000000-0000-0000-0000-000000000001"
			}
			c.Set("tenant_id", tenantID)
			c.Set("user_id", userID)
			c.Set("roles", []string{"tenant-admin"})
			c.Next(ctx)
			return
		}

		// 1. Try Bearer token
		authHeader := string(c.GetHeader("Authorization"))
		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			tenantID, userID, roles, err := verifyJWT(token)
			if err != nil {
				respondError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid or expired token")
				return
			}
			c.Set("tenant_id", tenantID)
			c.Set("user_id", userID)
			c.Set("roles", roles)
			c.Next(ctx)
			return
		}

		// 2. Try API Key
		apiKey := string(c.GetHeader("X-API-Key"))
		if apiKey != "" {
			tenantID, roles, err := verifyAPIKey(apiKey)
			if err != nil {
				respondError(c, http.StatusUnauthorized, "UNAUTHORIZED", "invalid api key")
				return
			}
			c.Set("tenant_id", tenantID)
			c.Set("user_id", "")
			c.Set("roles", roles)
			c.Next(ctx)
			return
		}

		respondError(c, http.StatusUnauthorized, "UNAUTHORIZED", "authentication required")
	}
}

func isPublicPath(path string) bool {
	switch path {
	case "/health", "/ready", "/api/v1/branding", "/api/v1/auth/token", "/api/v1/auth/refresh":
		return true
	default:
		return false
	}
}

// verifyJWT is a stub; replace with real RS256 verification against Dex JWKS endpoint.
func verifyJWT(token string) (tenantID, userID string, roles []string, err error) {
	// TODO: fetch JWKS from Dex, verify signature, extract claims
	_ = token
	return "", "", nil, errAuthUnavailable
}

// verifyAPIKey looks up the hashed API key in the database.
func verifyAPIKey(key string) (tenantID string, roles []string, err error) {
	// TODO: SHA256(key) → query api_keys table → return tenant_id
	_ = key
	return "", nil, errAuthUnavailable
}

var errAuthUnavailable = errors.New("auth verifier is not configured")
