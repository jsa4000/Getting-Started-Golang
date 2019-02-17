package security

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
	cerr "webapp/core/errors"
	net "webapp/core/net/http"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// ServiceJwt Implementation used for the service
type ServiceJwt struct {
	config *Config
}

// NewServiceJwt Create a new ServiceImpl
func NewServiceJwt(config *Config) Service {
	return &ServiceJwt{
		config: config,
	}
}

func (s *ServiceJwt) validateCreateTokenRequest(req *CreateTokenRequest) error {
	if len(req.UserName) == 0 && len(req.UserEmail) == 0 {
		return net.ErrBadRequest.From(errors.New("UserName or UserEmail must not be empty"))
	}
	if len(req.Authorization) == 0 {
		return net.ErrUnauthorized.From(errors.New("Basic Authorization is required"))
	}
	if req.Authorization != basicAuth(s.config.ClientID, s.config.ClientSecret) {
		return net.ErrUnauthorized.From(errors.New("Authorization credentials are not valid"))
	}
	return nil
}

// CreateToken create the token - HMAC (HS256) based with secret key
// Check https://github.com/dgrijalva/jwt-go and 'Signing Methods and Key Types'
func (s *ServiceJwt) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	if err := s.validateCreateTokenRequest(req); err != nil {
		return nil, err
	}
	var user *UserData
	var err error
	if len(req.UserName) > 0 {
		user, err = s.config.uc.GetUserByName(ctx, req.UserName)
	} else {
		user, err = s.config.uc.GetUserByEmail(ctx, req.UserEmail)
	}
	if err != nil {
		herr, ok := err.(*cerr.Error)
		if !ok {
			herr = net.ErrInternalServer.From(err)
		}
		return nil, herr
	}

	fmt.Printf("** Expiration Time: %d\n", s.config.ExpirationTime)

	expireTime := time.Now().Add(time.Second * time.Duration(s.config.ExpirationTime))
	claims := jwt.MapClaims{
		"jti":   uuid.NewV4().String(),
		"iss":   s.config.Issuer,
		"sub":   user.ID,
		"name":  user.Name,
		"roles": user.Roles,
		"exp":   expireTime.Format(time.RFC3339),
		"iat":   time.Now().Format(time.RFC3339),
	}
	if s.config.tc != nil {

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.SecretKey))
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &CreateTokenResponse{
		Token:      tokenString,
		ExpireTime: expireTime,
	}, nil
}

// CheckToken returns deserialized token
func (s *ServiceJwt) CheckToken(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.SecretKey), nil
	})
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &CheckTokenResponse{
		Data: token,
	}, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
