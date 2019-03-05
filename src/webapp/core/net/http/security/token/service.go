package token

import (
	"context"
)

// CreateTokenRequest request
type CreateTokenRequest struct {
	UserName string `json:"username,required"`
	Scope    string `json:"scope"`
}

// CreateTokenResponse Response
type CreateTokenResponse struct {
	AccessToken    string
	RefreshToken   string
	TokenType      string
	ExpirationTime int
}

// CheckTokenRequest struct request
type CheckTokenRequest struct {
	Token string
}

// CheckTokenResponse struct Response
type CheckTokenResponse struct {
	Data  interface{}
	Valid bool
}

// Service Interface for Users
type Service interface {
	Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
	Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
