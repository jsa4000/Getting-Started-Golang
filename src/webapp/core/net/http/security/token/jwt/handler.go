package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"

	jwt "github.com/dgrijalva/jwt-go"
)

// AuthHandler Implementation used for the service
type AuthHandler struct {
	security.BaseHandler
	*Config
	provider security.UserInfoService
}

// Handle handler to authorize the JWT method
func (s *AuthHandler) Handle(w http.ResponseWriter, r *http.Request, target security.Target) error {
	log.Debugf("Handle JWT Request for %s", net.RemoveURLParams(r.RequestURI))
	basicAuth, ok := r.Header[authHeader]
	if !ok {
		return net.ErrUnauthorized.From(errors.New("Authorization is required"))
	}
	token, err := s.verify(r.Context(), strings.TrimPrefix(basicAuth[0], bearerPreffix))
	if err != nil {
		return err
	}
	security.SetContextValue(r, AuthKey, new(security.ContextValue))
	claims := token.Claims.(jwt.MapClaims)
	security.SetUserName(r, claims[userNameField].(string))
	security.SetUserID(r, claims[subjectField].(string))
	if val, ok := claims[rolesField]; ok {
		if iroles, err := val.([]interface{}); !err {
			roles := make([]string, len(iroles))
			for _, role := range iroles {
				roles = append(roles, role.(string))
			}
			security.SetUserRoles(r, roles)
		}
	}
	return nil
}

// Check returns deserialized token
func (s *AuthHandler) verify(ctx context.Context, val string) (*jwt.Token, error) {
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return nil, net.ErrForbidden.From(err)
	}
	return token, nil
}
