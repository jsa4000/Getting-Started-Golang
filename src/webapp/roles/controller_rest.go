package roles

import (
	"encoding/json"
	"net/http"

	errors "webapp/core/errors"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	valid "webapp/core/validation"
)

// RestController for http netport
type RestController struct {
	Service Service
}

// NewRestController create new RestController
func NewRestController(service Service) *RestController {
	return &RestController{
		Service: service,
	}
}

// GetRoutes gracefully shutdown rest controller
func (c *RestController) GetRoutes() []net.Route {
	return []net.Route{
		net.Route{
			Path:    "/roles",
			Method:  "GET",
			Handler: c.GetAll,
		},
		net.Route{
			Path:    "/roles/{id}",
			Method:  "GET",
			Handler: c.GetByID,
		},
		net.Route{
			Path:    "/roles",
			Method:  "POST",
			Handler: c.Create,
		},
		net.Route{
			Path:    "/roles/{id}",
			Method:  "DELETE",
			Handler: c.DeleteByID,
		},
	}
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Role Controller Shutdown")
}

// WriteError Sets the error from inner layers
func (c *RestController) WriteError(w http.ResponseWriter, err error) {
	herr, ok := err.(*errors.Error)
	if !ok {
		herr = net.ErrInternalServer.From(err)
	}
	w.WriteHeader(herr.Code)
	json.NewEncoder(w).Encode(herr)
	log.Error(herr)
}

// Decode Sets the error from inner layers
func (c *RestController) Decode(w http.ResponseWriter, r *http.Request, body interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(body)
	if err != nil {
		return net.ErrBadRequest.From(err)
	}
	valid, err := valid.Validate(body)
	if !valid && err != nil {
		return net.ErrBadRequest.From(err)
	}
	return nil
}

// GetAll handler to request the
func (c *RestController) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := c.Service.GetAll(r.Context(), &GetAllRequest{})
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Roles)
}

// GetByID handler to request the
func (c *RestController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := net.GetVars(r)
	req := GetByIDRequest{ID: vars["id"]}
	res, err := c.Service.GetByID(r.Context(), &req)
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Role)
}

// Create handler to request the
func (c *RestController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest
	if err := c.Decode(w, r, &req); err != nil {
		c.WriteError(w, err)
		return
	}
	res, err := c.Service.Create(r.Context(), &req)
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Role)
}

// DeleteByID handler to request the
func (c *RestController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := net.GetVars(r)
	req := DeleteByIDRequest{ID: vars["id"]}
	_, err := c.Service.DeleteByID(r.Context(), &req)
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
