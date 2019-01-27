package users

import (
	"context"
	"time"

	"webapp/core/database/mongo"

	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	driver "github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

const timeout = 10
const database = "webapp"
const collection = "users"

// MongoRepository to implement the Users Repository
type MongoRepository struct {
	Client     *mongo.Client
	Collection *driver.Collection
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository(client *mongo.Client) Repository {
	result := &MongoRepository{
		Client:     client,
		Collection: client.Db.Database(database).Collection(collection),
	}
	result.CreateIndexes(context.Background())
	return result
}

// CreateIndexes create index for the collection
func (c *MongoRepository) CreateIndexes(ctx context.Context) {
	// Create ascending index (1) for email Set index as unique index
	indexes := []driver.IndexModel{
		mongo.CreateIndexModel("name", true, false),
		mongo.CreateIndexModel("email", true, true),
	}
	mongo.CreateIndex(ctx, c.Collection, indexes...)
}

// FindAll fetches all the values form the database
func (c *MongoRepository) FindAll(ctx context.Context) ([]*User, error) {
	options := options.Find()
	options.SetLimit(100)
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	cur, err := c.Collection.Find(ctx, bson.M{}, options)
	users := []*User{}
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)
	var result User

	for cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			log.Error(err)
			continue
		}
		users = append(users, &result)
	}
	if err := cur.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// FindByID User by Id
func (c *MongoRepository) FindByID(ctx context.Context, id string) (*User, error) {
	var result User
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	hexID, _ := primitive.ObjectIDFromHex(id)
	idDoc := bson.M{"_id": hexID}
	err := c.Collection.FindOne(ctx, idDoc).Decode(&result)
	if err != nil && err == driver.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create Add user into the datbase
func (c *MongoRepository) Create(ctx context.Context, user User) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	result, err := c.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	user.ID = id.Hex()
	return &user, nil
}

// DeleteByID user from the database
func (c *MongoRepository) DeleteByID(ctx context.Context, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	hexID, _ := primitive.ObjectIDFromHex(id)
	idDoc := bson.M{"_id": hexID}
	result, err := c.Collection.DeleteOne(ctx, idDoc)
	if err != nil {
		return false, err
	}
	if result.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}