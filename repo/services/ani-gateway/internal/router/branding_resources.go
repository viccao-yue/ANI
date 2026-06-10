package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerBranding(v1 *route.RouterGroup) {
	v1.GET("/branding", getBranding)
	v1.PUT("/branding", updateBranding)
	v1.POST("/branding/logo", uploadBrandingLogo)
}

func getBranding(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusOK, map[string]any{
		"logo_url":      "",
		"primary_color": "#000000",
		"name":          "",
	})
	_ = ctx
}

func updateBranding(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{
		"logo_url":      "",
		"primary_color": "#000000",
		"name":          "",
	})
	_ = ctx
}

func uploadBrandingLogo(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"logo_url": ""})
	_ = ctx
}
