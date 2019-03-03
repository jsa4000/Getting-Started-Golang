package oauth

import (
	"net/http"

	net "webapp/core/net/http"
	"webapp/core/net/http/security/token"
)

// RestController for http transport
type RestController struct {
	net.RestController
	Service token.Service
}

// NewRestController create new RestController
func NewRestController(service token.Service) net.Controller {
	return &RestController{
		Service: service,
	}
}

// GetRoutes gracefully shutdown rest controller
func (c *RestController) GetRoutes() []net.Route {
	return []net.Route{
		net.Route{
			Path:    "/auth/token",
			Method:  "POST",
			Handler: c.CreateToken,
		},
		net.Route{
			Path:    "/auth/check_token",
			Method:  "GET",
			Handler: c.CheckToken,
		},
	}
}

// CreateToken handler to request the
func (c *RestController) CreateToken(w http.ResponseWriter, r *http.Request) {
	var req CreateTokenRequest
	if err := c.Decode(w, r, &req); err != nil {
		c.Error(w, err)
		return
	}
	res, err := c.Service.Create(r.Context(), &token.CreateTokenRequest{
		UserName: req.UserName,
	})
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, res, http.StatusOK)
}

// CheckToken handler to request the
func (c *RestController) CheckToken(w http.ResponseWriter, r *http.Request) {
	req := CheckTokenRequest{Token: r.FormValue("token")}
	res, err := c.Service.Check(r.Context(), &token.CheckTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, res.Data, http.StatusOK)
}
