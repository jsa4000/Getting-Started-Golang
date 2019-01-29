package roles

import (
	"context"

	mongow "webapp/core/storage/mongo"

	"github.com/mongodb/mongo-go-driver/mongo"
)

const timeout = 10
const database = "webapp"
const collection = "roles"

// MongoRepository to implement the Users Repository
type MongoRepository struct {
	Wrapper    *mongow.Wrapper
	Collection *mongo.Collection
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository(wrapper *mongow.Wrapper) Repository {
	result := &MongoRepository{
		Wrapper:    wrapper,
		Collection: wrapper.Client.Database(database).Collection(collection),
	}
	result.CreateIndexes(context.Background())
	return result
}

// CreateIndexes create index for the collection
func (c *MongoRepository) CreateIndexes(ctx context.Context) {

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
