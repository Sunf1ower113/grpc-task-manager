package repository

import "github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"

type TaskRepository interface {
	CreateTask(task *models.Task) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	GetTask(id int64) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(id int64) error
}
