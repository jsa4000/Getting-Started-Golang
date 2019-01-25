package users

import ex "webapp/core/exceptions"

// ErrServer Error connection to repository
var ErrServer = ex.New("Error: Internal Server Error", 500)

// ErrValidation Error connection to repository
var ErrValidation = ex.New("Error: Bad Request", 400)

// ErrNotFound Error resource not found
var ErrNotFound = ex.New("Error: Resource Not Found", 404)
