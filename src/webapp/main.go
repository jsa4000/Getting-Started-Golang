package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/core/config/viper"
	"webapp/core/logging"
	"webapp/core/logging/logrus"
	"webapp/roles"
	"webapp/server"
	"webapp/users"

	"github.com/gorilla/mux"
)

var logger = logrus.New()

func main() {

	logging.Log = logger

	// Set the log formatter
	logger.SetLevel(logging.DebugLevel)
	logger.SetFormatter(logging.TextFormat)

	// Create Parser (Configuration)
	parser := viper.NewParserFromFile("webapp.yaml", ".")

	// Read the Configuration
	serverConfig := server.Config{}
	parser.ReadFields(&serverConfig)

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	logger.Infof("Starting %s Server...", serverConfig.Name)

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
		logger.Info("Listening on " + server.Addr)
		logger.Info("Press Ctrl+c to shutdown the server")
		if err := server.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	// Waits until an interrupt is sent from the OS
	<-stop

	logger.Info("Shutting down the server...")

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

	logger.Info("Server gracefully stopped")
}
