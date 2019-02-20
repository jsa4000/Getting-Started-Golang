package main

import (
	"context"
	"webapp/core/starter"
)

func main() {
	// Start the Application
	starter.StartApp(context.Background(), &App{})
}
