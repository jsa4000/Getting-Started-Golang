package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"webapp/core/config"
	"webapp/core/config/viper"
	log "webapp/core/logging"
	"webapp/core/logging/logrus"
	"webapp/roles"
	"webapp/servers"
	"webapp/users"
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

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	log.Infof("Starting Services...")

	// Create Repositories
	rolesRepository := roles.NewMockRepository()
	usersRepository := users.NewMockRepository()

	// Create Services
	rolesService := roles.NewServiceImpl(rolesRepository)
	usersService := users.NewServiceImpl(usersRepository)

	// Create The HTTP Server
	httpServer := servers.NewHTTPServer()

	// Create controllers
	rolesRestCtrl := roles.NewRestController(httpServer.Router, rolesService)
	usersRestCtrl := users.NewRestController(httpServer.Router, usersService)

	// Start the HTTP server
	httpServer.Start()

	log.Info("Press Ctrl+c to shutdown the server")

	// Waits until an interrupt is sent from the OS
	<-stop

	log.Infof("Stopping Services...")

	ctx := context.Background()

	// Shutdown the HTTP server
	httpServer.Shutdown(ctx)

	// Shutdown controllers
	rolesRestCtrl.Close()
	usersRestCtrl.Close()

	// Close Repositories
	rolesRepository.Close()
	usersRepository.Close()

	log.Info("Server gracefully stopped")
}
