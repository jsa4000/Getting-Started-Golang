package security

// UserData structure to identify the user
type UserData struct {
	ID       string
	Name     string
	Email    string
	Password string
	Roles    []string
}

// TokenData data structure for token generation (claims)
type TokenData map[string]interface{}
