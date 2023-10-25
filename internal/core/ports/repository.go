package ports

import (
	"hugeman/internal/core/domain"
)

// Repository interface
type Repository interface {
	CreateTodo(request domain.TodoRequest) (*domain.TodoResponse, error)
	UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error)
	DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error)
	GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error)
}
