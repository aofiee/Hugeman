package services

import (
	"hugeman/internal/core/domain"
	"hugeman/internal/core/ports"
)

// Service struct
type Service struct {
	repo ports.Repository
}

// New func
func New(repo ports.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateTodo func
func (s *Service) CreateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return nil, nil
}

// UpdateTodo func
func (s *Service) UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return nil, nil
}

// DeleteTodo func
func (s *Service) DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return nil, nil
}

// GetTodo func
func (s *Service) GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error) {
	return nil, nil
}
