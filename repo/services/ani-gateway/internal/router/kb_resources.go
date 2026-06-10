package router

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/kubercloud/ani/services/ani-gateway/internal/middleware"
)

func registerKnowledgeBases(svc *route.RouterGroup) {
	svc.GET("/knowledge-bases", listKnowledgeBases)
	svc.POST("/knowledge-bases", createKnowledgeBase)
	svc.GET("/knowledge-bases/:kb_id", getKnowledgeBase)
	svc.DELETE("/knowledge-bases/:kb_id", deleteKnowledgeBase)
	svc.GET("/knowledge-bases/:kb_id/documents", listKnowledgeBaseDocuments)
	svc.POST("/knowledge-bases/:kb_id/documents", uploadKnowledgeBaseDocument)
	svc.DELETE("/knowledge-bases/:kb_id/documents/:doc_id", deleteKnowledgeBaseDocument)
	svc.POST("/knowledge-bases/:kb_id/query", queryKnowledgeBase)
	// SSE streaming query (separate endpoint from JSON to allow clean SDK generation)
	svc.GET("/knowledge-bases/:kb_id/query/stream", streamQueryKnowledgeBase)
}

func listKnowledgeBases(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func createKnowledgeBase(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func getKnowledgeBase(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": c.Param("kb_id")})
	_ = ctx
}

func deleteKnowledgeBase(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func listKnowledgeBaseDocuments(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"items": []any{}, "next_cursor": nil})
	_ = ctx
}

func uploadKnowledgeBaseDocument(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"id": "", "status": "pending"})
	_ = ctx
}

func deleteKnowledgeBaseDocument(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.Status(http.StatusNoContent)
	_ = ctx
}

func queryKnowledgeBase(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	c.JSON(http.StatusOK, map[string]any{"results": []any{}, "total": 0})
	_ = ctx
}

func streamQueryKnowledgeBase(ctx context.Context, c *app.RequestContext) {
	_ = middleware.GetTenantID(c)
	// SSE: write headers then end stream immediately
	c.Response.Header.Set("Content-Type", "text/event-stream")
	c.Response.Header.Set("Cache-Control", "no-cache")
	c.Response.SetStatusCode(http.StatusOK)
	_ = ctx
}
