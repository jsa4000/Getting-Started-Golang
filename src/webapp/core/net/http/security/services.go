package security

import (
	"context"
	"net/http"
)

// AuthHandler handler to manage the authorization methods available
type AuthHandler func(w http.ResponseWriter, r *http.Request, config *Config) error

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
