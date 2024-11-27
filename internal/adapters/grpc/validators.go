package grpc

import (
	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	MaxLength = 255
	MinId     = 1
)

func trimAndValidateCreateTaskRequest(req *pb.CreateTaskRequest) error {
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	if req.Title == "" {
		return status.Error(codes.InvalidArgument, "Title cannot be empty")
	}
	if len(req.Title) > MaxLength {
		return status.Error(codes.InvalidArgument, "Title exceeds maximum length of 255 characters")
	}
	if req.Description == "" {
		return status.Error(codes.InvalidArgument, "Description cannot be empty")
	}
	return nil
}

func trimAndValidateUpdateTaskRequest(req *pb.UpdateTaskRequest) error {
	req.Title = strings.TrimSpace(req.Title)
	req.Description = strings.TrimSpace(req.Description)

	if req.Id < MinId {
		return status.Error(codes.InvalidArgument, "ID must be greater than 0")
	}
	if req.Title == "" {
		return status.Error(codes.InvalidArgument, "Title cannot be empty")
	}
	if len(req.Title) > MaxLength {
		return status.Error(codes.InvalidArgument, "Title exceeds maximum length of 255 characters")
	}
	if req.Description == "" {
		return status.Error(codes.InvalidArgument, "Description cannot be empty")
	}
	return nil
}

func trimAndValidateGetTaskRequest(req *pb.GetTaskRequest) error {
	if req.Id < MinId {
		return status.Error(codes.InvalidArgument, "ID must be greater than 0")
	}
	return nil
}

func trimAndValidateDeleteTaskRequest(req *pb.DeleteTaskRequest) error {
	if req.Id < MinId {
		return status.Error(codes.InvalidArgument, "ID must be greater than 0")
	}
	return nil
}
