package access

const (
	// AllowAnyOrigin allows control-access origin for all domains
	AllowAnyOrigin string = "*"

	// HeaderAccessControlAllowOrigin to allow specifi origin
	HeaderAccessControlAllowOrigin = "Access-Control-Allow-Origin"

	// HeaderAccessControlAllowMethods to allow specific methods
	HeaderAccessControlAllowMethods = "Access-Control-Allow-Methods"

	// HeaderAccessControlAllowHeaders to allow specific headers
	HeaderAccessControlAllowHeaders = "Access-Control-Allow-Headers"

	// HeaderAccessControlAllowCredentials to allow credentials
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
)
