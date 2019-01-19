package roles

import "context"

// ServiceImpl Implementation used for the service
type ServiceImpl struct {
	Repository Repository
}

// NewServiceImpl Create a new ServiceImpl
func NewServiceImpl(repo Repository) Service {
	return &ServiceImpl{Repository: repo}
}

// GetAll fetches all the roles from the repository
func (s *ServiceImpl) GetAll(ctx context.Context) ([]Role, error) {
	roles, err := s.Repository.FindAll(ctx)
	return roles, err
}

// GetByID Role by Id
func (s *ServiceImpl) GetByID(ctx context.Context, id string) (Role, error) {
	role, err := s.Repository.FindByID(ctx, id)
	return role, err
}

// Create Add role into the repository
func (s *ServiceImpl) Create(ctx context.Context, role Role) (Role, error) {
	roles, err := s.Repository.Create(ctx, role)
	return roles, err
}

// DeleteByID role from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, id string) error {
	return s.Repository.DeleteByID(ctx, id)
}
