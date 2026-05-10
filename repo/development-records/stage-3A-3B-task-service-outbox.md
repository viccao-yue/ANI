# Stage 3A + 3B Development Record

Date: 2026-05-11

Scope accepted by product owner:
- Stage 3A: implement the minimal `task-service` query interface.
- Stage 3B: implement the outbox publisher loop.

Design sources:
- `ANI-11-代码实现规范.md`
- `ANI-09-数据模型设计.md`
- `ANI-06-开发计划.md`
- `api/proto/task/v1/task_service.proto`

## Implemented

Stage 3A:
- Added `services/task-service`.
- Added environment-based config loading.
- Added gRPC registration through shared `bootstrap.RunGRPC`.
- Implemented `TaskService.GetTask`.
- `GetTask` sets tenant context before reading `async_tasks`, so PostgreSQL RLS remains active.
- Worker mutation RPCs remain explicitly `Unimplemented` because current proto requests do not carry tenant context for safe RLS-backed writes.

Stage 3B:
- Added shared `pkg/repo.OutboxRepo`.
- Implemented `FetchUnpublished` using `FOR UPDATE SKIP LOCKED`.
- Implemented `MarkPublished`.
- Added `task-service/internal/worker.OutboxPublisher`.
- Publisher publishes to NATS JetStream first, then marks rows as published in the same DB transaction.
- If NATS publish fails, the transaction rolls back and the event remains unpublished.

## Files Added

- `pkg/repo/outbox_repo.go`
- `services/task-service/go.mod`
- `services/task-service/main.go`
- `services/task-service/internal/config/config.go`
- `services/task-service/internal/service/task_service.go`
- `services/task-service/internal/worker/outbox_publisher.go`
- `development-records/stage-3A-3B-task-service-outbox.md`

## Files Updated

- `go.work`
- `Makefile`

## Validation

Required commands:

```bash
make gen-proto
make test
make build
```

Result:
- `make gen-proto`: passed.
- `make test`: passed.
- `make build`: passed.

Generated binaries:
- `bin/ani-gateway`
- `bin/model-service`
- `bin/task-service`

## Deferred

- `CancelTask` state transition.
- Worker mutation RPCs: `AcquireTaskLease`, `HeartbeatTaskLease`, `UpdateTaskProgress`, `FailTask`, `CompleteTask`.
- These need a tenant-aware internal authentication design or proto changes before implementation.
