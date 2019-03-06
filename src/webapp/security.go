package main

import (
	"webapp/core/net/http"
	"webapp/core/net/http/security"
	"webapp/core/net/http/security/access"
	"webapp/core/net/http/security/basic"
	"webapp/core/net/http/security/oauth2"
	"webapp/core/net/http/security/open"
	"webapp/core/net/http/security/roles"
	"webapp/core/net/http/security/token/jwt"
	"webapp/core/net/http/security/users"
)

func jwtHandler() security.Handler {
	return jwt.NewBuilder().
		WithTargets().
		WithURL("/*").
		And().
		WithPriority(2).
		Build()
}

func basicHandler(usersManager *users.Manager) security.Handler {
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

func openHandler() security.Handler {
	return open.NewBuilder().
		WithTargets().
		//WithURL("/debug/pprof/").
		WithURL("/oauth/*").
		And().
		WithPriority(0).
		Build()
}

func rolesHandler() security.Handler {
	return roles.NewBuilder().
		WithTargets().
		WithURL("/*").
		And().
		Build()
}

func corsAuthFilter() security.Handler {
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
		WithUser("user-trusted").WithPassword("mypassword$").WithRole("ADMIN", "WRITE", "READ", "TRUSTED").
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
		WithAuthorization(openHandler(), basicHandler(usersManager), jwtHandler()).
		WithFilter(rolesHandler(), corsAuthFilter()).
		Build()
}
