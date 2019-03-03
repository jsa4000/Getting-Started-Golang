package http

import (
	"encoding/json"
	"net/http"
	"webapp/core/errors"
	log "webapp/core/logging"
)

// Error Sets the error from inner layers
func Error(w http.ResponseWriter, err error) {
	herr, ok := err.(*errors.Error)
	if !ok {
		herr = ErrInternalServer.From(err)
	}
	JSON(w, herr, herr.Code)
	log.Error(herr)
}

// JSON Sets the error from inner layers
func JSON(w http.ResponseWriter, body interface{}, code int) {
	w.Header().Set(HeaderContentType, JSONMime)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}
