package composites

import (
	"database/sql"
	storage "github.com/Sunf1ower113/grpc-task-manager/internal/adapters/db"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/repository"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/services"
	"go.uber.org/zap"
)

type TaskComposite struct {
	Repository repository.TaskRepository
	Service    services.TaskService
}

func NewTaskComposite(db *sql.DB, logger *zap.Logger) (*TaskComposite, error) {
	taskRepository := storage.NewPostgresTaskRepository(db, logger)
	taskService := services.NewTaskService(taskRepository, logger)
	return &TaskComposite{
		Repository: taskRepository,
		Service:    taskService,
	}, nil
}
