package access

import (
	"net/http"
	"strconv"
	"strings"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
)

// AuthHandler struct to handle access control methods
type AuthHandler struct {
	security.BaseHandler
	*Config
}

// Handle handler to manage access control methods
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request, target security.Target) error {
	log.Debugf("Handle Access Request for %s", net.RemoveURLParams(r.RequestURI))
	if access, ok := target.(*Target); ok && access.Allow {
		cors(w, access.Origin)
		methods(w, access.Methods)
		headers(w, access.Headers)
		if access.Crendentials != -1 {
			credentials(w, access.Crendentials == 1)
		}
	}
	return nil
}

func methods(w http.ResponseWriter, methods []string) {
	if len(methods) > 0 {
		w.Header().Set(HeaderAccessControlAllowMethods, strings.Join(methods, ", "))
	}
}

func headers(w http.ResponseWriter, headers []string) {
	if len(headers) > 0 {
		w.Header().Set(HeaderAccessControlAllowHeaders, strings.Join(headers, ", "))
	}
}

func credentials(w http.ResponseWriter, allow bool) {
	w.Header().Set(HeaderAccessControlAllowCredentials, strconv.FormatBool(allow))

}

func cors(w http.ResponseWriter, origin string) {
	w.Header().Set(HeaderAccessControlAllowOrigin, origin)
}
