package oauth

import (
	"context"
	"fmt"
	"net/http"
	"webapp/core/net/http/security/token"
)

// Decode retrieve params from header or paramrs, not in the body
func Decode(r *http.Request, req *CreateTokenRequest) error {
	client, secret, hasAuth := r.BasicAuth()
	if hasAuth {
		req.ClientID = client
		req.ClientSecret = secret
	}

	fmt.Println(r.FormValue("redirect_uri"))
	fmt.Println(r.FormValue("scope"))
	return nil
}

// Create return a token response or error depending on the parameters
func Create(ctx context.Context, service token.Service, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	fmt.Println(req)
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
