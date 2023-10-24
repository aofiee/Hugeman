package validator

import (
	validators "github.com/go-playground/validator/v10"
)

// Validator interface
type Validator interface {
	ValidateStruct(inf interface{}) error
}

type validator struct {
	validator *validators.Validate
}

// New Validator func
func New() Validator {
	v := validators.New()
	return &validator{
		validator: v,
	}
}

// ValidateStruct func
func (v *validator) ValidateStruct(inf interface{}) error {

	return v.validator.Struct(inf)
}
