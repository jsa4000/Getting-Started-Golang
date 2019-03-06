package oauth2

import (
	"context"
	"errors"
	"fmt"

	net "webapp/core/net/http"
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

// Token return a token response or error depending on the parameters
func (s *ServiceImpl) Token(ctx context.Context, req *BasicOauth2Request) (*BasicOauth2Response, error) {
	if len(req.GranType) == 0 {
		return nil, net.ErrBadRequest.From(errors.New("grant_type is a mandatory field"))
	}
	var err error
	var res *token.CreateTokenResponse
	switch req.GranType {
	case GrantTypeAuthorizationCode:
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported yet", req.GranType))
	case GrantTypeClientCredentials:
		res, err = s.grantTypeClientCredentials(ctx, req)
	case GrantTypePassword:
		res, err = s.grantTypePassword(ctx, req)
	case GrantTypeRefreshToken:
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported yet", req.GranType))
	default:
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported", req.GranType))
	}
	if err != nil {
		return nil, err
	}
	return &BasicOauth2Response{
		AccessToken:    res.AccessToken,
		TokenType:      res.TokenType,
		RefreshToken:   res.RefreshToken,
		ExpirationTime: res.ExpirationTime,
		State:          req.State,
	}, nil
}

// Authorize return a token response or error depending on the parameters
func (s *ServiceImpl) Authorize(ctx context.Context, req *BasicOauth2Request) (*BasicOauth2Response, error) {
	if len(req.ResponseType) == 0 {
		return nil, net.ErrBadRequest.From(errors.New("response_type is a mandatory field"))
	}
	var err error
	var res *token.CreateTokenResponse
	switch req.ResponseType {
	case ResponseTypeCode:
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported yet", req.ResponseType))
	case ResponseTypeImplicit:
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported yet", req.ResponseType))
	default:
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported", req.ResponseType))
	}
	if err != nil {
		return nil, err
	}
	return &BasicOauth2Response{
		AccessToken:    res.AccessToken,
		TokenType:      res.TokenType,
		RefreshToken:   res.RefreshToken,
		ExpirationTime: res.ExpirationTime,
		State:          req.State,
	}, nil
}

// Check the token passed by parameter
func (s *ServiceImpl) Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
	client, err := s.clientService.Fetch(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}
	if !ValidateClient(client, req.ClientID, req.ClientSecret) {
		return nil, net.ErrUnauthorized.From(fmt.Errorf("Invalid credentials for client %s", req.ClientID))
	}
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

// Token return a token response or error depending on the parameters
func (s *ServiceImpl) grantTypePassword(ctx context.Context, req *BasicOauth2Request) (*token.CreateTokenResponse, error) {
	client, err := s.clientService.Fetch(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}
	if !ValidateClient(client, req.ClientID, req.ClientSecret) {
		return nil, net.ErrUnauthorized.From(fmt.Errorf("Invalid credentials for client %s", req.ClientID))
	}
	return s.tokenService.Create(ctx, &token.CreateTokenRequest{
		UserName: req.UserName,
		Scope:    req.Scope,
	})
}

// Token return a token response or error depending on the parameters
func (s *ServiceImpl) grantTypeClientCredentials(ctx context.Context, req *BasicOauth2Request) (*token.CreateTokenResponse, error) {
	client, err := s.clientService.Fetch(ctx, req.ClientID)
	if err != nil {
		return nil, err
	}
	if !ValidateClient(client, req.ClientID, req.ClientSecret) {
		return nil, net.ErrUnauthorized.From(fmt.Errorf("Invalid credentials for client %s", req.ClientID))
	}
	return s.tokenService.Create(ctx, &token.CreateTokenRequest{
		Scope: req.Scope,
	})
}
