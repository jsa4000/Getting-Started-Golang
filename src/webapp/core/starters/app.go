package starters

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// App inteface for generic Application
type App interface {
	Startup(ctx context.Context)
	Shutdown(ctx context.Context)
}

// StartApp main function
func StartApp(ctx context.Context, app App) {

	// Initialize the default components: loggin, parser, validation, etc..
	Init()

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	app.Startup(ctx)
	defer app.Shutdown(ctx)

	// Waits until an interrupt is sent from the OS
	<-stop
}
