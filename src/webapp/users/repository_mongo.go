package users

import (
	"context"

	log "webapp/core/logging"
)

// MongoRepository to implement the Users Repository
type MongoRepository struct {
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository() Repository {

	return &MongoRepository{}
}

// Close gracefully shutdown repository
func (c *MongoRepository) Close() {
	log.Info("Users Repository disconnected")
}

// FindAll fetches all the values form the database
func (c *MongoRepository) FindAll(_ context.Context) ([]User, error) {
	result := make([]User, 0, 0)
	return result, nil
}

// FindByID User by Id
func (c *MongoRepository) FindByID(_ context.Context, id string) (*User, error) {
	return nil, nil

}

// Create Add user into the datbase
func (c *MongoRepository) Create(_ context.Context, user User) (*User, error) {

	return nil, nil
}

// DeleteByID user from the database
func (c *MongoRepository) DeleteByID(_ context.Context, id string) (bool, error) {

	return true, nil
}
