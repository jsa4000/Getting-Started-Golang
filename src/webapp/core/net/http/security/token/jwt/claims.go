package jwt

import "webapp/core/net/http/security"

// Claims data structure for token enhacements and generation
type Claims map[string]interface{}

// ClaimsEnhancer Interface
type ClaimsEnhancer interface {
	Write(c Claims, u *security.UserInfo)
}
