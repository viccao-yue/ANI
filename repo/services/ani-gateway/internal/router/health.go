package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

func registerHealth(r *route.RouterGroup) {
	r.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})
	r.GET("/ready", func(ctx context.Context, c *app.RequestContext) {
		// TODO: check DB, NATS, Redis connectivity
		c.JSON(http.StatusOK, map[string]string{"status": "ready"})
	})
}
