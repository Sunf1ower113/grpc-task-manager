package services

import (
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/repository"
	"go.uber.org/zap"
)

type TaskService interface {
	CreateTask(task *models.Task) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	GetTask(id int64) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(id int64) error
}

type taskService struct {
	repo   repository.TaskRepository
	logger *zap.Logger
}

func NewTaskService(repo repository.TaskRepository, logger *zap.Logger) TaskService {
	return &taskService{
		repo:   repo,
		logger: logger,
	}
}

func (s *taskService) CreateTask(task *models.Task) (*models.Task, error) {
	s.logger.Info("Creating task", zap.String("title", task.Title))

	if task.Title == "" {
		s.logger.Warn("Task validation failed: title is empty")
		return nil, ErrInvalidInput("Title cannot be empty")
	}

	createdTask, err := s.repo.CreateTask(task)
	if err != nil {
		s.logger.Error("Failed to create task", zap.Error(err))
		return nil, err
	}
	return createdTask, nil
}

func (s *taskService) ListTasks() ([]*models.Task, error) {
	s.logger.Info("Listing tasks")
	tasks, err := s.repo.ListTasks()
	if err != nil {
		s.logger.Error("Failed to list tasks", zap.Error(err))
		return nil, err
	}
	return tasks, nil
}

func (s *taskService) GetTask(id int64) (*models.Task, error) {
	s.logger.Info("Fetching task", zap.Int64("id", id))
	task, err := s.repo.GetTask(id)
	if err != nil {
		s.logger.Error("Failed to fetch task", zap.Error(err))
		return nil, err
	}
	if task == nil {
		s.logger.Warn("Task not found", zap.Int64("id", id))
		return nil, ErrNotFound("Task not found")
	}
	return task, nil
}

func (s *taskService) UpdateTask(task *models.Task) (*models.Task, error) {
	s.logger.Info("Updating task", zap.Int64("id", task.ID))
	if task.ID == 0 {
		s.logger.Warn("Task validation failed: ID is zero")
		return nil, ErrInvalidInput("Task ID cannot be zero")
	}
	updatedTask, err := s.repo.UpdateTask(task)
	if err != nil {
		s.logger.Error("Failed to update task", zap.Error(err))
		return nil, err
	}
	return updatedTask, nil
}

func (s *taskService) DeleteTask(id int64) error {
	s.logger.Info("Deleting task", zap.Int64("id", id))
	err := s.repo.DeleteTask(id)
	if err != nil {
		s.logger.Error("Failed to delete task", zap.Error(err))
		return err
	}
	return nil
}
