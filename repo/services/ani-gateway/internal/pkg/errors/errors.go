// Package errors defines ANI's standard error response format.
// Every API error MUST use this format: {code, message, request_id, details}
package errors

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

// APIError is the canonical error response body for all ANI APIs.
type APIError struct {
	Code      string         `json:"code"`
	Message   string         `json:"message"`
	RequestID string         `json:"request_id"`
	Details   map[string]any `json:"details,omitempty"`
}

// Predefined error codes
const (
	CodeNotFound          = "NOT_FOUND"
	CodeUnauthorized      = "UNAUTHORIZED"
	CodeForbidden         = "FORBIDDEN"
	CodeBadRequest        = "BAD_REQUEST"
	CodeConflict          = "CONFLICT"
	CodeInternalError     = "INTERNAL_ERROR"
	CodeRateLimitExceeded = "RATE_LIMIT_EXCEEDED"
	CodeModelNotReady     = "MODEL_NOT_READY"
)

func RespondError(ctx context.Context, c *app.RequestContext, statusCode int, code, message string) {
	reqID, _ := c.Get("request_id")
	reqIDString, _ := reqID.(string)
	c.JSON(statusCode, APIError{
		Code:      code,
		Message:   message,
		RequestID: reqIDString,
	})
	c.Abort()
	_ = ctx
}

func NotFound(ctx context.Context, c *app.RequestContext, resource string) {
	RespondError(ctx, c, http.StatusNotFound, CodeNotFound, resource+" not found")
}

func Unauthorized(ctx context.Context, c *app.RequestContext) {
	RespondError(ctx, c, http.StatusUnauthorized, CodeUnauthorized, "authentication required")
}
