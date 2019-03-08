package roles

import (
	"context"
	decode "encoding/json"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"
	mongow "webapp/core/store/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

const (
	timeout      = 10
	repositoryID = "rolesrepository"
)

// MongoConfig for mongoDB Repository
type MongoConfig struct {
	Database   string `config:"repository.mongodb.roles.database"`
	Collection string `config:"repository.mongodb.roles.collection"`
}

// MongoRepository to implement the Roles Repository
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
	log.Debug("Roles Repository received OnConnect event from Mongodb")
	c.boostrap()
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

func (c *MongoRepository) boostrap() error {

	data := `[
{
	"id": "ADMIN",
	"name": "ADMIN"
},
{
	"id": "WRITE",
	"name": "WRITE"
},
{
	"id": "READ",
	"name": "READ"
}
]`
	var roles []Role
	err := decode.Unmarshal([]byte(data), &roles)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, role := range roles {
		idDoc := bson.M{"_id": role.ID}
		if result := c.Collection.FindOne(context.Background(), idDoc); result.Err() == nil {
			_, err := c.Collection.InsertOne(context.Background(), role)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("Inserted default role: id: %s, name: %s ", role.ID, role.Name)
		}
	}
	return nil
}
