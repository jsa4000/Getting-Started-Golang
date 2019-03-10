package oauth2

import (
	"net/http"
	net "webapp/core/net/http"
)

// RestController for http transport
type RestController struct {
	net.RestController
	service Service
}

// NewRestController create new RestController
func NewRestController(s Service) net.Controller {
	return &RestController{
		service: s,
	}
}

// GetRoutes gracefully shutdown rest controller
func (c *RestController) GetRoutes() []net.Route {
	return []net.Route{
		net.Route{
			Path:    "/oauth/token",
			Method:  "POST",
			Handler: c.Token,
		},
		net.Route{
			Path:    "/oauth/authorize",
			Method:  "POST",
			Handler: c.Authorize,
		},
		net.Route{
			Path:    "/oauth/check_token",
			Method:  "GET",
			Handler: c.CheckToken,
		},
	}
}

// Token handler to request the
func (c *RestController) Token(w http.ResponseWriter, r *http.Request) {
	var req BasicOauth2Request
	if err := c.Decode(r, &req); err != nil {
		c.Error(w, err)
		return
	}
	if id, secret, hasAuth := r.BasicAuth(); hasAuth {
		req.ClientID = id
		req.ClientSecret = secret
	}
	res, err := c.service.Token(r.Context(), &req)
	if err != nil {
		c.Error(w, err)
		return
	}
	if len(req.RedirectURI) > 0 {
		nr, _ := http.NewRequest("GET", req.RedirectURI, nil)
		c.URL(nr, res)
		http.Redirect(w, nr, nr.URL.String(), http.StatusFound)
		return
	}
	c.JSON(w, res, http.StatusOK)
}

// Authorize handler to request the
func (c *RestController) Authorize(w http.ResponseWriter, r *http.Request) {
	var req BasicOauth2Request
	if err := c.Decode(r, &req); err != nil {
		c.Error(w, err)
		return
	}
	if id, secret, hasAuth := r.BasicAuth(); hasAuth {
		req.ClientID = id
		req.ClientSecret = secret
	}
	res, err := c.service.Authorize(r.Context(), &req)
	if err != nil {
		c.Error(w, err)
		return
	}
	if len(req.RedirectURI) > 0 {
		nr, _ := http.NewRequest("GET", req.RedirectURI, nil)
		c.URL(nr, res)
		http.Redirect(w, nr, nr.URL.String(), http.StatusFound)
		return
	}
	c.JSON(w, res, http.StatusOK)
}

// CheckToken handler to request the
func (c *RestController) CheckToken(w http.ResponseWriter, r *http.Request) {
	var req CheckTokenRequest
	if err := c.Decode(r, &req); err != nil {
		c.Error(w, err)
		return
	}
	if id, secret, hasAuth := r.BasicAuth(); hasAuth {
		req.ClientID = id
		req.ClientSecret = secret
	}
	res, err := c.service.Check(r.Context(), &req)
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, res.Data, http.StatusOK)
}
