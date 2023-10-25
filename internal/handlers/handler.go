package handlers

import (
	"hugeman/internal/core/domain"
	"hugeman/internal/core/ports"
	"hugeman/pkg/validator"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// HTTPHandler struct
type HTTPHandler struct {
	srv       ports.Service
	db        *gorm.DB
	validator validator.Validator
}

// New func
func New(srv ports.Service, db *gorm.DB) *HTTPHandler {
	return &HTTPHandler{
		srv:       srv,
		db:        db,
		validator: validator.New(),
	}
}

// HealthCheck func
func (hdl *HTTPHandler) HealthCheck(c *fiber.Ctx) error {
	sqlDB, err := hdl.db.DB()
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseBody{Status: domain.InternalServerError})
	}

	err = sqlDB.Ping()
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseBody{Status: domain.InternalServerError})
	}
	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{Status: domain.Success, Data: ""})
}

// CreateTodo func
func (hdl *HTTPHandler) CreateTodo(c *fiber.Ctx) error {
	return nil
}

// UpdateTodo func
func (hdl *HTTPHandler) UpdateTodo(c *fiber.Ctx) error {
	return nil
}

// DeleteTodo func
func (hdl *HTTPHandler) DeleteTodo(c *fiber.Ctx) error {
	return nil
}

// GetTodo func
func (hdl *HTTPHandler) GetTodo(c *fiber.Ctx) error {
	return nil
}
