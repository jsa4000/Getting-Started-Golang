package mongo

import (
	"context"
	"fmt"
	"time"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

// Client structure for MongoDb Client
type Client struct {
	Db *mongo.Client
}

// New returns new mongodb client
func New() *Client {
	return &Client{}
}

// Connect to Mongodb database
func (c *Client) Connect(conn string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, conn)
	if err != nil {
		log.Error(fmt.Sprintf("Error Connecting to mongodb: %s", err))
		return err
	}
	c.Db = client
	// Check the connection
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Error(fmt.Sprintf("Error trying to connect to mongodb. %s", err))
		}
		log.Info(fmt.Sprintf("Connected to mongodb at %s", conn))
	}()
	return nil
}

// Disconnect to Mongodb database
func (c *Client) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := c.Db.Disconnect(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("Error disconnecting from mongodb. %s", err))
		return err
	}
	log.Info("Connection to MongoDB closed.")
	return nil
}

// CreateIndexModel fetches all the values form the database
func CreateIndexModel(name string, asc, unique bool) mongo.IndexModel {
	// Set index options (unique)
	options := options.Index()
	options.SetUnique(unique)
	// Set asc desc index
	var value int32 = 1
	if !asc {
		value = -1
	}
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: name, Value: bsonx.Int32(value)}}
	index.Options = options
	return index
}

// CreateIndex fetches all the values form the database
func CreateIndex(ctx context.Context, c *mongo.Collection, indexes ...mongo.IndexModel) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var err error
	if len(indexes) == 1 {
		log.Debug("Creating index for collection ", c.Name)
		_, err = c.Indexes().CreateOne(ctx, indexes[0])
	} else {
		log.Debug("Creating multiple indexes for collection ", c.Name)
		_, err = c.Indexes().CreateMany(ctx, indexes)
	}
	if err != nil {
		log.Error(err)
	}
	return err
}
