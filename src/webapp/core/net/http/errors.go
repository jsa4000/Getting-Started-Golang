package http

import (
	"net/http"
	cerror "webapp/core/errors"
)

// Link With HTTP errors in net/http
// https://golang.org/src/net/http/status.go

var (
	// ErrInternalServer Internal Server Error
	ErrInternalServer = cerror.New(http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

	// ErrBadRequest Bad Request
	ErrBadRequest = cerror.New(http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

	// ErrNotFound Resource not found
	ErrNotFound = cerror.New(http.StatusText(http.StatusNotFound), http.StatusNotFound)

	// ErrUnauthorized request Unauthorized
	ErrUnauthorized = cerror.New(http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

	// ErrForbidden Forbidden for resource
	ErrForbidden = cerror.New(http.StatusText(http.StatusForbidden), http.StatusForbidden)
)
