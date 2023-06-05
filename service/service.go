package service

import (
	"ToDoLIst/repository"
	"github.com/zhashkevych/todo-app"
)

type Authorization interface {
	CreatUser(user todo.User) (int, error)
}
type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}

}
