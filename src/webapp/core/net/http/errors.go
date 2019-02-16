package http

import (
	"net/http"
	ex "webapp/core/errors"
)

// Link With HTTP errors in net/http
// https://golang.org/src/net/http/status.go

// ErrInternalServer Internal Server Error
var ErrInternalServer = ex.New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

// ErrBadRequest Bad Request
var ErrBadRequest = ex.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

// ErrNotFound Resource not found
var ErrNotFound = ex.New(http.StatusText(http.StatusNotFound), http.StatusNotFound)

// ErrUnauthorized request Unauthorized
var ErrUnauthorized = ex.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

// ErrForbidden Forbidden for resource
var ErrForbidden = ex.New(http.StatusText(http.StatusForbidden), http.StatusForbidden)
