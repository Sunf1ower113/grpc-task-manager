package composites

import (
	"database/sql"
	"errors"

	storage "github.com/Sunf1ower113/grpc-task-manager/internal/adapters/db"
	"github.com/Sunf1ower113/grpc-task-manager/internal/adapters/grpc"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/repository"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/services"
	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"go.uber.org/zap"
)

type TaskComposite struct {
	Repository repository.TaskRepository
	Service    services.TaskService
	Handler    pb.TaskManagerServer
}

func NewTaskComposite(db *sql.DB, logger *zap.Logger) (*TaskComposite, error) {
	taskRepository := storage.NewPostgresTaskRepository(db, logger)
	if taskRepository == nil {
		return nil, errors.New("failed to initialize task repository")
	}

	taskService := services.NewTaskService(taskRepository, logger)

	if taskService == nil {
		return nil, errors.New("failed to initialize task service")
	}
	taskHandler := grpc.NewTaskHandler(taskService, logger)
	if taskHandler == nil {
		return nil, errors.New("failed to initialize task handler")
	}

	return &TaskComposite{
		Repository: taskRepository,
		Service:    taskService,
		Handler:    taskHandler,
	}, nil
}
