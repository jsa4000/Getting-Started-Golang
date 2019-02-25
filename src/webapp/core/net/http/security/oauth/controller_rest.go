package oauth

import (
	"encoding/json"
	"net/http"

	net "webapp/core/net/http"
)

// RestController for http transport
type RestController struct {
	net.RestController
	Service Service
}

// NewRestController create new RestController
func NewRestController(service Service) net.Controller {
	return &RestController{
		Service: service,
	}
}

// GetRoutes gracefully shutdown rest controller
func (c *RestController) GetRoutes() []net.Route {
	return []net.Route{
		net.Route{
			Path:    "/oauth/token",
			Method:  "POST",
			Handler: c.CreateToken,
		},
		net.Route{
			Path:    "/oauth/check_token",
			Method:  "GET",
			Handler: c.CheckToken,
		},
	}
}

// CreateToken handler to request the
func (c *RestController) CreateToken(w http.ResponseWriter, r *http.Request) {
	var req CreateTokenRequest
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
	json.NewEncoder(w).Encode(res)
}

// CheckToken handler to request the
func (c *RestController) CheckToken(w http.ResponseWriter, r *http.Request) {
	req := CheckTokenRequest{Token: r.FormValue("token")}
	res, err := c.Service.Check(r.Context(), &req)
	if err != nil {
		c.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Data)
}
