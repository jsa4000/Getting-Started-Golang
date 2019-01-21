package govalidator

import (
	"github.com/asaskevich/govalidator"
	valid "github.com/asaskevich/govalidator"
)

// Validator interface
type Validator struct {
}

// New Creates a new govalidator instance
func New() *Validator {
	govalidator.SetFieldsRequiredByDefault(false)
	govalidator.SetNilPtrAllowedByRequired(false)

	return &Validator{}
}

// ValidateStruct validates the struct using reflection
func (v *Validator) ValidateStruct(data interface{}) (bool, error) {
	return valid.ValidateStruct(data)
}
