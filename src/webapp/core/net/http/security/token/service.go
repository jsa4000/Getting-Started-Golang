package token

import (
	"context"
	"time"
)

// CreateTokenRequest request
type CreateTokenRequest struct {
	UserName string `json:"username"`
}

// CreateTokenResponse Response
type CreateTokenResponse struct {
	Token          string
	ExpirationTime time.Duration
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
