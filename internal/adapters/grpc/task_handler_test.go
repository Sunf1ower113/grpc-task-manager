package grpc

import (
	"context"
	"testing"

	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/models"
	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) CreateTask(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockService) ListTasks() ([]*models.Task, error) {
	args := m.Called()
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockService) GetTask(id int64) (*models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockService) UpdateTask(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockService) DeleteTask(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupHandler() (*MockService, *TaskHandler) {
	mockService := new(MockService)
	logger, _ := zap.NewDevelopment()
	handler := &TaskHandler{
		service: mockService,
		logger:  logger,
	}
	return mockService, handler
}

func TestTaskHandler_CreateTask(t *testing.T) {
	mockService, handler := setupHandler()

	mockTask := &models.Task{
		Title:       "Test Task",
		Description: "Description for test task",
	}

	mockResponse := &models.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Description for test task",
	}

	mockService.On("CreateTask", mockTask).Return(mockResponse, nil)

	req := &pb.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Description for test task",
	}

	resp, err := handler.CreateTask(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, int64(1), resp.Id)
	require.Equal(t, "Test Task", resp.Title)

	mockService.AssertExpectations(t)
}

func TestTaskHandler_ListTasks(t *testing.T) {
	mockService, handler := setupHandler()

	mockTasks := []*models.Task{
		{
			ID:          1,
			Title:       "Task 1",
			Description: "Description 1",
		},
		{
			ID:          2,
			Title:       "Task 2",
			Description: "Description 2",
		},
	}

	mockService.On("ListTasks").Return(mockTasks, nil)

	req := &pb.ListTasksRequest{}

	resp, err := handler.ListTasks(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Len(t, resp.Tasks, 2)
	require.Equal(t, int64(1), resp.Tasks[0].Id)
	require.Equal(t, "Task 1", resp.Tasks[0].Title)

	mockService.AssertExpectations(t)
}

func TestTaskHandler_GetTask(t *testing.T) {
	mockService, handler := setupHandler()

	mockTask := &models.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
	}

	mockService.On("GetTask", int64(1)).Return(mockTask, nil)

	req := &pb.GetTaskRequest{Id: 1}

	resp, err := handler.GetTask(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, int64(1), resp.Id)
	require.Equal(t, "Test Task", resp.Title)

	mockService.AssertExpectations(t)
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	mockService, handler := setupHandler()

	mockTask := &models.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockService.On("UpdateTask", mockTask).Return(mockTask, nil)

	req := &pb.UpdateTaskRequest{
		Id:          1,
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	resp, err := handler.UpdateTask(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, int64(1), resp.Id)
	require.Equal(t, "Updated Task", resp.Title)

	mockService.AssertExpectations(t)
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	mockService, handler := setupHandler()

	mockService.On("DeleteTask", int64(1)).Return(nil)

	req := &pb.DeleteTaskRequest{Id: 1}

	resp, err := handler.DeleteTask(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.True(t, resp.Success)

	mockService.AssertExpectations(t)
}
