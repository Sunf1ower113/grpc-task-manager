package grpc

import (
	"context"
	"errors"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/services"
	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TaskHandler implements the gRPC TaskManagerServer interface and handles all task-related gRPC requests.
type TaskHandler struct {
	pb.UnimplementedTaskManagerServer
	service services.TaskService
	logger  *zap.Logger
}

// NewTaskHandler initializes a new TaskHandler instance.
func NewTaskHandler(service services.TaskService, logger *zap.Logger) pb.TaskManagerServer {
	return &TaskHandler{
		service: service,
		logger:  logger,
	}
}

// CreateTask handles the gRPC request to create a new task.
func (h *TaskHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	h.logger.Info("Received CreateTask request", zap.String("title", req.Title))

	if err := trimAndValidateCreateTaskRequest(req); err != nil {
		h.logger.Warn("Validation failed for CreateTask", zap.Error(err))
		return nil, err
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	createdTask, err := h.service.CreateTask(task)
	if err != nil {
		h.logger.Error("Failed to create task", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to create task")
	}

	return &pb.TaskResponse{
		Id:          createdTask.ID,
		Title:       createdTask.Title,
		Description: createdTask.Description,
		CreatedAt:   createdTask.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   createdTask.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// ListTasks handles the gRPC request to list all tasks.
func (h *TaskHandler) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	h.logger.Info("Received ListTasks request")

	tasks, err := h.service.ListTasks()
	if err != nil {
		h.logger.Error("Failed to list tasks", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to list tasks")
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

// GetTask handles the gRPC request to retrieve a specific task by its ID.
func (h *TaskHandler) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.TaskResponse, error) {
	h.logger.Info("Received GetTask request", zap.Int64("id", req.Id))

	if err := trimAndValidateGetTaskRequest(req); err != nil {
		h.logger.Warn("Validation failed for GetTask", zap.Error(err))
		return nil, err
	}

	task, err := h.service.GetTask(req.Id)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			h.logger.Warn("Task not found", zap.Int64("id", req.Id))
			return nil, status.Error(codes.NotFound, "Task not found")
		}
		h.logger.Error("Failed to fetch task", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to fetch task")
	}

	return &pb.TaskResponse{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// UpdateTask handles the gRPC request to update an existing task.
func (h *TaskHandler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.TaskResponse, error) {
	h.logger.Info("Received UpdateTask request", zap.Int64("id", req.Id))

	if err := trimAndValidateUpdateTaskRequest(req); err != nil {
		h.logger.Warn("Validation failed for UpdateTask", zap.Error(err))
		return nil, err
	}

	task := &models.Task{
		ID:          req.Id,
		Title:       req.Title,
		Description: req.Description,
	}

	updatedTask, err := h.service.UpdateTask(task)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			h.logger.Warn("Task not found for update", zap.Int64("id", req.Id))
			return nil, status.Error(codes.NotFound, "Task not found for update")
		}
		h.logger.Error("Failed to update task", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to update task")
	}

	return &pb.TaskResponse{
		Id:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		CreatedAt:   updatedTask.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   updatedTask.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

// DeleteTask handles the gRPC request to delete a task by its ID.
func (h *TaskHandler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	h.logger.Info("Received DeleteTask request", zap.Int64("id", req.Id))

	if err := trimAndValidateDeleteTaskRequest(req); err != nil {
		h.logger.Warn("Validation failed for DeleteTask", zap.Error(err))
		return nil, err
	}

	err := h.service.DeleteTask(req.Id)
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			h.logger.Warn("Task not found for deletion", zap.Int64("id", req.Id))
			return nil, status.Error(codes.NotFound, "Task not found for deletion")
		}
		h.logger.Error("Failed to delete task", zap.Error(err))
		return nil, status.Error(codes.Internal, "Failed to delete task")
	}

	return &pb.DeleteTaskResponse{Success: true}, nil
}
