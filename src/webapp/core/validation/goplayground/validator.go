package goplayground

// GoDoc: https://godoc.org/gopkg.in/go-playground/validator.v9

import (
	valid "gopkg.in/go-playground/validator.v9"
)

// Validator interface
type Validator struct {
	validator *valid.Validate
}

// New Creates a new govalidator instance
func New() *Validator {
	return &Validator{
		validator: valid.New(),
	}
}

// Validate validates the struct using reflection
func (v *Validator) Validate(data interface{}) (bool, error) {
	err := v.validator.Struct(data)
	if err != nil {
		return false, err
	}
	return true, err
}
