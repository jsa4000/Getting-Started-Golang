package users

import (
	"webapp/core/net/http/security"
)

// Manager struct to handle basic authentication
type Manager struct {
	services []security.UserInfoService
}
