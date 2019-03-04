package jwt

import (
	"context"
	"fmt"
	"strings"
	"time"
	cerr "webapp/core/errors"
	net "webapp/core/net/http"
	"webapp/core/net/http/security"
	service "webapp/core/net/http/security/token"
	global "webapp/core/time"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Service Implementation used for the service
type Service struct {
	*Config
	provider security.UserInfoService
	enhancer ClaimsEnhancer
}

// Create create the token - HMAC (HS256) based with secret key
// Check https://github.com/dgrijalva/jwt-go and 'Signing Methods and Key Types'
func (s *Service) Create(ctx context.Context, req *service.CreateTokenRequest) (*service.CreateTokenResponse, error) {
	var err error
	user, err := s.provider.Fetch(ctx, req.UserName)
	if err != nil {
		herr, ok := err.(*cerr.Error)
		if !ok {
			herr = net.ErrInternalServer.From(err)
		}
		return nil, herr
	}

	expirationTime := global.Now().Add(time.Second * time.Duration(s.ExpirationTime))
	issueAt := global.Now()

	claims := jwt.MapClaims{
		jsontokenIDfield:    uuid.NewV4().String(),
		issuerField:         s.Issuer,
		subjectField:        user.ID,
		userNameField:       user.Name,
		rolesField:          user.Roles,
		expirationDateField: expirationTime.Unix(),
		issuedAtField:       issueAt.Unix(),
	}
	if s.enhancer != nil {
		s.enhancer.Write(Claims(claims), user)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &service.CreateTokenResponse{
		AccessToken:    tokenString,
		TokenType:      strings.TrimSpace(strings.ToLower(bearerPreffix)),
		ExpirationTime: s.ExpirationTime,
	}, nil
}

// Check returns deserialized token
func (s *Service) Check(ctx context.Context, req *service.CheckTokenRequest) (*service.CheckTokenResponse, error) {
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
	return &service.CheckTokenResponse{
		Data:  token,
		Valid: token.Valid,
	}, nil
}
