package repositories

import (
	"encoding/base64"
	"hugeman/internal/core/domain"
	"time"

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
	var (
		err           error
		response      domain.TodoResponse
		encodingImage string
	)
	if request.Image != nil {
		encodingImage = base64.StdEncoding.EncodeToString([]byte(*request.Image))
	}
	todo := domain.Todo{
		Title:       request.Title,
		Description: request.Description,
		Image:       &encodingImage,
		Status:      request.Status,
	}
	if request.Date != nil {
		_date, err := time.Parse(layoutDateTimeRFC3339, *request.Date)
		if err != nil {
			logrus.Errorln(err)
			return &response, err
		}
		todo.Date = &_date
	}
	if err = p.dbGorm.Create(&todo).Error; err != nil {
		logrus.Errorln(err)
		return &response, err
	}
	df := todo.Date.Format(layoutDateTimeRFC3339)
	response = domain.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Date:        &df,
		Image:       todo.Image,
		Status:      todo.Status,
	}
	return &response, nil
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
