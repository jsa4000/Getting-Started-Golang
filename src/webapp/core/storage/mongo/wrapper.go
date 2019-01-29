package mongo

import (
	"context"
	"fmt"
	"time"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Wrapper structure for MongoDb Client
type Wrapper struct {
	Client *mongo.Client
}

// New returns new mongodb client
func New() *Wrapper {
	return &Wrapper{}
}

// Connect to Mongodb database
func (w *Wrapper) checkConnection(ctx context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := w.Client.Ping(ctx, nil)
	if err != nil {
		log.Error(fmt.Sprintf("Error trying to connect to mongodb. '%s'", err))
		return
	}
	log.Info(fmt.Sprintf("Connected to mongodb at '%s'", w.Client.ConnectionString()))
}

// Connect to Mongodb database
func (w *Wrapper) Connect(ctx context.Context, conn string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, conn)
	if err != nil {
		log.Error(fmt.Sprintf("Error Connecting to mongodb: '%s'", err))
		return err
	}
	w.Client = client
	go w.checkConnection(ctx)
	return nil
}

// Disconnect to Mongodb database
func (w *Wrapper) Disconnect(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	err := w.Client.Disconnect(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("Error disconnecting from mongodb. '%s'", err))
		return err
	}
	log.Info("Connection to MongoDB closed.")
	return nil
}
