package http

import (
	"encoding/json"
	"io"
	"net/http"

	valid "webapp/core/validation"
)

// RestController for http transport
type RestController struct {
}

// Error Sets the error from inner layers
func (c *RestController) Error(w http.ResponseWriter, err error) {
	Error(w, err)
}

// JSON Sets the error from inner layers
func (c *RestController) JSON(w http.ResponseWriter, body interface{}, code int) {
	JSON(w, body, code)
}

// Decode decode and validates. Also it sets the error for upper layers
func (c *RestController) Decode(r *http.Request, body interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(body)
	if err != nil && err != io.EOF {
		return ErrBadRequest.From(err)
	}
	valid, err := valid.Validate(body)
	if !valid && err != nil {
		return ErrBadRequest.From(err)
	}
	return nil
}

// DecodeParams decode the form params form the request
func (c *RestController) DecodeParams(r *http.Request, body interface{}, tag string) error {
	return DecodeParams(r, body, tag)
}
