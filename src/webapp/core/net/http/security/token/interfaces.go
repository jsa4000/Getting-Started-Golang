package token

import "context"

// UserFetcher Interface
type UserFetcher interface {
	Fetch(ctx context.Context, username string) (*UserData, error)
}

// Enhancer Interface
type Enhancer interface {
	Write(t *Claims)
}

// Service Interface for Users
type Service interface {
	Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
	Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
