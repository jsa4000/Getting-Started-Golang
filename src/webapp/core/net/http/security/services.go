package security

import "context"

// TokenData data structure for token generation (claims)
type TokenData map[string]interface{}

// UserFetcher Interface
type UserFetcher interface {
	Fetch(ctx context.Context, userID string) (*UserData, error)
}

// TokenEnhancer Interface
type TokenEnhancer interface {
	Write(t *TokenData)
}

// TokenService Interface for Users
type TokenService interface {
	Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
	Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
