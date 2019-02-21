package jwt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	cerr "webapp/core/errors"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
	global "webapp/core/time"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type key string

const (
	// JwtKey Key to get data from context in basicAuth
	JwtKey key = "JwtKey"

	bearerPreffix = "Bearer "
	authHeader    = "Authorization"
)

// Service Implementation used for the service
type Service struct {
	*Config
	targets       []string
	userFetcher   security.UserFetcher
	tokenEnhancer security.TokenEnhancer
}

// Create create the token - HMAC (HS256) based with secret key
// Check https://github.com/dgrijalva/jwt-go and 'Signing Methods and Key Types'
func (s *Service) Create(ctx context.Context, req *security.CreateTokenRequest) (*security.CreateTokenResponse, error) {
	var err error
	user, err := s.userFetcher.Fetch(ctx, req.UserName)
	if err != nil {
		herr, ok := err.(*cerr.Error)
		if !ok {
			herr = net.ErrInternalServer.From(err)
		}
		return nil, herr
	}

	expirationTime := global.Now().Add(time.Second * time.Duration(s.ExpirationTime))
	claims := jwt.MapClaims{
		"jti":   uuid.NewV4().String(),
		"iss":   s.Issuer,
		"sub":   user.ID,
		"name":  user.Name,
		"roles": user.Roles,
		"exp":   expirationTime.Unix(),
		"iat":   global.Unix(),
	}
	if s.tokenEnhancer != nil {

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &security.CreateTokenResponse{
		Token:          tokenString,
		ExpirationTime: expirationTime,
	}, nil
}

// Check returns deserialized token
func (s *Service) Check(ctx context.Context, req *security.CheckTokenRequest) (*security.CheckTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok || ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) == 0 {
			return nil, net.ErrInternalServer.From(err)
		}
	}
	return &security.CheckTokenResponse{
		Data:  token,
		Valid: token.Valid,
	}, nil
}

// Handle handler to manage basic authenticaiton method
func (s *Service) Handle(w http.ResponseWriter, r *http.Request) error {
	log.Debugf("Handle JWT Request for %s", r.RequestURI)
	basicAuth, ok := r.Header[authHeader]
	if !ok {
		return net.ErrUnauthorized.From(errors.New("Authorization has not been found"))
	}
	resp, err := s.Check(r.Context(), &security.CheckTokenRequest{
		Token: strings.TrimPrefix(basicAuth[0], bearerPreffix),
	})
	if err != nil || !resp.Valid {
		return net.ErrUnauthorized.From(errors.New("Authorization Beared invalid"))
	}
	r.WithContext(context.WithValue(r.Context(), JwtKey, resp.Data))
	return nil
}

//Targets returns the targets or urls the auth applies for
func (s *Service) Targets() []string {
	return s.targets
}
