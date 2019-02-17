package users

import (
	"context"
	"errors"
	"fmt"
	net "webapp/core/net/http"
	security "webapp/core/net/http/security"
)

// ServiceImpl Implementation used for the service
type ServiceImpl struct {
	Repository Repository
}

// NewServiceImpl Create a new ServiceImpl
// NOTE: Instead returning a Service interfaces here we are returning a pointer
// 	     since this struct implements two interfaces: Service & security.UserCallback
//       In Go there is no need to implicitly reference the interface we are implementing
//       like other languages.
func NewServiceImpl(r Repository) *ServiceImpl {
	return &ServiceImpl{Repository: r}
}

// GetAll fetches all the users from the repository
func (s *ServiceImpl) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	users, err := s.Repository.FindAll(ctx)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &GetAllResponse{Users: users}, nil
}

// GetByID User by Id
func (s *ServiceImpl) GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error) {
	user, err := s.Repository.FindByID(ctx, req.ID)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	if user == nil {
		return nil, net.ErrNotFound.From(errors.New("User has not been found with id " + req.ID))
	}
	return &GetByIDResponse{User: user}, nil
}

// Create Add user into the repository
func (s *ServiceImpl) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	user, err := s.Repository.Create(ctx, *New(req.Name, req.Email, req.Password))
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	return &CreateResponse{User: user}, nil
}

// DeleteByID user from the repository
func (s *ServiceImpl) DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error) {
	ok, err := s.Repository.DeleteByID(ctx, req.ID)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	if !ok {
		return nil, net.ErrNotFound.From(errors.New("User cannot be deleted with id " + req.ID))
	}
	return &DeleteByIDResponse{}, nil
}

// GetUserByName implements usercallback interface
func (s *ServiceImpl) GetUserByName(ctx context.Context, name string) (*security.UserData, error) {
	user, err := s.Repository.FindByName(ctx, name)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	if user == nil {
		return nil, net.ErrNotFound.From(errors.New("User has not been found with name " + name))
	}
	return &security.UserData{
		ID:       fmt.Sprintf("%v", user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// GetUserByEmail implements usercallback interface
func (s *ServiceImpl) GetUserByEmail(ctx context.Context, email string) (*security.UserData, error) {
	user, err := s.Repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, net.ErrInternalServer.From(err)
	}
	if user == nil {
		return nil, net.ErrNotFound.From(errors.New("User has not been found with email " + email))
	}
	return &security.UserData{
		ID:       fmt.Sprintf("%v", user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
