package dto

import "time"

const (
	StatusToDo Status = "todo"
	StatusWip  Status = "in_progress"
	StatusDone Status = "done"
)

type Task struct {
	ID          int
	Title       string
	Description string
	IsCompleted bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Priority    int
	Status      Status
}

type Status = string
