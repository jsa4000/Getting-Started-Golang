package password

import (
	"golang.org/x/crypto/bcrypt"
)

// New crtpy current password into standard from core
func New(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(p)
}

// Compare hash with current password
func Compare(hash, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}
