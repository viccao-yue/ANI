package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// RateLimit enforces per-tenant token-bucket rate limiting.
// Limits are stored in Redis; this stub always allows requests.
func RateLimit() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if isPublicPath(string(c.Path())) {
			c.Next(ctx)
			return
		}
		tenantID := GetTenantID(c)
		if tenantID == "" {
			c.Next(ctx)
			return
		}
		// TODO: check Redis counter for tenant_id, reject with 429 if over limit
		allowed := checkLimit(tenantID)
		if !allowed {
			respondError(c, http.StatusTooManyRequests, "RATE_LIMIT_EXCEEDED",
				"request rate limit exceeded for this tenant")
			return
		}
		c.Next(ctx)
	}
}

func checkLimit(_ string) bool {
	// stub: always allow; replace with Redis token-bucket implementation
	return true
}
