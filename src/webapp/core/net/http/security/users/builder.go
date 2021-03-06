package users

import (
	"webapp/core/net/http/security"
)

// ManagerBuilder struct to handle basic authentication
type ManagerBuilder struct {
	*Manager
	UsersBuilder *NestedUsersBuilder
}

// NewManagerBuilder Create a new UsersBuilder
func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{
		Manager: &Manager{
			services: make([]security.UserInfoService, 0),
		},
	}
}

// NestedUsersBuilder build struct
type NestedUsersBuilder struct {
	*security.UsersBuilder
	parent *ManagerBuilder
}

// NewUsersBuilder Create a new UsersBuilder
func newNestedUsersBuilder(parent *ManagerBuilder) *NestedUsersBuilder {
	return &NestedUsersBuilder{
		security.NewUsersBuilder(),
		parent,
	}
}

// WithUserService Add user service
func (b *ManagerBuilder) WithUserService(s security.UserInfoService) *ManagerBuilder {
	b.services = append(b.services, s)
	return b
}

// WithInMemoryUsers set the interface to use for fetching user info
func (b *ManagerBuilder) WithInMemoryUsers() *NestedUsersBuilder {
	if b.UsersBuilder == nil {
		b.UsersBuilder = newNestedUsersBuilder(b)
	}
	return b.UsersBuilder
}

// Build Add user service
func (b *ManagerBuilder) Build() *Manager {
	b.services = append([]security.UserInfoService{b.UsersBuilder.Build()}, b.services...)
	return b.Manager
}

// WithUser set the interface to use for fetching user info
func (c *NestedUsersBuilder) WithUser(name string) *NestedUsersBuilder {
	c.UsersBuilder.WithUser(name)
	return c
}

// WithPassword set the interface to use for fetching user info
func (c *NestedUsersBuilder) WithPassword(password string) *NestedUsersBuilder {
	c.UsersBuilder.WithPassword(password)
	return c
}

// WithRole set the interface to use for fetching user info
func (c *NestedUsersBuilder) WithRole(roles ...string) *NestedUsersBuilder {
	c.UsersBuilder.WithRole(roles...)
	return c
}

// And set the interface to use for fetching user info
func (c *NestedUsersBuilder) And() *ManagerBuilder {
	return c.parent
}
