package users

import "webapp/core/errors"

// ErrServer Error connection to repository
var ErrServer = errors.New("Unknown Server Error", 500)

// ErrConnRepo Error connection to repository
var ErrConnRepo = errors.New("Error Connecting to the Repository", 500)

// ErrUserNotFoud Error connection to repository
var ErrUserNotFoud = errors.New("Error Resource Not Found", 404)
