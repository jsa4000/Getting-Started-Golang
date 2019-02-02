package roles

import (
	"context"
	"time"

	log "webapp/core/logging"
	mongow "webapp/core/storage/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

const timeout = 10
const database = "webapp"
const collection = "roles"

// MongoRepository to implement the Roles Repository
type MongoRepository struct {
	Wrapper    *mongow.Wrapper
	Collection *mongo.Collection
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository(wrapper *mongow.Wrapper) Repository {
	return &MongoRepository{
		Wrapper:    wrapper,
		Collection: wrapper.Client.Database(database).Collection(collection),
	}
}

// FindAll fetches all the values form the database
func (c *MongoRepository) FindAll(ctx context.Context) ([]*Role, error) {
	Roles := []*Role{}
	options := options.Find()
	options.SetLimit(100)
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	cur, err := c.Collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return Roles, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result Role
		err := cur.Decode(&result)
		if err != nil {
			log.Error(err)
			continue
		}
		Roles = append(Roles, &result)
	}
	if err := cur.Err(); err != nil {
		return Roles, err
	}
	return Roles, nil
}

// FindByID Role by Id
func (c *MongoRepository) FindByID(ctx context.Context, id string) (*Role, error) {
	var result Role
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	idDoc := bson.M{"_id": id}
	err := c.Collection.FindOne(ctx, idDoc).Decode(&result)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create Add Role into the datbase
func (c *MongoRepository) Create(ctx context.Context, Role Role) (*Role, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	_, err := c.Collection.InsertOne(ctx, Role)
	if err != nil {
		return nil, err
	}
	return &Role, nil
}

// DeleteByID Role from the database
func (c *MongoRepository) DeleteByID(ctx context.Context, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	idDoc := bson.M{"_id": id}
	result, err := c.Collection.DeleteOne(ctx, idDoc)
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}
