package server

import (
	"context"
	log "webapp/core/logging"
	net "webapp/core/net/http"
	pprof "webapp/core/net/http/pprof"
	httpSec "webapp/core/net/http/security"

	// Go-Core Starters
	_ "webapp/core/config/viper/starter"
	_ "webapp/core/logging/logrus/starter"
	_ "webapp/core/net/http/gorillamux/starter"
	_ "webapp/core/storage/mongo/starter"
	_ "webapp/core/validation/goplayground/starter"

	"webapp/roles"
	"webapp/security"
	"webapp/users"
)

// App struct
type App struct {
	httpServer *net.Server
}

//Startup the server
func (a *App) Startup(ctx context.Context) {
	log.Infof("Starting Services...")

	// Create repositories
	rolesRepository := roles.NewMongoRepository()
	usersRepository := users.NewMongoRepository()

	// Create Services
	rolesService := roles.NewServiceImpl(rolesRepository)
	usersService := users.NewServiceImpl(usersRepository)
	credentialService := security.NewCredentialService(usersRepository)

	// Security Config
	secConfig := httpSec.NewConfig().
		WithUserCallback(credentialService)

	// Create The HTTP Server
	a.httpServer = net.NewServer().
		WithControllers(pprof.NewController()).                                       // Add Controller for Profiling
		WithControllers(roles.NewRestController(rolesService)).                       // Add roles controller
		WithControllers(users.NewRestController(usersService)).                       // Add users controller
		WithStatic("/swagger/", "./static/swaggerui/").                               // Create swagger static content '/swagger/index.html'
		WithMiddleware(net.NewLoggingMiddleware(), net.NewCustomHeadersMiddleware()). // Add global middlewares
		WithSecurity(httpSec.New(secConfig)).                                         // Add security to http requests
		Start()                                                                       // Start the HTTP server
}

// Shutdown the server
func (a *App) Shutdown(ctx context.Context) {
	log.Info("Server is shutting down")

	a.httpServer.Shutdown(ctx)

	log.Info("Server Stopped")
}
