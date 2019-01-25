package roles

import (
	"context"
	"errors"
)

// ServiceImpl Implementation used for the service
type ServiceImpl struct {
	Repository Repository
}

// NewServiceImpl Create a new ServiceImpl
func NewServiceImpl(r Repository) Service {
	return &ServiceImpl{Repository: r}
}

// GetAll fetches all the roles from the repository
func (s *ServiceImpl) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	roles, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	return &GetAllResponse{Roles: roles}, nil
}

// GetByID Role by Id
func (s *ServiceImpl) GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	role, err := s.Repository.FindByID(ctx, req.ID)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	if role == nil {
		return nil, ErrNotFound.From(errors.New("Role has not been found with id " + req.ID))
	}
	return &GetByIDResponse{Role: role}, nil
}

// Create Add role into the repository
func (s *ServiceImpl) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	role := New(req.Name, req.Name)
	newRole, err := s.Repository.Create(ctx, role)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	return &CreateResponse{Role: newRole}, nil
}

// DeleteByID role from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error) {
	ok, err := s.Repository.DeleteByID(ctx, req.ID)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	if !ok {
		return nil, ErrNotFound.From(errors.New("Role cannot be deleted with id " + req.ID))
	}
	return &DeleteByIDResponse{}, nil
}
