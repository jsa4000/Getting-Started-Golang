package main

import (
	"context"
	"webapp/core/starters"
)

func main() {
	// Start the Application
	starters.StartApp(context.Background(), &App{})
}
