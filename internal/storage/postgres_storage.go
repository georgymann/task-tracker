package storage

import (
	"context"
	"fmt"

	"task-tracker/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
}

func NewPostgresStorage(pool *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{pool: pool}
}

func (s *PostgresStorage) CreateTask(ctx context.Context, task dto.Task) (dto.Task, error) {
	if task.Status == "" {
		task.Status = dto.StatusToDo
	}

	query := `INSERT INTO tasks (title, description, is_completed, priority, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`

	err := s.pool.QueryRow(ctx, query, task.Title, task.Description,
		task.IsCompleted, task.Priority, task.Status).Scan(
		&task.ID, &task.CreatedAt, &task.UpdatedAt)
	return task, err
}

func (s *PostgresStorage) GetTask(ctx context.Context, id int) (dto.Task, error) {
	var task dto.Task
	query := `SELECT id, title, description, is_completed, created_at, updated_at, priority, status
		FROM tasks WHERE id = $1`

	err := s.pool.QueryRow(ctx, query, id).Scan(
		&task.ID, &task.Title, &task.Description, &task.IsCompleted,
		&task.CreatedAt, &task.UpdatedAt, &task.Priority, &task.Status)
	return task, err
}

func (s *PostgresStorage) UpdateTask(ctx context.Context, task dto.Task) error {
	query := `UPDATE tasks SET title = $1, description = $2, is_completed = $3,
		priority = $4, status = $5, updated_at = now() WHERE id = $6`

	result, err := s.pool.Exec(ctx, query, task.Title, task.Description,
		task.IsCompleted, task.Priority, task.Status, task.ID)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("задача с ID %d не найдена", task.ID)
	}
	return nil
}

func (s *PostgresStorage) DeleteTask(ctx context.Context, id int) error {
	result, err := s.pool.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return nil
}
