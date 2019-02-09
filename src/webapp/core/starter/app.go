package starter

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "webapp/core/logging"
)

// App interface for generic Application
type App interface {
	Startup(ctx context.Context)
	Shutdown(ctx context.Context)
}

// StartApp main function
func StartApp(ctx context.Context, app App) {
	start := time.Now()

	// Initialize the default components: config, logging, validation, etc..
	Init(ctx)

	// Load App Config
	c := LoadConfig()

	log.Infof("Starting Application %s", c.Name)

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	app.Startup(ctx)

	log.Debugf("%s Started in %d ns", c.Name, time.Since(start).Nanoseconds())

	// Waits until an interrupt is sent from the OS
	<-stop

	end := time.Now()

	// Create a new context to shutdown the application
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	// Shutdown the app
	app.Shutdown(ctx)

	// Shutdown the components registered previously
	Shutdown(ctx)

	log.Debugf("%s shutdown in %d ns", c.Name, time.Since(end).Nanoseconds())

}
