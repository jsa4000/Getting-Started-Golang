package http

import (
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

// JSON Sets the JSON response
func (c *RestController) JSON(w http.ResponseWriter, body interface{}, code int) {
	JSON(w, body, code)
}

// URL Sets the data encoded into the request url
func (c *RestController) URL(r *http.Request, data interface{}) {
	URL(r, data)
}

// Decode decode and validates. Also it sets the error for upper layers
func (c *RestController) Decode(r *http.Request, body interface{}) error {
	if r.Header.Get(HeaderContentType) == JSONMime {
		if err := DecodeJSON(r, body); err != nil {
			return ErrBadRequest.From(err)
		}
	}
	if err := DecodeParams(r, body); err != nil {
		return err
	}
	if err := DecodeVars(r, body); err != nil {
		return err
	}
	if valid, err := valid.Validate(body); !valid && err != nil {
		return ErrBadRequest.From(err)
	}
	return nil
}
