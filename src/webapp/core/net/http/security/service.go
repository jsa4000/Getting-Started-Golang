package security

import "context"

// Service Interface for Users
type Service interface {
	CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
	CheckToken(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
