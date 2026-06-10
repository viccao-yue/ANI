// Package router contains shared stub handler for not-yet-implemented endpoints.
package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func notImplemented(ctx context.Context, c *app.RequestContext) {
	c.JSON(http.StatusNotImplemented, map[string]any{
		"code":       "NOT_IMPLEMENTED",
		"message":    "this endpoint is not yet implemented",
		"request_id": middleware.GetRequestID(c),
	})
	_ = ctx
}

// ── OpenAI-Compatible Inference Proxy ─────────────────────────────────────────

// inferenceProxy routes /v1/chat/completions to the correct vLLM service
// based on X-Model-Name header, then streams the response back via SSE.
func inferenceProxy(ctx context.Context, c *app.RequestContext) {
	// TODO: read X-Model-Name header → lookup endpoint_url in Redis/DB
	// → reverse-proxy to vLLM with SSE streaming
	notImplemented(ctx, c)
}
