package validation

// Validator interface
type Validator interface {
	Validate(data interface{}) (bool, error)
}

// Validator Global logger
var validator Validator

// SetGlobal sets the Global Validator (singletone)
func SetGlobal(v Validator) {
	validator = v
}

// ValidateStruct validates the struct using reflection
func Validate(data interface{}) (bool, error) {
	return validator.Validate(data)
}
