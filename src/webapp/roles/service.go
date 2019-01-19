package roles

import "context"

// Service Inteface for Roles
type Service interface {
	GetAll(ctx context.Context) ([]Role, error)
	GetByID(ctx context.Context, id string) (Role, error)
	Create(ctx context.Context, role Role) (Role, error)
	DeleteByID(ctx context.Context, id string) error
}
