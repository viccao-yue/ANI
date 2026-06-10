package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerTasks(v1 *route.RouterGroup) {
	v1.GET("/tasks/:task_id", getTask)
	v1.DELETE("/tasks/:task_id", cancelTask)
}

func getTask(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{
		"id":     c.Param("task_id"),
		"status": "unknown",
	})
	_ = ctx
}

func cancelTask(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}
