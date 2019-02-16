package security

import (
	"context"
	"errors"
	net "webapp/core/net/http"
	security "webapp/core/net/http/security"
	users "webapp/users"
)

// CredentialService Implementation used for the service
type CredentialService struct {
	Repository users.Repository
}

// NewCredentialService Create a new SecurityService
func NewCredentialService(r users.Repository) security.UserCallback {
	return &CredentialService{Repository: r}
}

// GetUserByName or security
func (s *CredentialService) GetUserByName(ctx context.Context, name string) (*security.UserData, error) {
	user, err := s.Repository.FindByName(ctx, name)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	if user == nil {
		return nil, net.ErrNotFound.From(errors.New("User has not been found with name " + name))
	}
	return &security.UserData{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// GetUserByEmail or security
func (s *CredentialService) GetUserByEmail(ctx context.Context, email string) (*security.UserData, error) {
	user, err := s.Repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	if user == nil {
		return nil, net.ErrNotFound.From(errors.New("User has not been found with email " + email))
	}
	return &security.UserData{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
