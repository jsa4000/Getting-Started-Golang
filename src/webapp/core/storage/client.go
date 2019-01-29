package storage

import "context"

// Client structure for MongoDb Client
type Client interface {
	Connect(ctx context.Context, conn string) error
	Disconnect(ctx context.Context) error
}
