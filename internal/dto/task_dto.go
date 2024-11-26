package dto

import "time"

type TaskCreateRequestDTO struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
}

type TaskCreateResponseDTO struct {
	ID int64 `json:"id"`
}

type TaskUpdateRequestDTO struct {
	ID          int64  `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
}

type TaskUpdateResponseDTO struct {
	Success bool `json:"success"`
}

type TaskGetResponseDTO struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskListResponseDTO struct {
	Tasks []TaskGetResponseDTO `json:"tasks"`
}

type TaskDeleteRequestDTO struct {
	ID int64 `json:"id" validate:"required"`
}

type TaskDeleteResponseDTO struct {
	Success bool `json:"success"`
}
