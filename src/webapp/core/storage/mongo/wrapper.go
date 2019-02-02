package mongo

import (
	"context"
	"fmt"
	"time"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// OnConnect Handler that will be called every time is connects
type OnConnect func()

// Wrapper structure for MongoDb Client
type Wrapper struct {
	Client    *mongo.Client
	OnConnect map[string]OnConnect
}

// New returns new mongodb client
func New() *Wrapper {
	return &Wrapper{
		OnConnect: map[string]OnConnect{},
	}
}

func (w *Wrapper) onConnectEvent() {
	for _, sub := range w.OnConnect {
		go sub()
	}
}

// Subscribe to OnConnect event
func (w *Wrapper) Subscribe(id string, f OnConnect) error {
	_, exist := w.OnConnect[id]
	if exist {
		return fmt.Errorf("Key '%s' already subscribed to MongoDb OnConnect event", id)
	}
	w.OnConnect[id] = f
	return nil
}

// Unsubscribe to OnConnect event
func (w *Wrapper) Unsubscribe(id string) error {
	_, ok := w.OnConnect[id]
	if !ok {
		return fmt.Errorf("Key '%s' has not been subscribed to MongoDb OnConnect event", id)
	}
	delete(w.OnConnect, id)
	return nil
}

// Connect to Mongodb database
func (w *Wrapper) checkConnection(ctx context.Context) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := w.Client.Ping(ctx, nil)
		if err != nil {
			log.Error(fmt.Sprintf("Error trying to connect to mongodb. '%s'", err))
			cancel()
			continue
		}
		log.Info(fmt.Sprintf("Connected to mongodb at '%s'", w.Client.ConnectionString()))
		w.onConnectEvent()
		break
	}

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
