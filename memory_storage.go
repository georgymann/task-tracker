package main

import (
	"fmt"
)

type MemoryStorage struct {
	tasks map[int]Task
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
	tasks: make(map[int]Task),
	nextID: 1,
	}
}

func (s *MemoryStorage) CreateTask(task Task) (Task, error) {
	task.ID = s.nextID
	s.tasks[task.ID] = task
	s.nextID++
	return task, nil
}

func (s *MemoryStorage) GetTask(id int) (Task, error) {
	task, ok := s.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("задача с ID %d не найдена", id)
	}
	return task, nil
}

func (s *MemoryStorage) UpdateTask(task Task) error {
	_, ok := s.tasks[task.ID]
	if !ok {
		return fmt.Errorf("задача %+v с ID %d не найдена", task, task.ID)
	}
	s.tasks[task.ID] = task
	return nil
}

func (s *MemoryStorage) DeleteTask(id int) error {
	_, ok := s.tasks[id]
	if !ok {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	newTasks := make(map[int]Task)
	for key, task := range s.tasks {
		if key != id {
			newTasks[key] = task
		}
	}
	s.tasks = newTasks
	return nil
}