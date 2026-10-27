package storage

import (
	"context"
	"fmt"
	"sort"
	"time"

	"task-tracker/internal/dto"
)

type MemoryStorage struct {
	tasks  map[int]dto.Task
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int]dto.Task),
		nextID: 1,
	}
}

func (s *MemoryStorage) CreateTask(ctx context.Context, task dto.Task) (dto.Task, error) {
	if err := ctx.Err(); err != nil {
		return dto.Task{}, err
	}

	task.ID = s.nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt
	s.tasks[task.ID] = task
	s.nextID++
	return task, nil
}

func (s *MemoryStorage) GetTask(ctx context.Context, id int) (dto.Task, error) {
	if err := ctx.Err(); err != nil {
		return dto.Task{}, err
	}

	task, ok := s.tasks[id]
	if !ok {
		return dto.Task{}, fmt.Errorf("задача с ID %d не найдена", id)
	}
	return task, nil
}

func (s *MemoryStorage) GetTasks(ctx context.Context) ([]dto.Task, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	tasks := make([]dto.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks, nil
}

func (s *MemoryStorage) UpdateTask(ctx context.Context, task dto.Task) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	oldTask, ok := s.tasks[task.ID]
	if !ok {
		return fmt.Errorf("задача с ID %d не найдена", task.ID)
	}
	task.CreatedAt = oldTask.CreatedAt
	task.UpdatedAt = time.Now()
	s.tasks[task.ID] = task
	return nil
}

func (s *MemoryStorage) DeleteTask(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	_, ok := s.tasks[id]
	if !ok {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	delete(s.tasks, id)
	return nil
}
