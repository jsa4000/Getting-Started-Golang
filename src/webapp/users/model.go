package users

import "golang.org/x/crypto/bcrypt"

// User struct to define an User
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Bcrypt. Online Converter: https://8gwifi.org/bccrypt.jsp
	Password string `json:"password"`
}

// New Creates a new User
func New(name, email, password string) *User {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return &User{
		Name:     name,
		Email:    email,
		Password: string(p),
	}
}
