package users

import (
	"context"
)

// Repository Interface for Users
type Repository interface {
	Close()
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) (User, error)
	DeleteByID(ctx context.Context, id string) error
}
