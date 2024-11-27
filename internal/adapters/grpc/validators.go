package grpc

import (
	"strings"

	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	MaxLength = 255 // Maximum length for string fields like Title.
	MinId     = 1   // Minimum valid ID value.
)

// trimAndValidateCreateTaskRequest validates and trims a CreateTaskRequest.
// Ensures Title and Description are not empty and Title does not exceed MaxLength.
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

// trimAndValidateUpdateTaskRequest validates and trims an UpdateTaskRequest.
// Ensures ID is valid, Title and Description are not empty, and Title does not exceed MaxLength.
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

// trimAndValidateGetTaskRequest validates a GetTaskRequest.
// Ensures ID is valid.
func trimAndValidateGetTaskRequest(req *pb.GetTaskRequest) error {
	if req.Id < MinId {
		return status.Error(codes.InvalidArgument, "ID must be greater than 0")
	}
	return nil
}

// trimAndValidateDeleteTaskRequest validates a DeleteTaskRequest.
// Ensures ID is valid.
func trimAndValidateDeleteTaskRequest(req *pb.DeleteTaskRequest) error {
	if req.Id < MinId {
		return status.Error(codes.InvalidArgument, "ID must be greater than 0")
	}
	return nil
}
