package roles

import (
	"context"

	log "webapp/core/logging"
)

// MongoRepository to implement the Roles Repository
type MongoRepository struct {
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository() Repository {

	return &MongoRepository{}
}

// Close gracefully shutdown repository
func (c *MongoRepository) Close() {
	log.Info("Roles Repository disconnected")
}

// FindAll fetches all the values form the database
func (c *MongoRepository) FindAll(_ context.Context) ([]Role, error) {
	result := make([]Role, 0, 0)
	return result, nil
}

// FindByID Role by Id
func (c *MongoRepository) FindByID(_ context.Context, id string) (*Role, error) {
	return nil, nil

}

// Create Add Role into the datbase
func (c *MongoRepository) Create(_ context.Context, Role Role) (*Role, error) {

	return nil, nil
}

// DeleteByID Role from the database
func (c *MongoRepository) DeleteByID(_ context.Context, id string) (bool, error) {

	return true, nil
}
