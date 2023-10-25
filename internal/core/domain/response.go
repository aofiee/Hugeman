package domain

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

var (
	// Success response
	Success = Status{Code: http.StatusOK, Message: []string{"Success"}}
	// BadRequest response
	BadRequest = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Not responding because of incorrect syntax"}}
	// Unauthorized response
	Unauthorized = Status{Code: http.StatusUnauthorized, Message: []string{"Sorry, We are not able to process your request. Please try again"}}
	// Forbidden response
	Forbidden = Status{Code: http.StatusForbidden, Message: []string{"Sorry, Permission denied"}}
	// InternalServerError response
	InternalServerError = Status{Code: http.StatusInternalServerError, Message: []string{"Internal Server Error"}}
	// ConFlict response
	ConFlict = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Data is conflict"}}
	// FieldsPermission response
	FieldsPermission = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Fields are not able to update"}}
)

// ResponseBody struct
type ResponseBody struct {
	Status Status      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`

	CurrentPage *int   `json:"current_page,omitempty"`
	PerPage     *int   `json:"per_page,omitempty"`
	TotalItem   *int64 `json:"total_item,omitempty"`
}

// Status struct
type Status struct {
	Code    int      `json:"code,omitempty"`
	Message []string `json:"message,omitempty"`
}

type (
	// TodoResponse struct
	TodoResponse struct {
		ID          *uuid.UUID  `json:"id,omitempty" mapstructure:"id"`
		Title       *string     `json:"title,omitempty" mapstructure:"title"`
		Description *string     `json:"description,omitempty" mapstructure:"description"`
		Date        *string     `json:"date,omitempty" mapstructure:"date"`
		Image       *string     `json:"image,omitempty" mapstructure:"image"`
		Status      *TodoStatus `json:"status,omitempty" mapstructure:"status"`
		CreatedAt   *time.Time  `json:"created_at,omitempty" mapstructure:"created_at"`
		UpdatedAt   *time.Time  `json:"updated_at,omitempty" mapstructure:"updated_at"`
		DeletedAt   *time.Time  `json:"deleted_at,omitempty" mapstructure:"deleted_at"`
	}

	// TodoListResponse struct
	TodoListResponse struct {
		Todos []TodoResponse `json:"todos,omitempty" mapstructure:"todos"`

		CurrentPage *int   `json:"current_page,omitempty" mapstructure:"current_page"`
		PerPage     *int   `json:"per_page,omitempty" mapstructure:"per_page"`
		TotalItem   *int64 `json:"total_item,omitempty" mapstructure:"total_item"`
	}
)
