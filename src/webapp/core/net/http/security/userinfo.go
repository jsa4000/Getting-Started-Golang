package security

import "context"

// UserInfo structure to retrieve (fetch) the user information
type UserInfo struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Email     string                 `json:"email"`
	Password  string                 `json:"password"`
	Roles     []string               `json:"roles"`
	Resources map[string]interface{} `json:"resources"`
}

// UserInfoService Interface
type UserInfoService interface {
	Fetch(ctx context.Context, username string) (*UserInfo, error)
}
