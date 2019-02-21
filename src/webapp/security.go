package main

import (
	"webapp/core/net/http"
	"webapp/core/net/http/security"
	"webapp/core/net/http/security/basic"
	"webapp/core/net/http/security/jwt"
	"webapp/core/net/http/security/open"
	"webapp/core/net/http/security/role"
)

func jwtService(uf security.UserFetcher) *jwt.Service {
	return jwt.NewBuilder().
		WithUserFetcher(uf).
		WithTargets([]string{"/*"}...).
		Build()
}

func basicAuthService(uf security.UserFetcher) security.AuthHandler {
	return basic.NewBuilder().
		WithUserFetcher(uf).
		WithTargets([]string{"/oauth/*"}...).
		Build()
}

func openAuthService() security.AuthHandler {
	return open.NewBuilder().
		WithTargets([]string{"/swagger/*"}...).
		Build()
}

func roleAuthService() security.AuthHandler {
	return role.NewBuilder().
		WithTargets([]string{"/users"}...).
		Build()
}

// Security creates the security model
func Security(uf security.UserFetcher) http.Security {
	jwtService := jwtService(uf)
	return security.NewBuilder().
		WithTokenService(jwtService).
		WithAuthenticationHandlers(openAuthService(), basicAuthService(uf), jwtService).
		WithAuthorizationHandlers(roleAuthService()).
		Build()
}
