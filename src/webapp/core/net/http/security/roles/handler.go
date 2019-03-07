package roles

import (
	"fmt"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// Handler struct to handle basic authentication
type Handler struct {
	security.BaseHandler
	*Config
}

// Handle handler to manage basic authenticaiton method
func (s *Handler) Handle(w http.ResponseWriter, r *http.Request, target security.Target) error {
	auth, exist := security.AuthType(r)
	if !exist {
		return nil
	}
	log.Debugf("Handle Roles Request for %s and auth type %s", net.RemoveURLParams(r.RequestURI), auth)
	route, exist := net.RouteInfo(r)
	if !exist || len(route.Roles) == 0 {
		return nil
	}
	userRoles, hasRoles := security.UserRoles(r)
	username, _ := security.GetUserName(r)
	if (!hasRoles || len(userRoles) == 0) && len(route.Roles) > 0 ||
		len(security.RolesMatches(userRoles, route.Roles)) != len(route.Roles) {
		return net.ErrForbidden.From(fmt.Errorf("User %s has not enough privileges", username))
	}
	return nil
}
