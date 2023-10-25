package repositories

import (
	"hugeman/internal/core/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	layoutDateTimeRFC3339 = "2006-01-02T15:04:05Z"
)

// Postgres struct
type Postgres struct {
	dbGorm *gorm.DB
}

// NewPostgres func
func NewPostgres(dbGorm *gorm.DB) *Postgres {
	logrus.Info("Migrate database ...", layoutDateTimeRFC3339)
	domain.MigrateDatabase(dbGorm)
	return &Postgres{
		dbGorm: dbGorm,
	}
}

// CreateTodo func
func (p *Postgres) CreateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return nil, nil
}

// UpdateTodo func
func (p *Postgres) UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return nil, nil
}

// DeleteTodo func
func (p *Postgres) DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	return nil, nil
}

// GetTodo func
func (p *Postgres) GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error) {
	return nil, nil
}
