// Package bootstrap provides standardized dependency initialization for ANI services.
// Every Go microservice calls bootstrap.MustConnect() to get a *Deps,
// then passes it to bootstrap.RunGRPC() to start serving.
package bootstrap

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

// Deps holds all initialized external dependencies.
// All fields are non-nil after MustConnect returns successfully.
type Deps struct {
	DB     *pgxpool.Pool
	NATS   *nats.Conn
	JS     nats.JetStreamContext
	Redis  *redis.Client
	Logger *slog.Logger
}

// Close releases all connections. Call with defer after MustConnect.
func (d *Deps) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.NATS != nil {
		d.NATS.Close()
	}
	if d.Redis != nil {
		_ = d.Redis.Close()
	}
}
