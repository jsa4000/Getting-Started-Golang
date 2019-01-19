package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/core/config"
	"webapp/core/config/viper"
	log "webapp/core/logging"
	"webapp/core/logging/logrus"
	"webapp/roles"
	"webapp/server"
	"webapp/users"

	"github.com/gorilla/mux"
)

func setGlobalLogger() {
	log.SetGlobal(logrus.New())
	// Set the log formatter
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(log.TextFormat)
}

func setGlobalParser() {
	config.SetGlobal(viper.NewParserFromFile("webapp.yaml", "."))
}

func main() {

	// Set Global Logger
	setGlobalLogger()
	// Set Global Parser
	setGlobalParser()

	// Read the Configuration
	serverConfig := server.Config{}
	config.ReadFields(&serverConfig)

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	log.Infof("Starting %s Server...", serverConfig.Name)

	// Create Repositories
	rolesRepository := roles.NewMockRepository()
	usersRepository := users.NewMockRepository()

	// Create Services
	rolesService := roles.NewServiceImpl(rolesRepository)
	usersService := users.NewServiceImpl(usersRepository)

	// Create the router
	router := mux.NewRouter()

	// Create the routings
	router.HandleFunc("/", server.HomeHandler).Methods("GET")
	router.Use(server.LoggingMiddleware)

	// Create controllers
	rolesRestCtrl := roles.NewRestController(router, rolesService)
	usersRestCtrl := users.NewRestController(router, usersService)

	// Create server with parameters
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", serverConfig.Port),
		WriteTimeout: time.Second * time.Duration(serverConfig.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(serverConfig.ReadTimeout),
		IdleTimeout:  time.Second * serverConfig.IdleTimeout,
		Handler:      router,
	}

	// Start the server
	go func() {
		log.Info("Listening on " + server.Addr)
		log.Info("Press Ctrl+c to shutdown the server")
		if err := server.ListenAndServe(); err != nil {
			log.Error(err)
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
	rolesRepository.Close()
	usersRepository.Close()

	log.Info("Server gracefully stopped")
}
