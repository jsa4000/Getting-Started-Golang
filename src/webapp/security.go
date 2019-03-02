package main

import (
	"webapp/core/net/http"
	"webapp/core/net/http/security"
	"webapp/core/net/http/security/basic"
	"webapp/core/net/http/security/open"
	"webapp/core/net/http/security/scopes"
	"webapp/core/net/http/security/token/jwt"
)

func jwtHandler() security.AuthHandler {
	return jwt.NewBuilder().
		WithTargets().
		WithURL("/*").
		And().
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
		WithTargets().
		WithURL("/oauth/*").
		And().
		Build()
}

func openAuthHandler() security.AuthHandler {
	return open.NewBuilder().
		WithTargets().
		WithURL("/swaggerui/*").
		And().
		Build()
}

func authManager(service security.UserInfoService) *security.AuthManager {
	return security.NewAuthManagerBuilder().
		WithInMemoryUsers().
		WithUser("client-trusted").WithPassword("mypassword$").WithRoles([]string{"ADMIN", "WRITE", "READ"}).
		WithUser("client-readonly").WithPassword("mypassword$").WithRoles([]string{"READ"}).
		And().
		WithUserService(service).
		Build()
}

func scopesAuthHandler() security.AuthHandler {
	return scopes.NewBuilder().
		WithTargets().
		WithURL("/users").
		WithURL("/oauth").
		And().
		Build()
}

// Security creates the security model
func Security(us security.UserInfoService) http.Security {
	authManager := authManager(us)
	return security.NewBuilder().
		WithTokenService(jwtService(us)).
		WithAuthorization(openAuthHandler(), basicAuthHandler(authManager), jwtHandler()).
		WithResourceFilter(scopesAuthHandler()).
		Build()
}
