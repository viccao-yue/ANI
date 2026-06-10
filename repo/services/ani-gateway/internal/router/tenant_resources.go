package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerTenant(svc *route.RouterGroup) {
	svc.GET("/tenant/members", listTenantMembers)
	svc.POST("/tenant/members", inviteTenantMember)
	svc.DELETE("/tenant/members/:member_id", removeTenantMember)
	svc.GET("/tenant/roles", listTenantRoles)
	svc.GET("/tenant/sso", getTenantSSO)
	svc.PUT("/tenant/sso", updateTenantSSO)
	svc.GET("/tenant/webhooks", listTenantWebhooks)
	svc.POST("/tenant/webhooks", createTenantWebhook)
	svc.DELETE("/tenant/webhooks/:webhook_id", deleteTenantWebhook)
}

func listTenantMembers(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func inviteTenantMember(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "invited"})
	_ = ctx
}

func removeTenantMember(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func listTenantRoles(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}})
	_ = ctx
}

func getTenantSSO(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"enabled": false})
	_ = ctx
}

func updateTenantSSO(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"enabled": false})
	_ = ctx
}

func listTenantWebhooks(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createTenantWebhook(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "active"})
	_ = ctx
}

func deleteTenantWebhook(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}
