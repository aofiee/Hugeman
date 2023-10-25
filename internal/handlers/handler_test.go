package handlers

import (
	"database/sql"
	"errors"
	"hugeman/internal/core/domain"
	mockservice "hugeman/pkg/testings/mock_service"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Suite struct
type Suite struct {
	dbGorm *gorm.DB
	mock   sqlmock.Sqlmock
}

// MockDB func
func MockDB() *Suite {
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	if err != nil {
		logrus.Error(err)
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.dbGorm, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	// defer db.Close()
	return s
}

// TestCreateTodoFailed func
func TestCreateTodoFailed(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("failed status 400 json broken body", func(t *testing.T) {
		srv.EXPECT().CreateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Post("/v1/api/todo", h.CreateTodo)
		req := httptest.NewRequest(fiber.MethodPost, "/v1/api/todo",
			strings.NewReader(`{
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE",
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	})

	t.Run("failed status 400", func(t *testing.T) {
		srv.EXPECT().CreateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Post("/v1/api/todo", h.CreateTodo)
		req := httptest.NewRequest(fiber.MethodPost, "/v1/api/todo",
			strings.NewReader(`{
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	})

	t.Run("failed status 500", func(t *testing.T) {
		srv.EXPECT().CreateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Post("/v1/api/todo", h.CreateTodo)
		req := httptest.NewRequest(fiber.MethodPost, "/v1/api/todo",
			strings.NewReader(`{
				"title": "world",
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

	})

}

// TestCreateTodoSuccess func
func TestCreateTodoSuccess(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("success status 200", func(t *testing.T) {
		srv.EXPECT().CreateTodo(gomock.Any()).Return(&domain.TodoResponse{}, nil)
		app := fiber.New()
		app.Post("/v1/api/todo", h.CreateTodo)
		req := httptest.NewRequest(fiber.MethodPost, "/v1/api/todo",
			strings.NewReader(`{
				"title": "world",
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	})
}

// TestUpdateTodoFailed func
func TestUpdateTodoFailed(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("failed status 400 json broken body", func(t *testing.T) {
		srv.EXPECT().UpdateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Put("/v1/api/todo", h.UpdateTodo)
		req := httptest.NewRequest(fiber.MethodPut, "/v1/api/todo",
			strings.NewReader(`{
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE",
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed status 400", func(t *testing.T) {
		srv.EXPECT().UpdateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Put("/v1/api/todo", h.UpdateTodo)
		req := httptest.NewRequest(fiber.MethodPut, "/v1/api/todo",
			strings.NewReader(`{
				"id": "539b68f0-cf3c-4ac8-8988-8e8de0ff9168",
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed status 400", func(t *testing.T) {
		srv.EXPECT().UpdateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Put("/v1/api/todo", h.UpdateTodo)
		req := httptest.NewRequest(fiber.MethodPut, "/v1/api/todo",
			strings.NewReader(`{
				"title": "world",
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed status 400", func(t *testing.T) {
		srv.EXPECT().UpdateTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Put("/v1/api/todo", h.UpdateTodo)
		req := httptest.NewRequest(fiber.MethodPut, "/v1/api/todo",
			strings.NewReader(`{
				"id": "539b68f0-cf3c-4ac8-8988-8e8de0ff9168",
				"title": "world",
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}

// TestUpdateTodoSuccess func
func TestUpdateTodoSuccess(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("success status 200", func(t *testing.T) {
		srv.EXPECT().UpdateTodo(gomock.Any()).Return(&domain.TodoResponse{}, nil)
		app := fiber.New()
		app.Put("/v1/api/todo", h.UpdateTodo)
		req := httptest.NewRequest(fiber.MethodPut, "/v1/api/todo",
			strings.NewReader(`{
				"id": "539b68f0-cf3c-4ac8-8988-8e8de0ff9168",
				"title": "world",
				"description": "golang",
				"date": "2023-10-25T19:46:05Z",
				"image": "https://scontent.fbkk8-3.fna.fbcdn.net/v/t39.30808-6/393759110_798549782070028_522951429028478775_n.jpg?_nc_cat=1&ccb=1-7&_nc_sid=5f2048&_nc_eui2=AeFvsI6YA_c4FeHdRi0zW2W48XPmwMFG_Uzxc-bAwUb9TFcrLmmmIIpAKPybpTBsiRJIZYoAmbkGGhw6ZdiDCDXb&_nc_ohc=CHQsmIugG-0AX8fQ--x&_nc_ht=scontent.fbkk8-3.fna&oh=00_AfBgFL0IgF2CLk-U3XUmqJ27Kswkfd1xWbilcRhBrjl9PQ&oe=653DF20F",
				"status": "COMPLETE"
			}`))
		req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

// TestDeleteTodoFailed func
func TestDeleteTodoFailed(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("failed status 500", func(t *testing.T) {
		srv.EXPECT().DeleteTodo(gomock.Any()).Return(&domain.TodoResponse{}, errors.New("error"))
		app := fiber.New()
		app.Delete("/v1/api/todo/:id", h.DeleteTodo)
		req := httptest.NewRequest(fiber.MethodDelete, "/v1/api/todo/539b68f0-cf3c-4ac8-8988-8e8de0ff9168", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("failed status 500", func(t *testing.T) {
		srv.EXPECT().DeleteTodo(gomock.Any()).Return(&domain.TodoResponse{}, errors.New("error"))
		app := fiber.New()
		app.Delete("/v1/api/todo/:id", h.DeleteTodo)
		req := httptest.NewRequest(fiber.MethodDelete, "/v1/api/todo/539b68f0-cf3c-4ac8-8988-8e8de0ff916", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

// TestDeleteTodoSuccess func
func TestDeleteTodoSuccess(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("success status 200", func(t *testing.T) {
		srv.EXPECT().DeleteTodo(gomock.Any()).Return(&domain.TodoResponse{}, nil)
		app := fiber.New()
		app.Delete("/v1/api/todo/:id", h.DeleteTodo)
		req := httptest.NewRequest(fiber.MethodDelete, "/v1/api/todo/539b68f0-cf3c-4ac8-8988-8e8de0ff9168", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

// TestGetTodoFailed func
func TestGetTodoFailed(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("failed status 500", func(t *testing.T) {
		srv.EXPECT().GetTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Get("/v1/api/todo", h.GetTodo)
		req := httptest.NewRequest(fiber.MethodGet, "/v1/api/todo", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("failed status 500", func(t *testing.T) {
		srv.EXPECT().GetTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Get("/v1/api/todo", h.GetTodo)
		req := httptest.NewRequest(fiber.MethodGet, "/v1/api/todo?id=xxx", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("failed status 500", func(t *testing.T) {
		srv.EXPECT().GetTodo(gomock.Any()).Return(nil, errors.New("error"))
		app := fiber.New()
		app.Get("/v1/api/todo/:id", h.GetTodo)
		req := httptest.NewRequest(fiber.MethodGet, "/v1/api/todo/xxx", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})
}

// TestGetTodoSuccess func
func TestGetTodoSuccess(t *testing.T) {
	s := MockDB()
	ctrl := gomock.NewController(t)
	srv := mockservice.NewMockService(ctrl)
	h := New(srv, s.dbGorm)

	t.Run("success status 200", func(t *testing.T) {
		srv.EXPECT().GetTodo(gomock.Any()).Return(&domain.TodoListResponse{
			Todos: []domain.TodoResponse{
				{
					ID: func() *uuid.UUID { s, _ := uuid.NewRandom(); return &s }(),
				},
			},
		}, nil)
		app := fiber.New()
		app.Get("/v1/api/todo", h.GetTodo)
		req := httptest.NewRequest(fiber.MethodGet, "/v1/api/todo", nil)
		resp, err := app.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
