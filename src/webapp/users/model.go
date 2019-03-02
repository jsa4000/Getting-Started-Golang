package users

import (
	pcrypt "webapp/core/crypto/password"
)

// User struct to define an User
type User struct {
	ID    interface{} `json:"id" bson:"_id,omitempty"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	// Bcrypt. Online Converter: https://8gwifi.org/bccrypt.jsp
	Password string `json:"password"`
}

// New Creates a new User
func New(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: pcrypt.New(password),
	}
}
