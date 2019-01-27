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

const database = "webapp"
const collection = "users"

// UserM struct to define an User
type UserM struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Email    string              `json:"email" bson:"email"`
	Password string              `json:"password" bson:"password"`
}

// Convert Returns a user from User
func (u *UserM) Convert() *User {
	return &User{
		ID:       u.ID.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

// From Returns a user from User
func From(user User) *UserM {
	return &UserM{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

// MongoRepository to implement the Users Repository
type MongoRepository struct {
	Client     *mongo.Client
	Collection *driver.Collection
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository(client *mongo.Client) Repository {
	return &MongoRepository{
		Client:     client,
		Collection: client.Db.Database(database).Collection(collection),
	}
}

// FindAll fetches all the values form the database
func (c *MongoRepository) FindAll(ctx context.Context) ([]*User, error) {
	options := options.Find()
	options.SetLimit(100)
	cur, err := c.Collection.Find(ctx, bson.M{}, options)
	users := []*User{}
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)
	var result UserM

	for cur.Next(ctx) {
		err := cur.Decode(&result)
		if err != nil {
			log.Error(err)
			continue
		}
		users = append(users, result.Convert())
	}
	if err := cur.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// FindByID User by Id
func (c *MongoRepository) FindByID(ctx context.Context, id string) (*User, error) {
	var result UserM
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	hexID, _ := primitive.ObjectIDFromHex(id)
	idDoc := bson.M{"_id": hexID}
	err := c.Collection.FindOne(ctx, idDoc).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Convert(), nil
}

// Create Add user into the datbase
func (c *MongoRepository) Create(ctx context.Context, user User) (*User, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	result, err := c.Collection.InsertOne(ctx, From(user))
	if err != nil {
		return nil, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	user.ID = id.Hex()
	return &user, nil
}

// DeleteByID user from the database
func (c *MongoRepository) DeleteByID(ctx context.Context, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
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
