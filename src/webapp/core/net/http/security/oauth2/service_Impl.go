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
		err = net.ErrBadRequest.From(fmt.Errorf("response_type %s not supported yet", req.GranType))
	case GrantTypePassword:
		res, err = s.tokenService.Create(ctx, &token.CreateTokenRequest{
			UserName: req.UserName,
			Scope:    req.Scope,
		})
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
		State:          req.Scope,
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
		State:          req.Scope,
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
