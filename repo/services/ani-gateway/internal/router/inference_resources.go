package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerInferenceServices(svc *route.RouterGroup) {
	svc.GET("/inference-services", listInferenceServices)
	svc.POST("/inference-services", createInferenceService)
	svc.GET("/inference-services/:service_id", getInferenceService)
	svc.PATCH("/inference-services/:service_id", updateInferenceService)
	svc.DELETE("/inference-services/:service_id", deleteInferenceService)
	svc.GET("/inference-services/:service_id/logs", getInferenceServiceLogs)
}

func listInferenceServices(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createInferenceService(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func getInferenceService(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("service_id")})
	_ = ctx
}

func updateInferenceService(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("service_id")})
	_ = ctx
}

func deleteInferenceService(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func getInferenceServiceLogs(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}
