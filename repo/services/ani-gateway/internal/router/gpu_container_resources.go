package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerGpuContainers(svc *route.RouterGroup) {
	svc.GET("/gpu-containers", listGpuContainers)
	svc.POST("/gpu-containers", createGpuContainer)
	// static sub-path before /:container_id to avoid wildcard match
	svc.GET("/gpu-containers/available-gpus", listAvailableGpus)
	svc.GET("/gpu-containers/:container_id", getGpuContainer)
	svc.PATCH("/gpu-containers/:container_id", updateGpuContainer)
	svc.DELETE("/gpu-containers/:container_id", deleteGpuContainer)
	svc.GET("/gpu-containers/:container_id/metrics", getGpuContainerMetrics)
	svc.GET("/gpu-containers/:container_id/versions", listGpuContainerVersions)
	svc.POST("/gpu-containers/:container_id/rollback", rollbackGpuContainer)
	svc.POST("/gpu-containers/:container_id/rebuild", rebuildGpuContainer)
}

func listGpuContainers(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createGpuContainer(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func listAvailableGpus(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}})
	_ = ctx
}

func getGpuContainer(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("container_id")})
	_ = ctx
}

func updateGpuContainer(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("container_id")})
	_ = ctx
}

func deleteGpuContainer(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func getGpuContainerMetrics(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"container_id": c.Param("container_id"), "metrics": []any{}})
	_ = ctx
}

func listGpuContainerVersions(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func rollbackGpuContainer(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("container_id"), "status": "rolling_back"})
	_ = ctx
}

func rebuildGpuContainer(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("container_id"), "status": "rebuilding"})
	_ = ctx
}
