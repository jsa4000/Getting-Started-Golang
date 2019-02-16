package server

import (
	"context"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	pprof "webapp/core/net/http/pprof"

	// Go-Core Starters
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
	httpServer *net.Server
}

//Startup the server
func (a *App) Startup(ctx context.Context) {
	log.Infof("Starting Services...")

	// Create Services
	rolesService := roles.NewServiceImpl(roles.NewMongoRepository())
	usersService := users.NewServiceImpl(users.NewMongoRepository())

	// Create The HTTP Server
	a.httpServer = net.NewServer().
		AddController(pprof.NewController()).                                        // Add Controller for Profiling
		AddController(roles.NewRestController(rolesService)).                        // Add roles controller
		AddController(users.NewRestController(usersService)).                        // Add users controller
		Static("/swagger/", "./static/swaggerui/").                                  // Create swagger static content '/swagger/index.html'
		AddMiddleware(net.NewLoggingMiddleware(), net.NewCustomHeadersMiddleware()). // Add global middlewares
		Start()                                                                      // Start the HTTP server
}

// Shutdown the server
func (a *App) Shutdown(ctx context.Context) {
	log.Info("Server is shutting down")

	a.httpServer.Shutdown(ctx)

	log.Info("Server Stopped")
}
