package users

import (
	"context"
	"time"

	"webapp/core/database/mongo"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	driver "github.com/mongodb/mongo-go-driver/mongo"
)

const database = "webapp"
const collection = "users"

// MongoRepository to implement the Users Repository
type MongoRepository struct {
	Client     *mongo.Client
	Database   *driver.Database
	Collection *driver.Collection
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository(client *mongo.Client) Repository {
	return &MongoRepository{
		Client:     client,
		Database:   client.Db.Database(database),
		Collection: client.Db.Database(database).Collection(collection),
	}
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
	var result User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//idDoc := bson.NewDocument(bson.EC.ObjectID("_id", "5c4ca5e1fa98e587c68295fc"))
	idDoc := bson.D{{"name", "user_test"}}
	err := c.Collection.FindOne(ctx, idDoc).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Create Add user into the datbase
func (c *MongoRepository) Create(_ context.Context, user User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := c.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	user = User{
		ID:       id.Hex(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return &user, nil
}

// DeleteByID user from the database
func (c *MongoRepository) DeleteByID(_ context.Context, id string) (bool, error) {

	return true, nil
}
