package roles

import (
	"encoding/json"
	"net/http"

	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// RestController for http netport
type RestController struct {
	net.RestController
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
