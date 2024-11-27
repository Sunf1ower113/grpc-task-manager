package services

import (
	"database/sql"
	"errors"

	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/repository"
	"go.uber.org/zap"
)

var (
	ErrTaskNotFound   = errors.New("task not found")
	ErrTaskCreateFail = errors.New("failed to create task")
	ErrTaskUpdateFail = errors.New("failed to update task")
	ErrTaskDeleteFail = errors.New("failed to delete task")
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

	createdTask, err := s.repo.CreateTask(task)
	if err != nil {
		s.logger.Error("Failed to create task", zap.Error(err))
		return nil, ErrTaskCreateFail
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
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warn("Task not found", zap.Int64("id", id))
			return nil, ErrTaskNotFound
		}
		s.logger.Error("Failed to fetch task", zap.Error(err))
		return nil, err
	}

	return task, nil
}

func (s *taskService) UpdateTask(task *models.Task) (*models.Task, error) {
	s.logger.Info("Updating task", zap.Int64("id", task.ID))

	updatedTask, err := s.repo.UpdateTask(task)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warn("Task not found for update", zap.Int64("id", task.ID))
			return nil, ErrTaskNotFound
		}
		s.logger.Error("Failed to update task", zap.Error(err))
		return nil, ErrTaskUpdateFail
	}

	return updatedTask, nil
}

func (s *taskService) DeleteTask(id int64) error {
	s.logger.Info("Deleting task", zap.Int64("id", id))

	err := s.repo.DeleteTask(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warn("Task not found for deletion", zap.Int64("id", id))
			return ErrTaskNotFound
		}
		s.logger.Error("Failed to delete task", zap.Error(err))
		return ErrTaskDeleteFail
	}

	return nil
}
