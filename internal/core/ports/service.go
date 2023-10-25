package ports

import "hugeman/internal/core/domain"

// Service interface
type Service interface {
	CreateTodo(request domain.TodoRequest) (*domain.TodoResponse, error)
	UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error)
	DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error)
	GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error)
}
