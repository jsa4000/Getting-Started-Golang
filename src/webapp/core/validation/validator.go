package validation

// Validator interface
type Validator interface {
	ValidateStruct(data interface{}) (bool, error)
}

// Validator Global logger
var validator Validator

// SetGlobal sets the Global Validator (singletone)
func SetGlobal(v Validator) {
	validator = v
}

// ValidateStruct validates the struct using reflection
func ValidateStruct(data interface{}) (bool, error) {
	return validator.ValidateStruct(data)
}
