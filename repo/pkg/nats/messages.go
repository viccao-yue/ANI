// Package nats defines the canonical payload types for every ANI NATS subject.
//
// RULE: All publishers and consumers MUST use these structs.
// Never define task payload structs in individual services.
//
// Subject naming convention:
//
//	ani.tasks.{domain}.{action}   — task messages (WorkQueue, at-least-once)
//	ani.events.task.{id}.{event}  — task lifecycle events (Interest, fan-out)
package nats

import (
	"time"

	"github.com/google/uuid"
)

// ── Subject constants ──────────────────────────────────────────────────────────

const (
	SubjectInferenceDeploy = "ani.tasks.inference.deploy"
	SubjectInferenceDelete = "ani.tasks.inference.delete"
	SubjectKBParse         = "ani.tasks.kb.parse"
	SubjectKBIndex         = "ani.tasks.kb.index"
	SubjectModelImport     = "ani.tasks.model.import"
	SubjectTaskCompleted   = "ani.events.task.completed" // append ".{task_id}" when publishing
)

// ── Inference payloads ────────────────────────────────────────────────────────

// InferenceDeployMsg is published to SubjectInferenceDeploy when a user
// requests a new InferenceService. The Operator consumes this and creates the K8s CRD.
type InferenceDeployMsg struct {
	TaskID         uuid.UUID `json:"task_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	TenantID       uuid.UUID `json:"tenant_id"`
	ServiceID      uuid.UUID `json:"service_id"`
	ModelVersionID uuid.UUID `json:"model_version_id"`
	IsEncrypted    bool      `json:"is_encrypted"`
	EncryptAlgo    string    `json:"encrypt_algo,omitempty"` // sm4 | zuc | aes256gcm
	GPUType        string    `json:"gpu_type,omitempty"`
	GPUCount       int32     `json:"gpu_count"`
	Replicas       int32     `json:"replicas"`
	MaxConcurrency int32     `json:"max_concurrency"`
	DrainSeconds   int32     `json:"drain_seconds"`
	PublishedAt    time.Time `json:"published_at"`
}

// InferenceDeleteMsg is published to SubjectInferenceDelete when a user
// deletes an InferenceService. The Operator drains requests then removes K8s resources.
type InferenceDeleteMsg struct {
	TaskID       uuid.UUID `json:"task_id"`
	TenantID     uuid.UUID `json:"tenant_id"`
	ServiceID    uuid.UUID `json:"service_id"`
	DrainSeconds int32     `json:"drain_seconds"`
	PublishedAt  time.Time `json:"published_at"`
}

// ── Knowledge base payloads ───────────────────────────────────────────────────

// KBParseMsg is published to SubjectKBParse after a document is uploaded.
// The doc-parser service consumes this and extracts text chunks.
type KBParseMsg struct {
	TaskID         uuid.UUID `json:"task_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	TenantID       uuid.UUID `json:"tenant_id"`
	KBID           uuid.UUID `json:"kb_id"`
	DocID          uuid.UUID `json:"doc_id"`
	StoragePath    string    `json:"storage_path"` // MinIO object path
	FileType       string    `json:"file_type"`    // pdf | docx | xlsx | txt | md
	FileSizeBytes  int64     `json:"file_size_bytes"`
	ChecksumSHA256 string    `json:"checksum_sha256"`
	PublishedAt    time.Time `json:"published_at"`
}

// ParsedChunk is a single text chunk produced by the doc-parser.
type ParsedChunk struct {
	Index      int    `json:"index"`
	Content    string `json:"content"`
	PageNumber int    `json:"page_number,omitempty"` // for PDFs
	Heading    string `json:"heading,omitempty"`     // section heading if available
}

// KBIndexMsg is published to SubjectKBIndex after parsing completes.
// The rag-engine service consumes this, embeds chunks, and upserts to Milvus.
type KBIndexMsg struct {
	TaskID         uuid.UUID     `json:"task_id"`
	IdempotencyKey string        `json:"idempotency_key"`
	TenantID       uuid.UUID     `json:"tenant_id"`
	KBID           uuid.UUID     `json:"kb_id"`
	DocID          uuid.UUID     `json:"doc_id"`
	FileName       string        `json:"file_name"`
	Chunks         []ParsedChunk `json:"chunks"`
	PublishedAt    time.Time     `json:"published_at"`
}

// ── Model import payloads ─────────────────────────────────────────────────────

// ModelImportMsg is published to SubjectModelImport when a user requests
// importing a model from HuggingFace or ModelScope.
type ModelImportMsg struct {
	TaskID         uuid.UUID `json:"task_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	TenantID       uuid.UUID `json:"tenant_id"`
	ModelID        uuid.UUID `json:"model_id"`
	Source         string    `json:"source"`        // huggingface | modelscope
	RepoID         string    `json:"repo_id"`       // e.g. "Qwen/Qwen2.5-72B-Instruct"
	Revision       string    `json:"revision"`      // branch / tag / commit
	TargetBucket   string    `json:"target_bucket"` // MinIO bucket
	TargetPrefix   string    `json:"target_prefix"` // MinIO object prefix
	WebhookURL     string    `json:"webhook_url,omitempty"`
	PublishedAt    time.Time `json:"published_at"`
}

// ── Task lifecycle events ─────────────────────────────────────────────────────

// TaskCompletedEvent is published to "ani.events.task.completed.{task_id}"
// when any async task finishes (success or failure).
// Consumers: ANI Gateway SSE pusher, Webhook dispatcher.
type TaskCompletedEvent struct {
	TaskID      uuid.UUID `json:"task_id"`
	TenantID    uuid.UUID `json:"tenant_id"`
	TaskType    string    `json:"task_type"`
	Status      string    `json:"status"`              // completed | failed | dead_letter
	Result      any       `json:"result,omitempty"`    // task-specific result payload
	ErrorMsg    string    `json:"error_msg,omitempty"` // present when status=failed
	PublishedAt time.Time `json:"published_at"`
}

// TaskCompletedSubject returns the per-task event subject for a given task ID.
func TaskCompletedSubject(taskID uuid.UUID) string {
	return SubjectTaskCompleted + "." + taskID.String()
}
