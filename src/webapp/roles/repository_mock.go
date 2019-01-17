package roles

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

// MockRepository to implement the Roles Repository
type MockRepository struct {
	Roles map[string]Role
}

// NewMockRepository Create a Mock repository
func NewMockRepository() Repository {
	roles := make(map[string]Role)
	data := []Role{
		New("ROLE_ADMIN", "ROLE_ADMIN"),
		New("ROLE_USER", "ROLE_USER"),
	}
	for _, value := range data {
		roles[value.ID] = value
	}
	return &MockRepository{Roles: roles}
}

// Close gracefully shutdown repository
func (c *MockRepository) Close() {
	log.Info("Roles Repository disconnected")
}

// FindAll fetches all the values form the database
func (c *MockRepository) FindAll(_ context.Context) ([]Role, error) {
	result := make([]Role, 0, len(c.Roles))
	for _, val := range c.Roles {
		result = append(result, val)
	}
	return result, nil
}

// FindByID Role by Id
func (c *MockRepository) FindByID(_ context.Context, id string) (Role, error) {
	_, ok := c.Roles[id]
	if !ok {
		return Role{}, errors.New("Role has not been found with id " + id)
	}
	return c.Roles[id], nil
}

// Create Add role into the datbase
func (c *MockRepository) Create(_ context.Context, role Role) (Role, error) {
	role = New(role.Name, role.Name)
	c.Roles[role.ID] = role
	return role, nil
}

// DeleteByID role from the database
func (c *MockRepository) DeleteByID(_ context.Context, id string) error {
	_, ok := c.Roles[id]
	if !ok {
		return errors.New("Role has not been found with id " + id)
	}
	delete(c.Roles, id)
	return nil
}
