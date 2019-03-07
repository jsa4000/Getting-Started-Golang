package users

import (
	pcrypt "webapp/core/crypto/password"
)

// User struct to define an User
type User struct {
	ID       interface{} `json:"id" bson:"_id,omitempty"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Roles    []string    `json:"roles"`
}

// New Creates a new User
func New(name, email, password string, roles []string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: pcrypt.New(password),
		Roles:    roles,
	}
}
