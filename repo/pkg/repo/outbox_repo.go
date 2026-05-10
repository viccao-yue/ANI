package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// OutboxEvent is a pending event that must be published to NATS.
type OutboxEvent struct {
	ID            int64
	AggregateType string
	AggregateID   uuid.UUID
	EventType     string
	TenantID      uuid.UUID
	Payload       []byte
	CreatedAt     time.Time
}

// OutboxRepo manages unpublished outbox events.
//
// The publisher should connect with a DB role allowed to scan outbox_events
// across tenants, as documented by ANI-09. It intentionally does not set
// app.current_tenant_id because outbox publication is a platform-internal
// cross-tenant operation.
type OutboxRepo interface {
	FetchUnpublished(ctx context.Context, tx pgx.Tx, limit int) ([]OutboxEvent, error)
	MarkPublished(ctx context.Context, tx pgx.Tx, ids []int64) error
}

type PostgresOutboxRepo struct{}

func NewPostgresOutboxRepo() *PostgresOutboxRepo {
	return &PostgresOutboxRepo{}
}

func (r *PostgresOutboxRepo) FetchUnpublished(ctx context.Context, tx pgx.Tx, limit int) ([]OutboxEvent, error) {
	if limit <= 0 || limit > 500 {
		limit = 100
	}
	rows, err := tx.Query(ctx, `
		SELECT id, aggregate_type, aggregate_id, event_type, tenant_id, payload, created_at
		FROM outbox_events
		WHERE NOT published
		ORDER BY created_at ASC, id ASC
		LIMIT $1
		FOR UPDATE SKIP LOCKED
	`, limit)
	if err != nil {
		return nil, fmt.Errorf("outboxRepo.FetchUnpublished query: %w", err)
	}
	defer rows.Close()

	var events []OutboxEvent
	for rows.Next() {
		var event OutboxEvent
		if err := rows.Scan(
			&event.ID,
			&event.AggregateType,
			&event.AggregateID,
			&event.EventType,
			&event.TenantID,
			&event.Payload,
			&event.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("outboxRepo.FetchUnpublished scan: %w", err)
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("outboxRepo.FetchUnpublished rows: %w", err)
	}
	return events, nil
}

func (r *PostgresOutboxRepo) MarkPublished(ctx context.Context, tx pgx.Tx, ids []int64) error {
	if len(ids) == 0 {
		return nil
	}
	tag, err := tx.Exec(ctx, `
		UPDATE outbox_events
		SET published=true, published_at=NOW()
		WHERE id = ANY($1)
	`, ids)
	if err != nil {
		return fmt.Errorf("outboxRepo.MarkPublished update: %w", err)
	}
	if tag.RowsAffected() != int64(len(ids)) {
		return fmt.Errorf("outboxRepo.MarkPublished updated %d rows, expected %d", tag.RowsAffected(), len(ids))
	}
	return nil
}

func BeginOutboxTx(ctx context.Context, pool *pgxpool.Pool) (pgx.Tx, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("begin outbox tx: %w", err)
	}
	return tx, nil
}
