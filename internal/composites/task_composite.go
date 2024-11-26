package composites

import (
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/repository"
	"go.uber.org/zap"
)

type TaskComposite struct {
	Repository repository.TaskRepository
}

func NewTaskComposite(repo repository.TaskRepository, logger *zap.Logger) (*TaskComposite, error) {
	logger.Info("Initializing TaskComposite")
	return &TaskComposite{
		Repository: repo,
	}, nil
}
