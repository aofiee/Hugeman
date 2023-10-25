package handlers

import (
	"hugeman/internal/core/domain"
	"hugeman/internal/core/ports"
	"hugeman/pkg/validator"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
/* create todo */
// CreateTodo godoc
// @Summary Create todo
// @Description Create todo
// @Tags TODO
// @Accept application/json
// @Success 200 {object} map[string]interface{}
// @Router /v1/api/todo	[post]
// @Produce json
// @param CreateTodo body domain.TodoRequest true "CreateTodo"
func (hdl *HTTPHandler) CreateTodo(c *fiber.Ctx) error {
	var request domain.TodoRequest
	if err := c.BodyParser(&request); err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}
	if err := hdl.validator.ValidateStruct(request); err != nil {
		msg := domain.ResponseBody{
			Status: domain.BadRequest,
		}
		msg.Status.Message = []string{
			err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(msg)
	}
	response, err := hdl.srv.CreateTodo(request)
	if err != nil {
		logrus.Errorln(err)
		msg := domain.ResponseBody{
			Status: domain.InternalServerError,
		}
		msg.Status.Message = []string{
			err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(msg)
	}
	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{Status: domain.Success, Data: response})
}

// UpdateTodo func
/* update todo */
// UpdateTodo godoc
// @Summary Update todo
// @Description Update todo
// @Tags TODO
// @Accept application/json
// @Success 200 {object} map[string]interface{}
// @Router /v1/api/todo	[put]
// @Produce json
// @param UpdateTodo body domain.TodoRequest true "UpdateTodo"
func (hdl *HTTPHandler) UpdateTodo(c *fiber.Ctx) error {
	var request domain.TodoRequest
	if err := c.BodyParser(&request); err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}
	if err := hdl.validator.ValidateStruct(request); err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}
	if request.ID == nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}
	response, err := hdl.srv.UpdateTodo(request)
	if err != nil {
		msg := domain.ResponseBody{
			Status: domain.InternalServerError,
		}
		msg.Status.Message = []string{
			err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(msg)
	}
	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{Status: domain.Success, Data: response})
}

// DeleteTodo func
/* delete todo */
// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete todo
// @Tags TODO
// @Accept application/json
// @Success 200 {object} map[string]interface{}
// @Router /v1/api/todo/{id}	[delete]
// @Produce json
// @param id path string true "uuid"
func (hdl *HTTPHandler) DeleteTodo(c *fiber.Ctx) error {
	var (
		uid uuid.UUID
		err error
	)
	id := c.Params("id")
	if id != "" {
		uid, err = uuid.Parse(id)
		if err != nil {
			logrus.Errorln(err)
			return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
		}
	}

	var request domain.TodoRequest
	request.ID = &uid
	response, err := hdl.srv.DeleteTodo(request)
	if err != nil {
		msg := domain.ResponseBody{
			Status: domain.InternalServerError,
		}
		msg.Status.Message = []string{
			err.Error(),
		}
		return c.Status(fiber.StatusInternalServerError).JSON(msg)
	}
	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{Status: domain.Success, Data: response})
}

// GetTodo func
/* delete todo */
// DeleteTodo godoc
// @Summary Delete todo
// @Description Delete todo
// @Tags TODO
// @Accept application/json
// @Success 200 {object} map[string]interface{}
// @Router /v1/api/todo	[get]
// @Produce json
// @param id query string false "uuid"
// @param page query int false "page"
// @param limit query int false "limit"
// @param orderBy query string false "order_by"
// @param asc query bool false "asc"
// @param title query string false "title"
// @param description query string false "description"
// @param status query string false "status"
func (hdl *HTTPHandler) GetTodo(c *fiber.Ctx) error {
	var (
		uid  uuid.UUID
		err  error
		data []domain.TodoResponse
	)
	condition := domain.QueryTodoRequest{}
	err = c.QueryParser(&condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	err = hdl.validator.ValidateStruct(condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
	}

	id := c.Params("id")
	if id != "" {
		uid, err = uuid.Parse(id)
		if err != nil {
			logrus.Errorln(err)
			return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseBody{Status: domain.BadRequest})
		}
		condition.ID = &uid
	}
	result, err := hdl.srv.GetTodo(condition)
	if err != nil {
		logrus.Errorln(err)
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseBody{Status: domain.InternalServerError})
	}
	if result.Todos == nil {
		data = make([]domain.TodoResponse, 0)
	} else {
		data = result.Todos
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseBody{
		Status:      domain.Success,
		Data:        data,
		CurrentPage: result.CurrentPage,
		PerPage:     result.PerPage,
		TotalItem:   result.TotalItem,
	})
}
