package grpc

import (
	"context"
	"github.com/Sunf1ower113/grpc-task-manager/internal/domain/services"
	"github.com/Sunf1ower113/grpc-task-manager/proto"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestErrTaskNotFound(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrTaskNotFound(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("ErrTaskNotFound() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewTaskHandler(t *testing.T) {
	type args struct {
		service services.TaskService
		logger  *zap.Logger
	}
	tests := []struct {
		name string
		args args
		want proto.TaskManagerServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskHandler(tt.args.service, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_CreateTask(t *testing.T) {
	type fields struct {
		UnimplementedTaskManagerServer proto.UnimplementedTaskManagerServer
		service                        services.TaskService
		logger                         *zap.Logger
	}
	type args struct {
		ctx context.Context
		req *pb.CreateTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.TaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TaskHandler{
				UnimplementedTaskManagerServer: tt.fields.UnimplementedTaskManagerServer,
				service:                        tt.fields.service,
				logger:                         tt.fields.logger,
			}
			got, err := h.CreateTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	type fields struct {
		UnimplementedTaskManagerServer proto.UnimplementedTaskManagerServer
		service                        services.TaskService
		logger                         *zap.Logger
	}
	type args struct {
		ctx context.Context
		req *pb.DeleteTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteTaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TaskHandler{
				UnimplementedTaskManagerServer: tt.fields.UnimplementedTaskManagerServer,
				service:                        tt.fields.service,
				logger:                         tt.fields.logger,
			}
			got, err := h.DeleteTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_GetTask(t *testing.T) {
	type fields struct {
		UnimplementedTaskManagerServer proto.UnimplementedTaskManagerServer
		service                        services.TaskService
		logger                         *zap.Logger
	}
	type args struct {
		ctx context.Context
		req *pb.GetTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.TaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TaskHandler{
				UnimplementedTaskManagerServer: tt.fields.UnimplementedTaskManagerServer,
				service:                        tt.fields.service,
				logger:                         tt.fields.logger,
			}
			got, err := h.GetTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_ListTasks(t *testing.T) {
	type fields struct {
		UnimplementedTaskManagerServer proto.UnimplementedTaskManagerServer
		service                        services.TaskService
		logger                         *zap.Logger
	}
	type args struct {
		ctx context.Context
		req *pb.ListTasksRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ListTasksResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TaskHandler{
				UnimplementedTaskManagerServer: tt.fields.UnimplementedTaskManagerServer,
				service:                        tt.fields.service,
				logger:                         tt.fields.logger,
			}
			got, err := h.ListTasks(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	type fields struct {
		UnimplementedTaskManagerServer proto.UnimplementedTaskManagerServer
		service                        services.TaskService
		logger                         *zap.Logger
	}
	type args struct {
		ctx context.Context
		req *pb.UpdateTaskRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.TaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TaskHandler{
				UnimplementedTaskManagerServer: tt.fields.UnimplementedTaskManagerServer,
				service:                        tt.fields.service,
				logger:                         tt.fields.logger,
			}
			got, err := h.UpdateTask(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}
