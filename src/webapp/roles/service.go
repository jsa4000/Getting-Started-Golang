package roles

import "context"

// Service Inteface for Roles
type Service interface {
	GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error)
	GetByID(ctx context.Context, req *GetByIDRequest) (*GetByIDResponse, error)
	Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error)
	DeleteByID(ctx context.Context, req *DeleteByIDRequest) (*DeleteByIDResponse, error)
}