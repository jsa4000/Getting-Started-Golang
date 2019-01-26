package main

import (
	"context"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/roles"
	"webapp/users"
)

// App stuct
type App struct {
	rolesRepository roles.Repository
	rolesService    roles.Service
	rolesRestCtrl   *roles.RestController

	usersRepository users.Repository
	usersService    users.Service
	usersRestCtrl   *users.RestController

	// Change to interface instead
	httpServer *net.Server
}

// HomeHandler handler for the HomePage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Info("Home Page")
}

//Startup the server
func (a *App) Startup(ctx context.Context) {

	log.Infof("Starting Services...")

	// Create Repositories
	a.rolesRepository = roles.NewMockRepository()
	a.usersRepository = users.NewMockRepository()

	// Create Services
	a.rolesService = roles.NewServiceImpl(a.rolesRepository)
	a.usersService = users.NewServiceImpl(a.usersRepository)

	// Create The HTTP Server
	a.httpServer = net.NewServer()

	// Create controllers
	a.rolesRestCtrl = roles.NewRestController(a.rolesService)
	a.usersRestCtrl = users.NewRestController(a.usersService)

	// Add controllers to the Http server
	a.httpServer.AddController(a.rolesRestCtrl)
	a.httpServer.AddController(a.usersRestCtrl)

	// Create main homepage http route
	a.httpServer.AddRoutes(net.Route{
		Path:    "/",
		Method:  "GET",
		Handler: HomeHandler,
	})

	// Add global middlewares
	a.httpServer.AddMiddleware(net.LoggingMiddleware, net.CustomHeaders)

	// Start the HTTP server
	a.httpServer.Start()

	// log.Info("Press Ctrl+c to shutdown the server")
}

// Shutdown the server
func (a *App) Shutdown(ctx context.Context) {
	log.Info("Server is shutting down")

	a.usersRestCtrl.Close()
	a.rolesRestCtrl.Close()
	a.usersRepository.Close()
	a.rolesRepository.Close()

	a.httpServer.Shutdown(context.Background())

	log.Info("Server Stopped")
}
