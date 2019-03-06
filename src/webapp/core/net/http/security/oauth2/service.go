package oauth2

import "context"

// Service Interface for oauth2
type Service interface {
	Token(ctx context.Context, req *BasicOauth2Request) (*BasicOauth2Response, error)
	Authorize(ctx context.Context, req *BasicOauth2Request) (*BasicOauth2Response, error)
	Check(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error)
}
