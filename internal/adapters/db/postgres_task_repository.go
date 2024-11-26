package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/repository"
	"go.uber.org/zap"
)

type PostgresTaskRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewPostgresTaskRepository(db *sql.DB, logger *zap.Logger) repository.TaskRepository {
	return &PostgresTaskRepository{db: db, logger: logger}
}

func (r *PostgresTaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	query := `
		INSERT INTO tasks (title, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	r.logger.Info("Creating task", zap.String("title", task.Title))

	err := r.db.QueryRowContext(context.Background(), query, task.Title, task.Description, task.CreatedAt, task.UpdatedAt).Scan(&task.ID)
	if err != nil {
		r.logger.Error("Failed to create task", zap.Error(err))
		return nil, err
	}

	r.logger.Info("Task created successfully", zap.Int64("task_id", task.ID))
	return task, nil
}

func (r *PostgresTaskRepository) ListTasks() ([]*models.Task, error) {
	query := `SELECT id, title, description, created_at, updated_at FROM tasks`

	r.logger.Info("Fetching tasks")

	rows, err := r.db.QueryContext(context.Background(), query)
	if err != nil {
		r.logger.Error("Failed to fetch tasks", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt); err != nil {
			r.logger.Error("Failed to scan task row", zap.Error(err))
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	r.logger.Info("Tasks fetched successfully", zap.Int("task_count", len(tasks)))
	return tasks, nil
}

func (r *PostgresTaskRepository) GetTask(id int64) (*models.Task, error) {
	query := `SELECT id, title, description, created_at, updated_at FROM tasks WHERE id = $1`

	r.logger.Info("Fetching task", zap.Int64("task_id", id))

	var task models.Task
	err := r.db.QueryRowContext(context.Background(), query, id).Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Warn("Task not found", zap.Int64("task_id", id))
			return nil, nil
		}
		r.logger.Error("Failed to fetch task", zap.Error(err))
		return nil, err
	}

	r.logger.Info("Task fetched successfully", zap.Int64("task_id", task.ID))
	return &task, nil
}

func (r *PostgresTaskRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, updated_at = $3
		WHERE id = $4
		RETURNING id, created_at;
	`

	task.UpdatedAt = time.Now()

	r.logger.Info("Updating task", zap.Int64("task_id", task.ID))

	err := r.db.QueryRowContext(context.Background(), query, task.Title, task.Description, task.UpdatedAt, task.ID).
		Scan(&task.ID, &task.CreatedAt)
	if err != nil {
		r.logger.Error("Failed to update task", zap.Error(err))
		return nil, err
	}

	r.logger.Info("Task updated successfully", zap.Int64("task_id", task.ID))
	return task, nil
}

func (r *PostgresTaskRepository) DeleteTask(id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`

	r.logger.Info("Deleting task", zap.Int64("task_id", id))

	res, err := r.db.ExecContext(context.Background(), query, id)
	if err != nil {
		r.logger.Error("Failed to delete task", zap.Error(err))
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		r.logger.Error("Failed to get rows affected count", zap.Error(err))
		return err
	}

	if rowsAffected == 0 {
		r.logger.Warn("No rows deleted", zap.Int64("task_id", id))
		return sql.ErrNoRows
	}

	r.logger.Info("Task deleted successfully", zap.Int64("task_id", id))
	return nil
}
