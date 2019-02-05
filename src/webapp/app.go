package main

import (
	"context"
	"net/http"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	pprof "webapp/core/net/http/pprof"

	_ "webapp/core/config/viper/starter"
	_ "webapp/core/logging/logrus/starter"
	_ "webapp/core/net/http/gorillamux/starter"
	_ "webapp/core/storage/mongo/starter"
	_ "webapp/core/validation/goplayground/starter"

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
	a.rolesRepository = roles.NewMongoRepository()
	a.usersRepository = users.NewMongoRepository()

	// Create Services
	a.rolesService = roles.NewServiceImpl(a.rolesRepository)
	a.usersService = users.NewServiceImpl(a.usersRepository)

	// Create controllers
	a.rolesRestCtrl = roles.NewRestController(a.rolesService)
	a.usersRestCtrl = users.NewRestController(a.usersService)

	// Create The HTTP Server
	a.httpServer = net.NewServer()

	// Add Controller for Profiling
	a.httpServer.AddController(pprof.NewController())

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
}

// Shutdown the server
func (a *App) Shutdown(ctx context.Context) {
	log.Info("Server is shutting down")

	a.httpServer.Shutdown(ctx)

	log.Info("Server Stopped")
}
