package grpc

import (
	"context"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/services"
	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"go.uber.org/zap"
)

type TaskHandler struct {
	pb.UnimplementedTaskManagerServer
	service services.TaskService
	logger  *zap.Logger
}

func NewTaskHandler(service services.TaskService, logger *zap.Logger) pb.TaskManagerServer {
	return &TaskHandler{
		service: service,
		logger:  logger,
	}
}

func (h *TaskHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	h.logger.Info("Received CreateTask request", zap.String("title", req.Title))

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	createdTask, err := h.service.CreateTask(task)
	if err != nil {
		h.logger.Error("Failed to create task", zap.Error(err))
		return nil, err
	}

	return &pb.TaskResponse{
		Id:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		CreatedAt:   createdTask.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   createdTask.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (h *TaskHandler) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	h.logger.Info("Received ListTasks request")

	tasks, err := h.service.ListTasks()
	if err != nil {
		h.logger.Error("Failed to list tasks", zap.Error(err))
		return nil, err
	}

	var taskResponses []*pb.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, &pb.TaskResponse{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	return &pb.ListTasksResponse{Tasks: taskResponses}, nil
}

func (h *TaskHandler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	h.logger.Info("Received GetTask request", zap.Int64("id", req.Id))

	task, err := h.service.GetTask(req.Id)
	if err != nil {
		h.logger.Error("Failed to get task", zap.Error(err))
		return nil, err
	}
	if task == nil {
		h.logger.Warn("Task not found", zap.Int64("id", req.Id))
		return nil, ErrTaskNotFound("Task not found")
	}

	return &pb.TaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (h *TaskHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	h.logger.Info("Received UpdateTask request", zap.Int64("id", req.Id))

	task := &models.Task{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
	}

	updatedTask, err := h.service.UpdateTask(task)
	if err != nil {
		h.logger.Error("Failed to update task", zap.Error(err))
		return nil, err
	}

	return &pb.TaskResponse{
		Id:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		CreatedAt:   updatedTask.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   updatedTask.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	h.logger.Info("Received DeleteTask request", zap.Int64("id", req.Id))

	err := h.service.DeleteTask(req.Id)
	if err != nil {
		h.logger.Error("Failed to delete task", zap.Error(err))
		return nil, err
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}
