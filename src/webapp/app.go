package main

import (
	"context"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	"webapp/core/storage/mongo"
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

	db *mongo.Wrapper

	// Change to interface instead
	httpServer *net.Server
}

// GetUsersRepository factory that returns the User Repository
func (a *App) GetUsersRepository(t string) users.Repository {
	switch t {
	case "mock":
		return users.NewMockRepository()
	case "mongo":
		return users.NewMongoRepository(a.db)
	default:
		return users.NewMockRepository()
	}
}

// GetRolesRepository factory that returns the User Repository
func (a *App) GetRolesRepository(t string) roles.Repository {
	switch t {
	case "mock":
		return roles.NewMockRepository()
	case "mongo":
		return roles.NewMongoRepository(a.db)
	default:
		return roles.NewMockRepository()
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

	// Repository Type: mongo, mock
	rType := "mongo"

	// Create Database Driver
	if rType == "mongo" {
		a.db = mongo.New()
		a.db.Connect(ctx, "mongodb://root:root@dockerhost:27017/admin")
	}

	// Create Repositories
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

	if a.db != nil {
		a.db.Disconnect(ctx)
	}
	a.httpServer.Shutdown(ctx)

	log.Info("Server Stopped")
}
