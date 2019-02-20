package security

import (
	"context"
	"fmt"
	"time"
	cerr "webapp/core/errors"
	net "webapp/core/net/http"
	global "webapp/core/time"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// TokenServiceJwt Implementation used for the service
type TokenServiceJwt struct {
	*Config
}

// NewTokenServiceJwt Create a new ServiceImpl
func NewTokenServiceJwt(config *Config) *TokenServiceJwt {
	return &TokenServiceJwt{config}
}

// Create create the token - HMAC (HS256) based with secret key
// Check https://github.com/dgrijalva/jwt-go and 'Signing Methods and Key Types'
func (s *TokenServiceJwt) Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	var err error
	user, err := s.UserFetcher.Fetch(ctx, req.UserName)
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
	if s.TokenEnhancer != nil {

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &CreateTokenResponse{
		Token:          tokenString,
		ExpirationTime: expirationTime,
	}, nil
}

// Check returns deserialized token
func (s *TokenServiceJwt) Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
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
	return &CheckTokenResponse{
		Data:  token,
		Valid: token.Valid,
	}, nil
}
