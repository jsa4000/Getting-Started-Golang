package security

// AuthenticationManager struct to handle basic authentication
type AuthenticationManager struct {
	*Config
	providers map[string]UserInfoService
}
