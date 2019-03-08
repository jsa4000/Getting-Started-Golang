package users

import (
	"context"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"
	mongow "webapp/core/store/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

const (
	timeout      = 10
	repositoryID = "usersrepository"
)

// MongoConfig main app configuration
type MongoConfig struct {
	Database   string `config:"repository.mongodb.users.database"`
	Collection string `config:"repository.mongodb.users.collection"`
}

// MongoRepository to implement the Users Repository
type MongoRepository struct {
	Collection *mongo.Collection
}

// NewMongoRepository Create a Mock repository
func NewMongoRepository() Repository {
	c := MongoConfig{}
	config.ReadFields(&c)
	result := &MongoRepository{
		Collection: mongow.Client().Database(c.Database).Collection(c.Collection),
	}
	err := mongow.Subscribe(repositoryID, result.onConnect)
	if err != nil {
		log.Error(err)
	}
	return result
}

func (c *MongoRepository) onConnect() {
	log.Debug("Users Repository received OnConnect event from Mongodb")
	c.CreateIndexes(context.Background())
}

// CreateIndexes create index for the collection
func (c *MongoRepository) CreateIndexes(ctx context.Context) {
	mongow.CreateIndex(ctx, c.Collection, []mongo.IndexModel{
		mongow.UniqueIndex("name", true, false),
		mongow.UniqueIndex("email", true, true),
	}...)
	c.boostrap()
}

// FindAll fetches all the values form the database
func (c *MongoRepository) FindAll(ctx context.Context) ([]*User, error) {
	users := []*User{}
	options := options.Find()
	options.SetLimit(100)
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	cur, err := c.Collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result User
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
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &result, nil
}

// FindByEmail User by Id
func (c *MongoRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var result User
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	idDoc := bson.M{"email": email}
	err := c.Collection.FindOne(ctx, idDoc).Decode(&result)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return normalize(&result), nil
}

// FindByName User by Id
func (c *MongoRepository) FindByName(ctx context.Context, name string) (*User, error) {
	var result User
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()
	idDoc := bson.M{"name": name}
	err := c.Collection.FindOne(ctx, idDoc).Decode(&result)
	if err != nil && err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return normalize(&result), nil
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

func normalize(user *User) *User {
	id, _ := user.ID.(primitive.ObjectID)
	user.ID = id.Hex()
	return user
}

func (c *MongoRepository) boostrap() {
	objects, err := BootstrapData()
	if err != nil {
		log.Error(err)
	}
	oo := make([]interface{}, 0, len(objects))
	for _, o := range objects {
		oo = append(oo, o)
	}
	ordered := false
	result, err := c.Collection.InsertMany(context.Background(), oo, &options.InsertManyOptions{
		Ordered: &ordered,
	})
	insertedIDs := len(result.InsertedIDs)
	if err != nil {
		if errors, ok := err.(mongo.BulkWriteException); ok {
			errorIDs := len(errors.WriteErrors)
			if errorIDs != insertedIDs {
				log.Debugf("Inserted %d new default %s", insertedIDs-errorIDs, c.Collection.Name())
			}
		}
		return
	}
	log.Debugf("Inserted %d new default %s", insertedIDs, c.Collection.Name())
}
