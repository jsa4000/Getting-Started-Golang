package main

import (
	"context"
	"webapp/core/starters"
)

func main() {
	// Start the server
	starters.StartApp(context.Background(), &App{})
}
