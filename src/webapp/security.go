package main

import (
	"webapp/core/net/http"
	"webapp/core/net/http/security"
	"webapp/core/net/http/security/access"
	"webapp/core/net/http/security/basic"
	"webapp/core/net/http/security/oauth"
	"webapp/core/net/http/security/open"
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
		WithURL("/auth/*").
		WithAuthority("ADMIN", "WRITE", "READ").
		And().
		Build()
}

func corsAuthFilter() security.AuthHandler {
	return access.NewBuilder().
		WithTargets().
		WithURL("/auth/*").WithOrigin("example.domain.com").Allow().
		WithURL("/users").WithOrigin("*").Allow().
		WithURL("/roles").
		And().
		Build()
}

func openAuthHandler() security.AuthHandler {
	return open.NewBuilder().
		WithTargets().
		WithURL("/swaggerui/*").
		WithURL("/debug/pprof/").
		And().
		Build()
}

func oAuthManager(jwt *jwt.Service) *oauth.Manager {
	return oauth.NewManagerBuilder().
		WithInMemoryClients().
		WithClient("client-trusted").WithSecret("mypassword$").WithScope("admin", "roles", "users").
		WithClient("client-readonly").WithSecret("mypassword$").WithScope("read").
		And().
		WithTokenService(jwt).
		Build()
}

func authManager(service security.UserInfoService) *security.AuthManager {
	return security.NewAuthManagerBuilder().
		WithInMemoryUsers().
		WithUser("user-trusted").WithPassword("mypassword$").WithRole("ADMIN", "WRITE", "READ").
		WithUser("user-readonly").WithPassword("mypassword$").WithRole("READ").
		And().
		WithUserService(service).
		Build()
}

// Security creates the security model
func Security(us security.UserInfoService) http.Security {
	authManager := authManager(us)
	jwtService := jwtService(authManager)
	return security.NewBuilder().
		WithAuthentication(oAuthManager(jwtService)).
		WithAuthorization(openAuthHandler(), basicAuthHandler(authManager), jwtHandler()).
		WithFilter(corsAuthFilter()).
		Build()
}