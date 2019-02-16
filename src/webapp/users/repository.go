package users

import "context"

// Repository Interface for Users
type Repository interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByName(ctx context.Context, name string) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	DeleteByID(ctx context.Context, id string) (bool, error)
}
