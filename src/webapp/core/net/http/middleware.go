package http

import (
	"net/http"
)

// Middleware for handle the requests
type Middleware func(http.Handler) http.Handler
