package mngmt

import (
	"net/http"
	net "webapp/core/net/http"
)

// RestController for http transport
type RestController struct {
	net.RestController
	service Service
	path    string
}

// NewRestController create new RestController
func NewRestController(service Service, path string) net.Controller {
	return &RestController{
		service: service,
		path:    path,
	}
}

// GetRoutes gracefully shutdown rest controller
func (c *RestController) GetRoutes() []net.Route {
	return []net.Route{
		net.Route{
			Path:    c.path + "/health",
			Method:  "GET",
			Handler: c.Health,
		},
		net.Route{
			Path:    c.path + "/metrics",
			Method:  "GET",
			Handler: c.Metrics,
		},
	}
}

// Health handler to request the health
func (c *RestController) Health(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.Health(r.Context(), &HealthRequest{})
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, res.Health, http.StatusOK)
}

// Metrics handler to request the metrics
func (c *RestController) Metrics(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.Metrics(r.Context(), &MetricsRequest{})
	if err != nil {
		c.Error(w, err)
		return
	}
	c.JSON(w, res.Metrics, http.StatusOK)
}
