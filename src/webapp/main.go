package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/config"
	"webapp/roles"
	"webapp/server"
	"webapp/users"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Set the log formatter
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Create Parser (Configuration)
	parser := config.NewViperParserFromFile("webapp.yaml", ".")

	// Read the Configuration
	appconfig := server.NewAppConfig(parser)

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	log.Infof("Starting %s Server...", appconfig.Name)

	// Create Repositories
	usersRepo := users.NewMockRepository()
	rolesRepo := roles.NewMockRepository()

	// Create the router
	router := mux.NewRouter()

	// Create the routings
	router.HandleFunc("/", server.HomeHandler).Methods("GET")
	router.Use(server.LoggingMiddleware)

	// Create controllers
	rolesRestCtrl := roles.NewRestController(router, rolesRepo)
	usersRestCtrl := users.NewRestController(router, usersRepo)

	// Create server with parameters
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", appconfig.Port),
		WriteTimeout: time.Second * time.Duration(appconfig.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(appconfig.ReadTimeout),
		IdleTimeout:  time.Second * appconfig.IdleTimeout,
		Handler:      router,
	}

	// Start the server
	go func() {
		log.Info("Listening on " + server.Addr)
		log.Info("Press Ctrl+c to shutdown the server")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Waits until an interrupt is sent from the OS
	<-stop

	log.Info("Shutting down the server...")

	// Shutdown the server (default context)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	// Shutdown controllers
	rolesRestCtrl.Close()
	usersRestCtrl.Close()

	// Close Repositories
	rolesRepo.Close()
	usersRepo.Close()

	log.Info("Server gracefully stopped")
}
