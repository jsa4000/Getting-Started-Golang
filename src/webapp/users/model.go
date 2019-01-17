package users

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct to define an User
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// New creates new instance
func New(name string, email string, password string) User {
	cyptPass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return User{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Email:    email,
		Password: string(cyptPass),
	}
}
