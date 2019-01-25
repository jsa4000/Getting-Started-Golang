package users

import (
	"encoding/json"
	"net/http"

	errors "webapp/core/errors"
	log "webapp/core/logging"
	"webapp/core/net"
	valid "webapp/core/validation"
)

// RestController for http transport
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
func (c *RestController) GetRoutes() []net.HTTPRoute {
	return []net.HTTPRoute{
		net.HTTPRoute{
			Path:    "/users",
			Method:  "GET",
			Handler: c.GetAll,
		},
		net.HTTPRoute{
			Path:    "/users/{id}",
			Method:  "GET",
			Handler: c.GetByID,
		},
		net.HTTPRoute{
			Path:    "/users",
			Method:  "POST",
			Handler: c.Create,
		},
		net.HTTPRoute{
			Path:    "/users/{id}",
			Method:  "DELETE",
			Handler: c.DeleteByID,
		},
	}
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Users Controller Shutdown")
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

// GetAll handler to request the
func (c *RestController) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := c.Service.GetAll(r.Context(), &GetAllRequest{})
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Users)
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
	json.NewEncoder(w).Encode(res.User)
}

// Create handler to request the
func (c *RestController) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		err = ErrBadRequest.From(err)
		c.WriteError(w, err)
		return
	}

	valid, err := valid.Validate(&req)
	if !valid && err != nil {
		err = ErrBadRequest.From(err)
		c.WriteError(w, err)
		return
	}

	res, err := c.Service.Create(r.Context(), &req)
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.User)
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
