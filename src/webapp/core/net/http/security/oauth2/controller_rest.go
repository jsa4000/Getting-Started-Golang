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
	if err := c.DecodeParams(r, &req, "json"); err != nil {
		c.Error(w, err)
		return
	}
	if client, secret, hasAuth := r.BasicAuth(); hasAuth {
		req.ClientID = client
		req.ClientSecret = secret
	}
	res, err := c.service.Token(r.Context(), &req)
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, &BasicOauth2Response{
		AccessToken:    res.AccessToken,
		TokenType:      res.TokenType,
		RefreshToken:   res.RefreshToken,
		ExpirationTime: res.ExpirationTime,
	}, http.StatusOK)
}

// Authorize handler to request the
func (c *RestController) Authorize(w http.ResponseWriter, r *http.Request) {
	var req BasicOauth2Request
	if err := c.Decode(r, &req); err != nil {
		c.Error(w, err)
		return
	}
	if err := c.DecodeParams(r, &req, "json"); err != nil {
		c.Error(w, err)
		return
	}
	if client, secret, hasAuth := r.BasicAuth(); hasAuth {
		req.ClientID = client
		req.ClientSecret = secret
	}
	res, err := c.service.Authorize(r.Context(), &req)
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, &BasicOauth2Response{
		AccessToken:    res.AccessToken,
		TokenType:      res.TokenType,
		RefreshToken:   res.RefreshToken,
		ExpirationTime: res.ExpirationTime,
	}, http.StatusOK)
}

// CheckToken handler to request the
func (c *RestController) CheckToken(w http.ResponseWriter, r *http.Request) {
	req := CheckTokenRequest{Token: r.FormValue("token")}
	res, err := c.service.Check(r.Context(), &req)
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, res.Data, http.StatusOK)
}
