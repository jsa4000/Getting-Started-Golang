package exceptions

import "fmt"

// Error to describe a complete error for tracing
type Error struct {
	Message     string `json:"message"`
	Description string `json:"error"`
	HTTPCode    int    `json:"code"`
}

// New Creates a new Error instance
func New(message string, httpCode int) *Error {
	return &Error{
		Message:  message,
		HTTPCode: httpCode,
	}
}

// From Return the error for the given error
func (e *Error) From(err error) *Error {
	e.Description = err.Error()
	return e
}

// Error implementation for error interface
func (e *Error) Error() string {
	return e.String()
}

// Error implementation for error interface
func (e *Error) String() string {
	return fmt.Sprintf("error: %s, description: %s", e.Message, e.Description)
}
