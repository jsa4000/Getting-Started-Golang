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
func (s *ServiceImpl) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	roles, err := s.Repository.FindAll(ctx)
	return &GetAllResponse{Roles: roles}, err
}

// GetByID Role by Id
func (s *ServiceImpl) GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	role, err := s.Repository.FindByID(ctx, req.ID)
	return &GetByIDResponse{Role: role}, err
}

// Create Add role into the repository
func (s *ServiceImpl) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	role := New(req.Name, req.Name)
	newRole, err := s.Repository.Create(ctx, role)
	return &CreateResponse{Role : newRole}, err
}

// DeleteByID role from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error) {
	err := s.Repository.DeleteByID(ctx, req.ID)
	return &DeleteByIDResponse{}, err
}
