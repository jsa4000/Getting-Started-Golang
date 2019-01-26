package http

import (
	"encoding/json"
	"net/http"

	errors "webapp/core/errors"
	log "webapp/core/logging"
	valid "webapp/core/validation"
)

// RestController for http transport
type RestController struct {
}

// WriteError Sets the error from inner layers
func (c *RestController) WriteError(w http.ResponseWriter, err error) {
	herr, ok := err.(*errors.Error)
	if !ok {
		herr = ErrInternalServer.From(err)
	}
	w.WriteHeader(herr.Code)
	json.NewEncoder(w).Encode(herr)
	log.Error(herr)
}

// Decode decode and validates. Also it sets the error for upper layers
func (c *RestController) Decode(w http.ResponseWriter, r *http.Request, body interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(body)
	if err != nil {
		return ErrBadRequest.From(err)
	}
	valid, err := valid.Validate(body)
	if !valid && err != nil {
		return ErrBadRequest.From(err)
	}
	return nil
}
