package roles

import (
	"context"
)

// Repository Inteface for Roles
type Repository interface {
	Close()
	FindAll(ctx context.Context) ([]Role, error)
	FindByID(ctx context.Context, id string) (Role, error)
	Create(ctx context.Context, role Role) (Role, error)
	DeleteByID(ctx context.Context, id string) error
}