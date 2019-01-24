package users

import (
	"context"
)

// Service Interface for Users
type Service interface {
	GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error)
	GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error)
}
