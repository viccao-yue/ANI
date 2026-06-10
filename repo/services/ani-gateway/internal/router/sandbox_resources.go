package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerSandboxes(svc *route.RouterGroup) {
	svc.GET("/sandboxes", listSandboxes)
	svc.POST("/sandboxes", createSandbox)
	svc.GET("/sandboxes/:sandbox_id", getSandbox)
	svc.PATCH("/sandboxes/:sandbox_id", updateSandbox)
	svc.DELETE("/sandboxes/:sandbox_id", deleteSandbox)
	svc.POST("/sandboxes/:sandbox_id/extend", extendSandbox)
	svc.POST("/sandboxes/:sandbox_id/pause", pauseSandbox)
	svc.GET("/sandboxes/:sandbox_id/security-events", listSandboxSecurityEvents)
	svc.GET("/sandboxes/:sandbox_id/security-overview", getSandboxSecurityOverview)
}

func listSandboxes(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createSandbox(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func getSandbox(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("sandbox_id")})
	_ = ctx
}

func updateSandbox(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("sandbox_id")})
	_ = ctx
}

func deleteSandbox(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func extendSandbox(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("sandbox_id"), "status": "active"})
	_ = ctx
}

func pauseSandbox(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("sandbox_id"), "status": "paused"})
	_ = ctx
}

func listSandboxSecurityEvents(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func getSandboxSecurityOverview(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"sandbox_id": c.Param("sandbox_id"), "risk_level": "unknown"})
	_ = ctx
}
