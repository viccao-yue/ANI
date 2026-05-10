// Package types defines shared domain types used across all ANI services.
package types

import (
	"errors"
	"fmt"
)

// Sentinel errors — use errors.Is() to check, never == comparison.
var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("already exists")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrBadRequest   = errors.New("bad request")
	ErrInvalidState = errors.New("invalid state transition")
	ErrLeaseTaken   = errors.New("task lease already held by another worker")
	ErrDeadLetter   = errors.New("task exceeded max attempts and is in dead-letter queue")
)

// ANIError carries a stable code, a user-facing message, and an optional
// cause. The cause is logged server-side but never sent to clients.
type ANIError struct {
	Code    string // e.g. "MODEL_NOT_FOUND" — safe to return to clients
	Message string // human-readable, safe to return to clients
	Cause   error  // underlying error, logged only
}

func (e *ANIError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *ANIError) Unwrap() error { return e.Cause }

// Error code constants — every code here must appear in OpenAPI ErrorResponse.
const (
	CodeNotFound          = "NOT_FOUND"
	CodeConflict          = "CONFLICT"
	CodeUnauthorized      = "UNAUTHORIZED"
	CodeForbidden         = "FORBIDDEN"
	CodeBadRequest        = "BAD_REQUEST"
	CodeRateLimitExceeded = "RATE_LIMIT_EXCEEDED"
	CodeInvalidState      = "INVALID_STATE"
	CodeInternalError     = "INTERNAL_ERROR"
	CodeNotImplemented    = "NOT_IMPLEMENTED"
)

// HTTP status codes for each sentinel error, used by the gateway error handler.
func HTTPStatusForError(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return 404
	case errors.Is(err, ErrConflict):
		return 409
	case errors.Is(err, ErrUnauthorized):
		return 401
	case errors.Is(err, ErrForbidden):
		return 403
	case errors.Is(err, ErrBadRequest), errors.Is(err, ErrInvalidState):
		return 400
	default:
		return 500
	}
}

// Wrapf wraps an error with context for logging while preserving sentinel type.
// Usage: return types.Wrapf(ErrNotFound, "inferenceRepo.GetByID id=%s", id)
func Wrapf(sentinel error, format string, args ...any) error {
	return fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), sentinel)
}
