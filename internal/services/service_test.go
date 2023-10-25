package services

import (
	"errors"
	"testing"

	"hugeman/internal/core/domain"
	mockrepository "hugeman/pkg/testings/mock_repository"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	ctrl *gomock.Controller
	repo *mockrepository.MockRepository
	svc  *Service
)

// SetupTest func
func SetupTest(t *testing.T) {
	ctrl = gomock.NewController(t)
	repo = mockrepository.NewMockRepository(ctrl)
	svc = New(repo)
}

// TestCreateTodo func
func TestCreateTodo(t *testing.T) {
	SetupTest(t)

	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	title := "title"
	description := "description"
	image := "image"
	status := domain.TodoStatus("COMPLETE")
	date := "2021-01-01T00:00:00Z"

	todo := domain.TodoRequest{
		Title:       &title,
		Description: &description,
		Image:       &image,
		Status:      &status,
		Date:        &date,
	}

	t.Run("success", func(t *testing.T) {
		repo.EXPECT().CreateTodo(todo).Return(&domain.TodoResponse{
			ID:          &uuid,
			Title:       &title,
			Description: &description,
			Image:       &image,
			Status:      &status,
			Date:        &date,
		}, nil)

		rs, err := svc.CreateTodo(todo)
		assert.NoError(t, err)
		assert.NotNil(t, rs)
		assert.Equal(t, &uuid, rs.ID)
		assert.Equal(t, &title, rs.Title)
		assert.Equal(t, &description, rs.Description)
		assert.Equal(t, &status, rs.Status)
		assert.Equal(t, &date, rs.Date)
	})

	t.Run("failed", func(t *testing.T) {
		repo.EXPECT().CreateTodo(todo).Return(&domain.TodoResponse{}, errors.New("error"))

		rs, err := svc.CreateTodo(todo)
		assert.Error(t, err)
		assert.Nil(t, rs)
	})
}

// TestUpdateTodo func
func TestUpdateTodo(t *testing.T) {
	SetupTest(t)

	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	title := "title"
	description := "description"
	image := "image"
	status := domain.TodoStatus("COMPLETE")
	date := "2021-01-01T00:00:00Z"

	todo := domain.TodoRequest{
		Title:       &title,
		Description: &description,
		Image:       &image,
		Status:      &status,
		Date:        &date,
	}

	t.Run("success", func(t *testing.T) {
		repo.EXPECT().UpdateTodo(todo).Return(&domain.TodoResponse{
			ID:          &uuid,
			Title:       &title,
			Description: &description,
			Image:       &image,
			Status:      &status,
			Date:        &date,
		}, nil)

		rs, err := svc.UpdateTodo(todo)
		assert.NoError(t, err)
		assert.NotNil(t, rs)
		assert.Equal(t, &uuid, rs.ID)
		assert.Equal(t, &title, rs.Title)
		assert.Equal(t, &description, rs.Description)
		assert.Equal(t, &status, rs.Status)
		assert.Equal(t, &date, rs.Date)
	})
}

// TestDeleteTodo func
func TestDeletetTodo(t *testing.T) {
	SetupTest(t)

	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	title := "title"
	description := "description"
	image := "image"
	status := domain.TodoStatus("COMPLETE")
	date := "2021-01-01T00:00:00Z"

	todo := domain.TodoRequest{
		Title:       &title,
		Description: &description,
		Image:       &image,
		Status:      &status,
		Date:        &date,
	}

	t.Run("success", func(t *testing.T) {
		repo.EXPECT().DeleteTodo(todo).Return(&domain.TodoResponse{
			ID:          &uuid,
			Title:       &title,
			Description: &description,
			Image:       &image,
			Status:      &status,
			Date:        &date,
		}, nil)

		rs, err := svc.DeleteTodo(todo)
		assert.NoError(t, err)
		assert.NotNil(t, rs)
		assert.Equal(t, &uuid, rs.ID)
		assert.Equal(t, &title, rs.Title)
		assert.Equal(t, &description, rs.Description)
		assert.Equal(t, &status, rs.Status)
		assert.Equal(t, &date, rs.Date)
	})
}

// TestGetTodo func
func TestGetTodo(t *testing.T) {
	SetupTest(t)

	t.Run("success", func(t *testing.T) {
		condition := domain.QueryTodoRequest{}
		page := 1
		perPage := 100
		condition.Page = &page
		condition.Limit = &perPage
		offset := (page - 1) * perPage
		condition.Pagination = &domain.Pagination{
			Limit:  perPage,
			Offset: offset,
		}
		condition.SortMethod = &domain.SortMethod{
			OrderBy: "ID",
			Asc:     true,
		}
		repo.EXPECT().GetTodo(condition).Return(&domain.TodoListResponse{}, nil)

		rs, err := svc.GetTodo(condition)
		assert.NoError(t, err)
		assert.NotNil(t, rs)
	})

	t.Run("success query nil", func(t *testing.T) {
		repo.EXPECT().GetTodo(gomock.Any()).Return(&domain.TodoListResponse{}, nil)

		rs, err := svc.GetTodo(domain.QueryTodoRequest{})
		assert.NoError(t, err)
		assert.NotNil(t, rs)
	})
}
