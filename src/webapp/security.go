package main

import (
	"webapp/core/net/http"
	"webapp/core/net/http/security"
	"webapp/core/net/http/security/access"
	"webapp/core/net/http/security/basic"
	"webapp/core/net/http/security/oauth2"
	"webapp/core/net/http/security/open"
	"webapp/core/net/http/security/token/jwt"
	"webapp/core/net/http/security/users"
)

func jwtHandler() security.AuthHandler {
	return jwt.NewBuilder().
		WithTargets().
		WithURL("/*").
		And().
		WithPriority(2).
		Build()
}

func basicAuthHandler(usersManager *users.Manager) security.AuthHandler {
	return basic.NewBuilder().
		WithUserInfoService(usersManager).
		WithTargets().
		WithURL("/management/*").
		WithURL("/swaggerui/*").
		WithAuthority("ADMIN", "WRITE", "READ").
		And().
		WithPriority(1).
		Build()
}

func openAuthHandler() security.AuthHandler {
	return open.NewBuilder().
		WithTargets().
		//WithURL("/debug/pprof/").
		WithURL("/oauth/*").
		And().
		WithPriority(0).
		Build()
}

func corsAuthFilter() security.AuthHandler {
	return access.NewBuilder().
		WithTargets().
		WithURL("/oauth/*").WithOrigin("example.domain.com").WithCredentials(true).Allow().
		WithURL("/users").WithOrigin("*").WithMethods("POST", "GET", "OPTIONS", "PUT", "DELETE").Allow().
		WithURL("/roles").WithCredentials(false).Allow().
		And().
		Build()
}

func oAuthManager(jwt *jwt.Service) *oauth2.Manager {
	return oauth2.NewManagerBuilder().
		WithInMemoryClients().
		WithClient("client-trusted").WithSecret("mypassword$").WithScope("read", "write", "admin").
		WithClient("client-readonly").WithSecret("mypassword$").WithScope("read").
		And().
		WithTokenService(jwt).
		Build()
}

func usersManager(service security.UserInfoService) *users.Manager {
	return users.NewManagerBuilder().
		WithInMemoryUsers().
		WithUser("user-trusted").WithPassword("mypassword$").WithRole("ADMIN", "WRITE", "READ").
		WithUser("user-readonly").WithPassword("mypassword$").WithRole("READ").
		And().
		WithUserService(service).
		Build()
}

type tags struct{}

func (t *tags) Write(c jwt.Claims, u *security.UserInfo) {
	c["region"] = "eu-west-1"
	c["tags"] = []string{"webapp", "secuity", "token"}
}

func jwtService(provider security.UserInfoService) *jwt.Service {
	return jwt.NewServiceBuilder().
		WithUserInfoService(provider).
		WithClaimsEnhancer(&tags{}).
		Build()
}

// Security creates the security model
func Security(us security.UserInfoService) http.Security {
	usersManager := usersManager(us)
	jwtService := jwtService(usersManager)
	return security.NewBuilder().
		WithAuthentication(oAuthManager(jwtService)).
		WithAuthorization(openAuthHandler(), basicAuthHandler(usersManager), jwtHandler()).
		WithFilter(corsAuthFilter()).
		Build()
}
