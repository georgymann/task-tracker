package main

type Storage interface {
	CreateTask(task Task) (Task, error)
	GetTask(id int) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
}