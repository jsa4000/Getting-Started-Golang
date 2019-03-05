package oauth

import "context"

// Service Interface for Users
type Service interface {
	Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
	Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
