package main

import (
	"context"
	"webapp/core/starter"
	"webapp/server"
)

func main() {
	// Start the Application
	starter.StartApp(context.Background(), &server.App{})
}
