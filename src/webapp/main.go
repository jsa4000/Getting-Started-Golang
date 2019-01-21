package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"webapp/core/config"
	"webapp/core/config/viper"
	log "webapp/core/logging"
	"webapp/core/logging/logrus"
	trans "webapp/core/transport"
	"webapp/core/validation"
	"webapp/core/validation/goplayground"
	"webapp/roles"
	"webapp/users"
)

func setGlobalLogger() {
	log.SetGlobal(logrus.New())
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(log.TextFormat)
}

func setGlobalParser() {
	config.SetGlobal(viper.NewParserFromFile("webapp.yaml", "."))
}

func setGlobalValidator() {
	validation.SetGlobal(goplayground.New())
}

// HomeHandler handler for the HomePage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Info("Home Page")
}

func main() {

	// Set Global Logger
	setGlobalLogger()
	// Set Global Parser
	setGlobalParser()
	// Set global Validator
	setGlobalValidator()

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
	httpServer := trans.NewHTTPServer()

	// Create controllers
	rolesRestCtrl := roles.NewRestController(rolesService)
	usersRestCtrl := users.NewRestController(usersService)

	// Assigen controllers to the Http server
	httpServer.AddRoutes(rolesRestCtrl.GetRoutes()...)
	httpServer.AddRoutes(usersRestCtrl.GetRoutes()...)

	// Create the routings
	httpServer.AddRoutes(trans.HTTPRoute{
		Path:    "/",
		Method:  "GET",
		Handler: HomeHandler,
	})
	httpServer.AddMiddleware(trans.LoggingMiddleware)

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
