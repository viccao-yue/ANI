package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	taskv1 "github.com/kubercloud/ani/pkg/generated/pb/task/v1"
	sharedrepo "github.com/kubercloud/ani/pkg/repo"
	"github.com/kubercloud/ani/pkg/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskService struct {
	taskv1.UnimplementedTaskServiceServer
	db   *pgxpool.Pool
	repo sharedrepo.AsyncTaskRepo
}

func NewTaskService(db *pgxpool.Pool, taskRepo sharedrepo.AsyncTaskRepo) *TaskService {
	return &TaskService{db: db, repo: taskRepo}
}

func (s *TaskService) Register(server *grpc.Server) {
	taskv1.RegisterTaskServiceServer(server, s)
}

func (s *TaskService) GetTask(ctx context.Context, req *taskv1.GetTaskRequest) (*taskv1.AsyncTask, error) {
	tenantID, taskID, err := parseTenantAndID(req.GetTenantId(), req.GetTaskId())
	if err != nil {
		return nil, toStatus(err)
	}
	ctx = types.WithTenant(ctx, &types.TenantContext{TenantID: tenantID})
	task, err := s.repo.GetByID(ctx, s.db, taskID)
	if err != nil {
		return nil, toStatus(err)
	}
	return taskToPB(task), nil
}

func (s *TaskService) CancelTask(ctx context.Context, req *taskv1.CancelTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "task cancellation requires explicit cancel state transition design")
}

func (s *TaskService) UpdateTaskProgress(ctx context.Context, req *taskv1.UpdateTaskProgressRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "worker task updates require tenant-aware internal auth context")
}

func (s *TaskService) AcquireTaskLease(ctx context.Context, req *taskv1.AcquireTaskLeaseRequest) (*taskv1.AcquireTaskLeaseResponse, error) {
	return nil, status.Error(codes.Unimplemented, "worker lease acquisition requires tenant-aware internal auth context")
}

func (s *TaskService) HeartbeatTaskLease(ctx context.Context, req *taskv1.HeartbeatTaskLeaseRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "worker heartbeat requires tenant-aware internal auth context")
}

func (s *TaskService) FailTask(ctx context.Context, req *taskv1.FailTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "worker failure reporting requires tenant-aware internal auth context")
}

func (s *TaskService) CompleteTask(ctx context.Context, req *taskv1.CompleteTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "worker completion reporting requires tenant-aware internal auth context")
}

func parseTenantAndID(tenantID, id string) (uuid.UUID, uuid.UUID, error) {
	tid, err := uuid.Parse(tenantID)
	if err != nil || tid == uuid.Nil {
		return uuid.Nil, uuid.Nil, types.Wrapf(types.ErrBadRequest, "invalid tenant_id")
	}
	parsedID, err := uuid.Parse(id)
	if err != nil || parsedID == uuid.Nil {
		return uuid.Nil, uuid.Nil, types.Wrapf(types.ErrBadRequest, "invalid task_id")
	}
	return tid, parsedID, nil
}

func taskToPB(task *sharedrepo.AsyncTask) *taskv1.AsyncTask {
	out := &taskv1.AsyncTask{
		TenantId:           task.TenantID.String(),
		Id:                 task.ID.String(),
		IdempotencyKey:     task.IdempotencyKey,
		TaskType:           task.TaskType,
		ResourceType:       task.ResourceType,
		ResourceId:         task.ResourceID.String(),
		Status:             task.Status,
		AttemptCount:       int32(task.AttemptCount),
		MaxAttempts:        int32(task.MaxAttempts),
		ProgressPct:        int32(task.ProgressPct),
		ErrorMessage:       task.ErrorMessage,
		CompensatingAction: task.CompensatingAction,
		WebhookUrl:         task.WebhookURL,
		CreatedAt:          timestamppb.New(task.CreatedAt),
	}
	if task.LeaseUntil != nil {
		out.LeaseUntil = timestamppb.New(*task.LeaseUntil)
	}
	if task.LastHeartbeatAt != nil {
		out.LastHeartbeatAt = timestamppb.New(*task.LastHeartbeatAt)
	}
	if task.DeadLetterAt != nil {
		out.DeadLetterAt = timestamppb.New(*task.DeadLetterAt)
	}
	if task.StartedAt != nil {
		out.StartedAt = timestamppb.New(*task.StartedAt)
	}
	if task.CompletedAt != nil {
		out.CompletedAt = timestamppb.New(*task.CompletedAt)
	}
	return out
}

func toStatus(err error) error {
	switch {
	case errors.Is(err, types.ErrNotFound):
		return status.Error(codes.NotFound, "task not found")
	case errors.Is(err, types.ErrBadRequest), errors.Is(err, types.ErrInvalidState):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, types.ErrForbidden):
		return status.Error(codes.PermissionDenied, "forbidden")
	case errors.Is(err, types.ErrUnauthorized):
		return status.Error(codes.Unauthenticated, "unauthorized")
	default:
		return status.Error(codes.Internal, "internal error")
	}
}
