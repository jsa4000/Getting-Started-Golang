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

var wrapper *Wrapper

// SetGlobal set global component component
func SetGlobal(w *Wrapper) {
	wrapper = w
}

// New returns new mongodb client
func New() *Wrapper {
	return &Wrapper{
		OnConnect: map[string]OnConnect{},
	}
}

// Client Returns Mongo Client
func Client() *mongo.Client {
	return wrapper.Client
}

// Subscribe to OnConnect event
func Subscribe(id string, f OnConnect) error {
	return wrapper.Subscribe(id, f)
}

// Unsubscribe to OnConnect event
func Unsubscribe(id string) error {
	return wrapper.Unsubscribe(id)
}

// Connect to Mongodb database
func Connect(ctx context.Context, conn string) error {
	return wrapper.Connect(ctx, conn)
}

// Disconnect to Mongodb database
func Disconnect(ctx context.Context) error {
	return wrapper.Disconnect(ctx)
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

// Connected check the connection to Mongodb database
func (w *Wrapper) Connected() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := w.Client.Ping(ctx, nil)
	if err != nil {
		cancel()
		return false
	}
	return true
}

// check the connection to Mongodb database
func (w *Wrapper) check(ctx context.Context) {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	go w.check(ctx)
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
