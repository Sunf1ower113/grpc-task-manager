package db

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *zap.Logger) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	logger := zap.NewNop()
	return db, mock, logger
}

func TestPostgresTaskRepository_CreateTask(t *testing.T) {
	db, mock, logger := setupMockDB(t)
	defer db.Close()

	repo := NewPostgresTaskRepository(db, logger)
	task := &models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
	}

	now := time.Now()
	mock.ExpectQuery("INSERT INTO tasks").
		WithArgs(task.Title, task.Description, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	createdTask, err := repo.CreateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, createdTask)
	assert.Equal(t, int64(1), createdTask.ID)
	assert.WithinDuration(t, now, createdTask.CreatedAt, time.Second)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresTaskRepository_ListTasks(t *testing.T) {
	db, mock, logger := setupMockDB(t)
	defer db.Close()

	repo := NewPostgresTaskRepository(db, logger)

	mock.ExpectQuery("SELECT id, title, description, created_at, updated_at FROM tasks").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "created_at", "updated_at"}).
			AddRow(1, "Test Task", "This is a test task", time.Now(), time.Now()).
			AddRow(2, "Another Task", "This is another test task", time.Now(), time.Now()))

	tasks, err := repo.ListTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresTaskRepository_GetTask(t *testing.T) {
	db, mock, logger := setupMockDB(t)
	defer db.Close()

	repo := NewPostgresTaskRepository(db, logger)

	mock.ExpectQuery("SELECT id, title, description, created_at, updated_at FROM tasks WHERE id =").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "created_at", "updated_at"}).
			AddRow(1, "Test Task", "This is a test task", time.Now(), time.Now()))

	task, err := repo.GetTask(1)
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, int64(1), task.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresTaskRepository_UpdateTask(t *testing.T) {
	db, mock, logger := setupMockDB(t)
	defer db.Close()

	repo := NewPostgresTaskRepository(db, logger)
	task := &models.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mock.ExpectQuery("UPDATE tasks SET").
		WithArgs(task.Title, task.Description, sqlmock.AnyArg(), task.ID).
		WillReturnRows(sqlmock.NewRows([]string{"created_at"}).AddRow(time.Now()))

	updatedTask, err := repo.UpdateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, updatedTask)
	assert.Equal(t, task.ID, updatedTask.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPostgresTaskRepository_DeleteTask(t *testing.T) {
	db, mock, logger := setupMockDB(t)
	defer db.Close()

	repo := NewPostgresTaskRepository(db, logger)

	mock.ExpectExec("DELETE FROM tasks WHERE id =").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteTask(1)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
