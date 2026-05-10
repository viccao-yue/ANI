package main

import (
	"github.com/kubercloud/ani/pkg/bootstrap"
	"github.com/kubercloud/ani/services/model-service/internal/config"
	"github.com/kubercloud/ani/services/model-service/internal/repo"
	"github.com/kubercloud/ani/services/model-service/internal/service"
)

func main() {
	cfg := config.Load()
	deps := bootstrap.MustConnect(cfg)
	defer deps.Close()

	modelRepo := repo.NewPostgresModelRepo()
	svc := service.NewModelService(deps.DB, modelRepo)
	bootstrap.RunGRPC(cfg.GRPCPort, svc.Register, deps)
}
