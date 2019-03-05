package oauth

import (
	"context"
	"webapp/core/net/http/security/token"
)

// ServiceImpl for http transport
type ServiceImpl struct {
	tokenService  token.Service
	clientService ClientService
}

// NewServiceImpl create new RestController
func NewServiceImpl(ts token.Service, cs ClientService) *ServiceImpl {
	return &ServiceImpl{
		tokenService:  ts,
		clientService: cs,
	}
}

// Create return a token response or error depending on the parameters
func (s *ServiceImpl) Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	res, err := s.tokenService.Create(ctx, &token.CreateTokenRequest{
		UserName: req.UserName,
		Scope:    req.Scope,
	})
	if err != nil {
		return nil, err
	}
	return &CreateTokenResponse{
		AccessToken:    res.AccessToken,
		TokenType:      res.TokenType,
		RefreshToken:   res.RefreshToken,
		ExpirationTime: res.ExpirationTime,
	}, nil
}

// Check the token passed by parameter
func (s *ServiceImpl) Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
	res, err := s.tokenService.Check(ctx, &token.CheckTokenRequest{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	return &CheckTokenResponse{
		Data: res.Data,
	}, nil
}
