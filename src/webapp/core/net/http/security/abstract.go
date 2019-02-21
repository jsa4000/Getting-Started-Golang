package security

import (
	"context"
	"net/http"
)

// UserFetcher Interface
type UserFetcher interface {
	Fetch(ctx context.Context, userID string) (*UserData, error)
}

// TokenEnhancer Interface
type TokenEnhancer interface {
	Write(t *TokenData)
}

// AuthHandler interface to manage the authorization method
type AuthHandler interface {
	Handle(w http.ResponseWriter, r *http.Request) error
	Targets() []string
}

// TokenService Interface for Users
type TokenService interface {
	Create(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error)
	Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
