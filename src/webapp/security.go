package main

import (
	"webapp/core/net/http"
	"webapp/core/net/http/security"
	"webapp/core/net/http/security/basic"
	"webapp/core/net/http/security/open"
	"webapp/core/net/http/security/scopes"
	"webapp/core/net/http/security/token/jwt"
)

var (
	authManager = AuthManager()
)

func jwtHandler() security.AuthHandler {
	return jwt.NewBuilder().
		WithTargets([]string{"/*"}...).
		Build()
}

func jwtService(provider security.UserInfoService) *jwt.Service {
	return jwt.NewServiceBuilder().
		WithUserInfoService(provider).
		Build()
}

func basicAuthHandler(authManager *security.AuthManager) security.AuthHandler {
	return basic.NewBuilder().
		WithUserInfoService(authManager).
		WithTargets([]string{"/oauth/*"}...).
		Build()
}

func openAuthHandler() security.AuthHandler {
	return open.NewBuilder().
		WithTargets([]string{"/swaggerui/*"}...).
		Build()
}

func AuthManager() *security.AuthManager {
	return security.NewAuthManagerBuilder().
		WithInMemoryUsers().
		WithUser("client-trusted").WithPassword("mypassword$").WithRoles([]string{"ADMIN", "WRITE", "READ"}).
		WithUser("client-readonly").WithPassword("mypassword$").WithRoles([]string{"READ"}).
		And().
		Build()
}

func scopesAuthHandler() security.AuthHandler {
	return scopes.NewBuilder().
		WithTargets([]string{"/users", "/oauth"}...).
		Build()
}

// Security creates the security model
func Security(provider security.UserInfoService) http.Security {
	return security.NewBuilder().
		WithTokenService(jwtService(provider)).
		WithAuthorization(openAuthHandler(), basicAuthHandler(authManager), jwtHandler()).
		WithResourceFilter(scopesAuthHandler()).
		Build()
}
