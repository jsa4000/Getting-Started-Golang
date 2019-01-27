package users

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// MockRepository to implement the Users Repository
type MockRepository struct {
	Users map[string]User
}

// NewMockRepository Create a Mock repository
func NewMockRepository() Repository {
	users := make(map[string]User)
	data := []User{
		User{
			ID:       uuid.NewV4().String(),
			Name:     "javier",
			Email:    "javier.golang@example.com",
			Password: "$2a$10$a/pDBjfkq5p9KEMJF6F21.jiMvXozhxcwnASjTfhzTTNghltk2MtG",
		},
		User{
			ID:       uuid.NewV4().String(),
			Name:     "simon",
			Email:    "simon.golang@example.com",
			Password: "$2a$10$a/pDBjfkq5p9KEMJF6F21.jiMvXozhxcwnASjTfhzTTNghltk2MtG",
		},
	}
	for _, value := range data {
		users[value.ID] = value
	}
	return &MockRepository{Users: users}
}

// FindAll fetches all the values form the database
func (c *MockRepository) FindAll(_ context.Context) ([]*User, error) {
	result := make([]*User, 0, len(c.Users))
	for _, val := range c.Users {
		result = append(result, &val)
	}
	return result, nil
}

// FindByID User by Id
func (c *MockRepository) FindByID(_ context.Context, id string) (*User, error) {
	user, ok := c.Users[id]
	if !ok {
		return nil, nil
	}
	return &user, nil
}

// Create Add user into the datbase
func (c *MockRepository) Create(_ context.Context, user User) (*User, error) {
	user = User{
		ID:       uuid.NewV4().String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	c.Users[user.ID] = user
	return &user, nil
}

// DeleteByID user from the database
func (c *MockRepository) DeleteByID(_ context.Context, id string) (bool, error) {
	_, ok := c.Users[id]
	if !ok {
		return false, nil
	}
	delete(c.Users, id)
	return true, nil
}
