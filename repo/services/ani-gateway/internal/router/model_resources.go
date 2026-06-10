package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerModels(svc *route.RouterGroup) {
	svc.GET("/models", listModels)
	svc.POST("/models", createModel)
	svc.POST("/models/import", importModel)
	svc.GET("/models/:model_id", getModel)
	svc.DELETE("/models/:model_id", deleteModel)
	svc.GET("/models/:model_id/versions", listModelVersions)
	svc.POST("/models/:model_id/versions", createModelVersion)
}

func listModels(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createModel(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func importModel(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func getModel(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("model_id")})
	_ = ctx
}

func deleteModel(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func listModelVersions(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createModelVersion(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "model_id": c.Param("model_id"), "status": "pending"})
	_ = ctx
}
