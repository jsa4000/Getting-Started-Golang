package oauth

import (
	"context"
	"webapp/core/net/http/security/token"
)

// Create return a token response or error depending on the parameters
func Create(ctx context.Context, service token.Service, req *CreateTokenRequest) (*CreateTokenResponse, error) {

	res, err := service.Create(ctx, &token.CreateTokenRequest{
		UserName: req.UserName,
	})
	if err != nil {
		return nil, err
	}
	return &CreateTokenResponse{
		AccessToken:    res.AccessToken,
		TokenType:      res.TokenType,
		RefreshToken:   res.RefreshToken,
		ExpirationTime: res.ExpirationTime,
	}, nil

}
