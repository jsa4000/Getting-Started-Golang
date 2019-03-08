package main

import (
	"context"
	log "webapp/core/logging"

	mngmt "webapp/core/mngmt"
	net "webapp/core/net/http"
	pprof "webapp/core/net/http/pprof"
	"webapp/core/net/http/security"

	// Go-Core Starters
	_ "webapp/core/config/viper/starter"
	_ "webapp/core/logging/logrus/starter"
	_ "webapp/core/net/http/gorillamux/starter"
	_ "webapp/core/store/mongo/starter"
	_ "webapp/core/store/redis/starter"
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

	// Create repositories
	rolesRepository := roles.NewCacheRepository(roles.NewMongoRepository())
	usersRepository := users.NewMongoRepository()

	// Create Services
	rolesService := roles.NewServiceImpl(rolesRepository)
	usersService := users.NewServiceImpl(usersRepository)

	// Create Management for health, metrics, etc..
	mngmt := mngmt.NewManagerBuilder().
		WithMetrics(true).
		WithRootPath("/management").
		Build()

	// Create Security for HTTP
	usersManager := usersManager(usersService)
	recurity := security.NewBuilder().
		WithAuthentication(oAuthManager(jwtService(usersManager))).
		WithAuthorization(openHandler(), basicHandler(usersManager), jwtHandler()).
		WithFilter(rolesHandler(), corsAuthFilter()).
		Build()

	// Create The HTTP Server
	a.httpServer = net.NewServer().
		WithControllers(mngmt.Controller()).
		WithControllers(pprof.NewController()).                 // Add Controller for Profiling
		WithControllers(roles.NewRestController(rolesService)). // Add roles controller
		WithControllers(users.NewRestController(usersService)). // Add users controller
		WithStatic("/swaggerui/", "./static/swaggerui/").       // Create swagger static content '/swagger/index.html'
		WithMiddleware(net.NewLoggingMiddleware()).             // Add global middlewares
		WithSecurity(recurity).                                 // Add security to HTTP Requests
		Start()                                                 // Start the HTTP server
}

// Shutdown the server
func (a *App) Shutdown(ctx context.Context) {
	log.Info("Server is shutting down")

	a.httpServer.Shutdown(ctx)

	log.Info("Server Stopped")
}
