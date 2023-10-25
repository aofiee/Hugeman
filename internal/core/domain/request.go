package domain

import "github.com/google/uuid"

type (
	// TodoRequest struct
	TodoRequest struct {
		ID          *uuid.UUID  `json:"id" validate:"omitempty,uuid4" form:"id" query:"id"`
		Title       *string     `json:"title" validate:"required,max=100" form:"title" query:"title"`
		Description *string     `json:"description" validate:"omitempty" form:"description" query:"description"`
		Date        *string     `json:"date" validate:"required" form:"date" query:"date"`
		Image       *string     `json:"image" validate:"omitempty" form:"image" query:"image"`
		Status      *TodoStatus `json:"status" validate:"required,oneof=IN_PROGRESS COMPLETE" form:"status" query:"status"`
	}

	// QueryTodoRequest struct
	QueryTodoRequest struct {
		ID          *uuid.UUID `json:"id" validate:"omitempty,uuid4" form:"id" query:"id"`
		Title       *string    `json:"title" validate:"required" form:"title" query:"title"`
		Description *string    `json:"description" validate:"omitempty" form:"description" query:"description"`

		IDs        []uuid.UUID `json:"ids,omitempty" form:"ids" query:"ids"`
		Limit      *int        `json:"limit,omitempty" form:"limit" query:"limit"`
		Page       *int        `json:"page,omitempty" form:"page" query:"page"`
		OrderBy    *string     `json:"order_by,omitempty" form:"order_by" query:"order_by"`
		Asc        *bool       `json:"asc,omitempty" form:"asc" query:"asc"`
		Pagination *Pagination `json:"-"`
		SortMethod *SortMethod `json:"-"`
	}
)

// Pagination struct
type Pagination struct {
	Limit  int `json:"limit" query:"limit" validate:"gte=-1,lte=100"`
	Offset int `json:"offset" query:"offset"`
}

// SortMethod struct
type SortMethod struct {
	Asc     bool   `json:"asc" query:"asc"`
	OrderBy string `json:"order_by" query:"order_by"`
}
