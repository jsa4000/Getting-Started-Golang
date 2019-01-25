package users

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

// GetAll fetches all the users from the repository
func (s *ServiceImpl) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	users, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	return &GetAllResponse{Users: users}, nil
}

// GetByID User by Id
func (s *ServiceImpl) GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	user, err := s.Repository.FindByID(ctx, req.ID)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	if user == nil {
		return nil, ErrNotFound.From(errors.New("User has not been found with id " + req.ID))
	}
	return &GetByIDResponse{User: user}, nil
}

// Create Add user into the repository
func (s *ServiceImpl) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	user, err := s.Repository.Create(ctx, *New(req.Name, req.Email, req.Password))
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	return &CreateResponse{User: user}, nil
}

// DeleteByID user from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error) {
	ok, err := s.Repository.DeleteByID(ctx, req.ID)
	if err != nil {
		return nil, ErrInternalServer.From(err)
	}
	if !ok {
		return nil, ErrNotFound.From(errors.New("User cannot be deleted with id " + req.ID))
	}
	return &DeleteByIDResponse{}, nil
}
