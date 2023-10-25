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
	return s.repo.CreateTodo(request)
}

// UpdateTodo func
func (s *Service) UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return s.repo.UpdateTodo(request)
}

// DeleteTodo func
func (s *Service) DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return s.repo.DeleteTodo(request)
}

// GetTodo func
func (s *Service) GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error) {
	var (
		page    int
		perPage int
		offset  int
	)
	if condition.Page != nil {
		page = *condition.Page
	} else {
		page = 1
		condition.Page = &page
	}
	if condition.Limit != nil {
		perPage = *condition.Limit
	} else {
		perPage = 100
		condition.Limit = &perPage
	}
	offset = (page - 1) * perPage
	condition.Pagination = &domain.Pagination{
		Limit:  perPage,
		Offset: offset,
	}
	if condition.OrderBy != nil {
		asc := true
		if condition.Asc != nil {
			asc = *condition.Asc
		}
		condition.SortMethod = &domain.SortMethod{
			Asc:     asc,
			OrderBy: *condition.OrderBy,
		}
	} else {
		condition.SortMethod = &domain.SortMethod{
			Asc:     true,
			OrderBy: "ID",
		}
	}
	return s.repo.GetTodo(condition)
}
