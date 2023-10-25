package repositories

import (
	"encoding/base64"
	"errors"
	"hugeman/internal/core/domain"
	"net/url"
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
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		DeletedAt:   todo.DeletedAt,
	}
	return &response, nil
}

// UpdateTodo func
func (p *Postgres) UpdateTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	var (
		todo          domain.Todo
		response      domain.TodoResponse
		decodingImage string
		df            string
	)
	payload := domain.QueryTodoRequest{
		ID: request.ID,
	}
	condition := p.condition(payload)
	columns := p.updateColumns(request)
	if len(columns) == 0 {
		return &response, errors.New("fields are not able to update")
	}
	tx := p.dbGorm.Begin()
	defer func() {
		tx.Rollback()
	}()
	tx.Table(todo.TableName()).Where(condition).Updates(columns)
	if tx.Error != nil {
		logrus.Errorln(tx.Error)
		return &response, tx.Error
	}
	tx.Where(condition).First(&todo)
	if tx.Error != nil {
		logrus.Errorln(tx.Error)
		return &response, tx.Error
	}
	tx.Commit()
	if todo.ID == nil {
		return &response, errors.New("data not found")
	}
	if todo.Image != nil {
		dec, err := base64.StdEncoding.DecodeString(*todo.Image)
		if err != nil {
			logrus.Errorln(err)
			return &response, err
		}
		decodingImage = string(dec)
	}
	if todo.Date != nil {
		df = todo.Date.Format(layoutDateTimeRFC3339)
	}
	response = domain.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Date:        &df,
		Image:       &decodingImage,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		DeletedAt:   todo.DeletedAt,
	}
	return &response, nil
}

func (p *Postgres) condition(condition domain.QueryTodoRequest) map[string]interface{} {
	expression := make(map[string]interface{})
	if condition.ID != nil {
		expression["id"] = *condition.ID
	}
	if condition.Status != nil {
		expression["status"] = *condition.Status
	}
	return expression
}

func (p *Postgres) updateColumns(request domain.TodoRequest) map[string]interface{} {
	expression := make(map[string]interface{})
	if request.Title != nil {
		expression["title"] = *request.Title
	}
	if request.Description != nil {
		expression["description"] = *request.Description
	}
	if request.Date != nil {
		_date, err := time.Parse(layoutDateTimeRFC3339, *request.Date)
		if err != nil {
			logrus.Errorln(err)
		}
		expression["date"] = _date
	}
	if request.Image != nil {
		encodingImage := base64.StdEncoding.EncodeToString([]byte(*request.Image))
		expression["image"] = encodingImage
	}
	if request.Status != nil {
		expression["status"] = *request.Status
	}
	return expression
}

// DeleteTodo func
func (p *Postgres) DeleteTodo(request domain.TodoRequest) (*domain.TodoResponse, error) {
	var (
		todo          domain.Todo
		response      domain.TodoResponse
		decodingImage string
		df            string
	)
	payload := domain.QueryTodoRequest{
		ID: request.ID,
	}
	condition := p.condition(payload)
	if err := p.dbGorm.Table(todo.TableName()).Where(condition).First(&todo).Error; err != nil {
		logrus.Errorln(err)
		return &response, err
	}
	tx := p.dbGorm.Begin()
	defer func() {
		tx.Rollback()
	}()
	tx.Delete(&todo)
	if tx.Error != nil {
		logrus.Errorln(tx.Error)
		return &response, tx.Error
	}
	tx.Commit()
	if todo.Image != nil {
		dec, err := base64.StdEncoding.DecodeString(*todo.Image)
		if err != nil {
			logrus.Errorln(err)
			return &response, err
		}
		decodingImage = string(dec)
	}
	if todo.Date != nil {
		df = todo.Date.Format(layoutDateTimeRFC3339)
	}
	response = domain.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Date:        &df,
		Image:       &decodingImage,
		Status:      todo.Status,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		DeletedAt:   todo.DeletedAt,
	}
	return &response, nil
}

// GetTodo func
func (p *Postgres) GetTodo(condition domain.QueryTodoRequest) (*domain.TodoListResponse, error) {
	var (
		todo  domain.Todo
		todos []domain.Todo
	)
	cond := p.condition(condition)
	tx := p.dbGorm.Where(cond)

	if condition.Title != nil {
		keyword, err := url.QueryUnescape(*condition.Title)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		tx = tx.Where("title ILIKE ? ", "%"+keyword+"%")
	}
	if condition.Description != nil {
		keyword, err := url.QueryUnescape(*condition.Description)
		if err != nil {
			logrus.Errorln(err)
			return nil, err
		}
		tx = tx.Where("description ILIKE ? ", "%"+keyword+"%")
	}

	var toatalItem int64
	tx.Model(&todo).Count(&toatalItem)

	if condition.ID == nil {
		var order string
		if condition.SortMethod.OrderBy != "" {
			order = condition.SortMethod.OrderBy
		} else {
			order = "id"
		}
		if condition.SortMethod.Asc {
			tx = tx.Order(order + " ASC")
		} else {
			tx = tx.Order(order + " DESC")
		}
		logrus.Info("order by ", condition.SortMethod.Asc)
		tx = tx.Limit(condition.Pagination.Limit).Offset(condition.Pagination.Offset)
	}

	tx.Find(&todos)
	if tx.Error != nil {
		logrus.Errorln(tx.Error)
		return nil, tx.Error
	}
	result := domain.TodoListResponse{
		Todos: []domain.TodoResponse{},
	}

	result.CurrentPage = condition.Page
	result.PerPage = &condition.Pagination.Limit
	result.TotalItem = &toatalItem
	for _, todo := range todos {
		var (
			decodingImage string
			df            string
		)
		todo := todo
		if todo.Image != nil {
			dec, err := base64.StdEncoding.DecodeString(*todo.Image)
			if err != nil {
				logrus.Errorln(err)
				return nil, err
			}
			decodingImage = string(dec)
		}
		if todo.Date != nil {
			df = todo.Date.Format(layoutDateTimeRFC3339)
		}
		data := domain.TodoResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Date:        &df,
			Image:       &decodingImage,
			Status:      todo.Status,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
			DeletedAt:   todo.DeletedAt,
		}
		result.Todos = append(result.Todos, data)
	}
	return &result, nil
}
