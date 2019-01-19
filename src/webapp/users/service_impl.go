package users

import "context"

// ServiceImpl Implementation used for the service
type ServiceImpl struct {
	Repository Repository
}

// NewServiceImpl Create a new ServiceImpl
func NewServiceImpl(repo Repository) Service {
	return &ServiceImpl{Repository: repo}
}

// GetAll fetches all the users from the repository
func (s *ServiceImpl) GetAll(ctx context.Context) ([]User, error) {
	users, err := s.Repository.FindAll(ctx)
	return users, err
}

// GetByID User by Id
func (s *ServiceImpl) GetByID(ctx context.Context, id string) (User, error) {
	user, err := s.Repository.FindByID(ctx, id)
	return user, err
}

// Create Add user into the repository
func (s *ServiceImpl) Create(ctx context.Context, user User) (User, error) {
	users, err := s.Repository.Create(ctx, user)
	return users, err
}

// DeleteByID user from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, id string) error {
	return s.Repository.DeleteByID(ctx, id)
}
