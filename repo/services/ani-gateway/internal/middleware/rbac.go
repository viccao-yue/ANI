package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// RBAC checks whether the authenticated user has permission to access the route.
// Permission is encoded as "{resource}:{action}", matched against the user's roles.
// This is a stub; production will call OPA or an internal RBAC service.
func RBAC() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if isPublicPath(string(c.Path())) {
			c.Next(ctx)
			return
		}
		tenantID := GetTenantID(c)
		if tenantID == "" {
			// Auth middleware should have already rejected unauthenticated requests.
			respondError(c, http.StatusForbidden, "FORBIDDEN", "tenant context missing")
			return
		}
		// TODO: extract required permission from route metadata, check against JWT roles
		c.Next(ctx)
	}
}
