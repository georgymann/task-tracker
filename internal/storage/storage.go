package storage

import (
	"context"

	"task-tracker/internal/dto"
)

type Storage interface {
	CreateTask(ctx context.Context, task dto.Task) (dto.Task, error)
	GetTask(ctx context.Context, id int) (dto.Task, error)
	UpdateTask(ctx context.Context, task dto.Task) error
	DeleteTask(ctx context.Context, id int) error
}
