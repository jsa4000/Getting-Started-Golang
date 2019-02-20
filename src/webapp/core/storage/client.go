package storage

import "context"

// Client structure for Storage Client
type Client interface {
	Connect(ctx context.Context, conn string) error
	Disconnect(ctx context.Context) error
}

// OnConnect Handler that will be called every time is connects
type OnConnect func()
