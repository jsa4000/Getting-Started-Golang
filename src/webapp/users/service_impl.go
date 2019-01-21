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
func (s *ServiceImpl) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	users, err := s.Repository.FindAll(ctx)
	return &GetAllResponse{Users: users}, err
}

// GetByID User by Id
func (s *ServiceImpl) GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	user, err := s.Repository.FindByID(ctx, req.ID)
	return &GetByIDResponse{User: user}, err
}

// Create Add user into the repository
func (s *ServiceImpl) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	user := New(req.Name, req.Email, req.Password)
	newUSer, err := s.Repository.Create(ctx, user)
	return &CreateResponse{User : newUSer}, err
}

// DeleteByID user from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error) {
	err := s.Repository.DeleteByID(ctx, req.ID)
	return &DeleteByIDResponse{}, err
}
