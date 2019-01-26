package main

import (
	"context"
	"net/http"
	"webapp/core/database/mongo"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/roles"
	"webapp/users"
)

// App struct
type App struct {
	rolesRepository roles.Repository
	rolesService    roles.Service
	rolesRestCtrl   net.Controller

	usersRepository users.Repository
	usersService    users.Service
	usersRestCtrl   net.Controller

	// Mongo Database Client
	mongodb *mongo.Client

	// Change to interface instead
	httpServer *net.Server
}

// GetUsersRepository factory that returns the User Repository
func (a *App) GetUsersRepository(t string) users.Repository {
	switch t {
	case "mock":
		return users.NewMockRepository()
	case "mongo":
		return users.NewMongoRepository(a.mongodb)
	default:
		return users.NewMongoRepository(a.mongodb)
	}
}

// GetRolesRepository factory that returns the User Repository
func (a *App) GetRolesRepository(t string) roles.Repository {
	switch t {
	case "mock":
		return roles.NewMockRepository()
	case "mongo":
		return roles.NewMongoRepository(a.mongodb)
	default:
		return roles.NewMongoRepository(a.mongodb)
	}
}

// HomeHandler handler for the HomePage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Info("Home Page")
}

//Startup the server
func (a *App) Startup(ctx context.Context) {

	log.Infof("Starting Services...")

	// Create Database Driver
	a.mongodb = mongo.New()
	a.mongodb.Connect("mongodb://root:root@dockerhost:27017/admin")

	// Create Repositories
	rType := "mongo"
	a.rolesRepository = a.GetRolesRepository(rType)
	a.usersRepository = a.GetUsersRepository(rType)

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

	a.usersRestCtrl.Shutdown()
	a.rolesRestCtrl.Shutdown()
	a.usersRepository.Close()
	a.rolesRepository.Close()

	if a.mongodb != nil {
		a.mongodb.Disconnect()
	}
	a.httpServer.Shutdown(context.Background())

	log.Info("Server Stopped")
}
