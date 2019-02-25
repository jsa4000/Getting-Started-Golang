package token

// UserData structure to retrieve (fetch) the user information
type UserData struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Password  string                 `json:"password"`
	Roles     []string               `json:"roles"`
	Resources map[string]interface{} `json:"resources"`
}

// Claims data structure for token enhacements and generation
type Claims map[string]interface{}
