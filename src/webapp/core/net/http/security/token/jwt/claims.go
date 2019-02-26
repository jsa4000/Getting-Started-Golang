package jwt

// Claims data structure for token enhacements and generation
type Claims map[string]interface{}

// ClaimsEnhancer Interface
type ClaimsEnhancer interface {
	Write(t *Claims)
}
