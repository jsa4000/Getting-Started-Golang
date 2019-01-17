package users

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

// MockRepository to implement the Users Repository
type MockRepository struct {
	Users map[string]User
}

// NewMockRepository Create a Mock repository
func NewMockRepository() Repository {
	users := make(map[string]User)
	data := []User{
		New("javier", "javier.golang@example.com", "password"),
		New("simon", "simon.golang@example.com", "password"),
	}
	for _, value := range data {
		users[value.ID] = value
	}
	return &MockRepository{Users: users}
}

// Close gracefully shutdown repository
func (c *MockRepository) Close() {
	log.Info("Users Repository disconnected")
}

// FindAll fetches all the values form the database
func (c *MockRepository) FindAll(_ context.Context) ([]User, error) {
	result := make([]User, 0, len(c.Users))
	for _, val := range c.Users {
		result = append(result, val)
	}
	return result, nil
}

// FindByID User by Id
func (c *MockRepository) FindByID(_ context.Context, id string) (User, error) {
	_, ok := c.Users[id]
	if !ok {
		return User{}, errors.New("User has not been found with id " + id)
	}
	return c.Users[id], nil
}

// Create Add user into the datbase
func (c *MockRepository) Create(_ context.Context, user User) (User, error) {
	user = New(user.Name, user.Email, user.Password)
	c.Users[user.ID] = user
	return user, nil
}

// DeleteByID user from the database
func (c *MockRepository) DeleteByID(_ context.Context, id string) error {
	_, ok := c.Users[id]
	if !ok {
		return errors.New("User has not been found with id " + id)
	}
	delete(c.Users, id)
	return nil
}
