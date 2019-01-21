package govalidator

// GoDoc: https://godoc.org/gopkg.in/go-playground/validator.v9

import (
	valid "gopkg.in/go-playground/validator.v9"
)

// Validator interface
type Validator struct {
	Validate *valid.Validate
}

// New Creates a new govalidator instance
func New() *Validator {
	return &Validator{
		Validate: valid.New(),
	}
}

// ValidateStruct validates the struct using reflection
func (v *Validator) ValidateStruct(data interface{}) (bool, error) {
	err := v.Validate.Struct(data)
	if err != nil {
		return false, err
	}
	return true, err
}
