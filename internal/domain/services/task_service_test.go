package services

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	"go.uber.org/zap/zaptest"
)

type mockTaskRepository struct {
	tasks map[int64]*models.Task
	err   error
}

func (m *mockTaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	task.ID = int64(len(m.tasks) + 1)
	m.tasks[task.ID] = task
	return task, nil
}

func (m *mockTaskRepository) ListTasks() ([]*models.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	var taskList []*models.Task
	for _, task := range m.tasks {
		taskList = append(taskList, task)
	}
	return taskList, nil
}

func (m *mockTaskRepository) GetTask(id int64) (*models.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	task, exists := m.tasks[id]
	if !exists {
		return nil, nil
	}
	return task, nil
}

func (m *mockTaskRepository) UpdateTask(task *models.Task) (*models.Task, error) {
	if m.err != nil {
		return nil, m.err
	}
	_, exists := m.tasks[task.ID]
	if !exists {
		return nil, sql.ErrNoRows
	}
	m.tasks[task.ID] = task
	return task, nil
}

func (m *mockTaskRepository) DeleteTask(id int64) error {
	if m.err != nil {
		return m.err
	}
	_, exists := m.tasks[id]
	if !exists {
		return sql.ErrNoRows
	}
	delete(m.tasks, id)
	return nil
}

func Test_taskService_CreateTask(t *testing.T) {
	logger := zaptest.NewLogger(t)
	mockRepo := &mockTaskRepository{tasks: make(map[int64]*models.Task)}

	svc := NewTaskService(mockRepo, logger)

	tests := []struct {
		name    string
		task    *models.Task
		wantErr bool
	}{
		{
			name:    "Valid task creation",
			task:    &models.Task{Title: "Test Task", Description: "Test Description"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.CreateTask(tt.task)
			log.Println(err)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_taskService_ListTasks(t *testing.T) {
	logger := zaptest.NewLogger(t)
	mockRepo := &mockTaskRepository{
		tasks: map[int64]*models.Task{
			1: {ID: 1, Title: "Task 1", Description: "Description 1"},
			2: {ID: 2, Title: "Task 2", Description: "Description 2"},
		},
	}

	svc := NewTaskService(mockRepo, logger)

	tasks, err := svc.ListTasks()
	if err != nil {
		t.Errorf("ListTasks() unexpected error: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("ListTasks() got %d tasks, want 2", len(tasks))
	}
}

func Test_taskService_GetTask(t *testing.T) {
	logger := zaptest.NewLogger(t)
	mockRepo := &mockTaskRepository{
		tasks: map[int64]*models.Task{
			1: {ID: 1, Title: "Task 1", Description: "Description 1"},
		},
	}

	svc := NewTaskService(mockRepo, logger)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Existing task",
			id:      1,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, err := svc.GetTask(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
			}
			if task != nil && task.ID != tt.id {
				t.Errorf("GetTask() got = %v, want %v", task.ID, tt.id)
			}
		})
	}
}

func Test_taskService_UpdateTask(t *testing.T) {
	logger := zaptest.NewLogger(t)
	mockRepo := &mockTaskRepository{
		tasks: map[int64]*models.Task{
			1: {ID: 1, Title: "Task 1", Description: "Description 1"},
		},
	}

	svc := NewTaskService(mockRepo, logger)

	tests := []struct {
		name    string
		task    *models.Task
		wantErr bool
	}{
		{
			name:    "Update existing task",
			task:    &models.Task{ID: 1, Title: "Updated Task", Description: "Updated Description"},
			wantErr: false,
		},
		{
			name:    "Update non-existing task",
			task:    &models.Task{ID: 2, Title: "New Task", Description: "New Description"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.UpdateTask(tt.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_taskService_DeleteTask(t *testing.T) {
	logger := zaptest.NewLogger(t)
	mockRepo := &mockTaskRepository{
		tasks: map[int64]*models.Task{
			1: {ID: 1, Title: "Task 1", Description: "Description 1"},
		},
	}

	svc := NewTaskService(mockRepo, logger)

	tests := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "Delete existing task",
			id:      1,
			wantErr: false,
		},
		{
			name:    "Delete non-existing task",
			id:      2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.DeleteTask(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
