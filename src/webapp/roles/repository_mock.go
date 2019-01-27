package roles

import (
	"context"
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

// FindAll fetches all the values form the database
func (c *MockRepository) FindAll(_ context.Context) ([]Role, error) {
	result := make([]Role, 0, len(c.Roles))
	for _, val := range c.Roles {
		result = append(result, val)
	}
	return result, nil
}

// FindByID Role by Id
func (c *MockRepository) FindByID(_ context.Context, id string) (*Role, error) {
	role, ok := c.Roles[id]
	if !ok {
		return nil, nil
	}
	return &role, nil
}

// Create Add role into the datbase
func (c *MockRepository) Create(_ context.Context, role Role) (*Role, error) {
	role = New(role.Name, role.Name)
	c.Roles[role.ID] = role
	return &role, nil
}

// DeleteByID role from the database
func (c *MockRepository) DeleteByID(_ context.Context, id string) (bool, error) {
	_, ok := c.Roles[id]
	if !ok {
		return false, nil
	}
	delete(c.Roles, id)
	return true, nil
}
