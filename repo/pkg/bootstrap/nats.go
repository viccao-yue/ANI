package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/nats-io/nats.go"
)

// connectNATS connects to NATS and enables JetStream.
// Retries for up to 30 seconds.
func connectNATS(natsURL string) (*nats.Conn, nats.JetStreamContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var nc *nats.Conn
	var err error

	for {
		nc, err = nats.Connect(natsURL,
			nats.RetryOnFailedConnect(true),
			nats.MaxReconnects(5),
			nats.ReconnectWait(2*time.Second),
			nats.Name("ani-service"),
		)
		if err == nil {
			break
		}

		select {
		case <-ctx.Done():
			return nil, nil, fmt.Errorf("nats connection timeout: %w", err)
		case <-time.After(2 * time.Second):
			slog.Warn("nats not ready, retrying...", "err", err)
		}
	}

	js, err := nc.JetStream()
	if err != nil {
		nc.Close()
		return nil, nil, fmt.Errorf("nats jetstream init: %w", err)
	}

	// Ensure core ANI streams exist (idempotent)
	if err := ensureStreams(js); err != nil {
		nc.Close()
		return nil, nil, fmt.Errorf("nats stream setup: %w", err)
	}

	slog.Info("nats connected", "url", natsURL)
	return nc, js, nil
}

func ensureStreams(js nats.JetStreamContext) error {
	streams := []struct {
		name     string
		subjects []string
	}{
		{
			name:     "ANI_TASKS",
			subjects: []string{"ani.tasks.>"},
		},
		{
			name:     "ANI_EVENTS",
			subjects: []string{"ani.events.>"},
		},
	}

	for _, s := range streams {
		_, err := js.StreamInfo(s.name)
		if err == nats.ErrStreamNotFound {
			_, err = js.AddStream(&nats.StreamConfig{
				Name:      s.name,
				Subjects:  s.subjects,
				Retention: nats.WorkQueuePolicy, // consumed messages are deleted
				MaxAge:    24 * time.Hour,
				Storage:   nats.FileStorage,
				Replicas:  1, // increase to 3 in production HA deployments
			})
		}
		if err != nil {
			return fmt.Errorf("stream %s: %w", s.name, err)
		}
	}
	return nil
}
