package jwt

import (
	"webapp/core/net/http/security"
)

const (
	// AuthKey Key to get data from context in basicAuth
	AuthKey security.ContextKey = "jwt-auth-key"

	bearerPreffix = "Bearer "
	authHeader    = "Authorization"

	jsontokenIDfield    = "jti"
	issuerField         = "iss"
	subjectField        = "sub"
	emailField          = "email"
	userNameField       = "name"
	scopesField         = "scopes"
	rolesField          = "roles"
	expirationDateField = "exp"
	issuedAtField       = "iat"
)
