package users

import "context"

// Service Inteface for Users
type Service interface {
	GetAll(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, user User) (User, error)
	RemoveByID(ctx context.Context, id string) error
}
