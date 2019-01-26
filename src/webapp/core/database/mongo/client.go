package mongo

import (
	"context"
	"fmt"
	"time"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/mongo"
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
