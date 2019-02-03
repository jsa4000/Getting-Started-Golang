package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"
)

// Log Global Router
var router Router

// SetGlobal sets the Global Logger (singletone)
func SetGlobal(r Router) {
	router = r
}

const nanoseconds = 1e6

const pprofPreffix = "/debug/pprof/"

// Handler for handle the requests
type Handler func(w http.ResponseWriter, r *http.Request)

// Middleware for handle the requests
type Middleware func(http.Handler) http.Handler

// Route route
type Route struct {
	Path    string
	Method  string
	Handler Handler
	secured bool
	roles   []string
}

// Router to handle http requests
type Router interface {
	Handler() http.Handler
	HandleRoute(route ...Route)
	Use(m ...Middleware)
	Vars(r *http.Request) map[string]string
}

// Config main app configuration
type Config struct {
	Name         string        `config:"app.name:ServerApp"`
	LogLevel     string        `config:"logging.level:info"`
	Port         int           `config:"server.port:8080"`
	WriteTimeout int           `config:"server.writeTimeout:60"`
	ReadTimeout  int           `config:"server.readTimeout:60"`
	IdleTimeout  time.Duration `config:"server.idleTimeout:60"`
	Status       bool
}

// Controller to handle http requests
type Controller interface {
	GetRoutes() []Route
}

// LoggingMiddleware decorator (closure)
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Received Request ", fmt.Sprintf("uri=%s args=%s ", r.RequestURI, Vars(r)))
		start := time.Now()
		defer func() {
			log.Debug(fmt.Sprintf("Processed Response in %.7f ms", float64(time.Since(start))/nanoseconds))
		}()
		next.ServeHTTP(w, r)
	})
}

// CustomHeaders decorator (closure)
func CustomHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.RequestURI, pprofPreffix) {
			w.Header().Set("Content-Type", "text/html")
		} else {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

// Server struct
type Server struct {
	Server *http.Server
	Config Config
}

// NewServer create
func NewServer() *Server {
	// Read the Configuration
	serverConfig := Config{}
	config.ReadFields(&serverConfig)

	// Create server with parameters
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", serverConfig.Port),
		WriteTimeout: time.Second * time.Duration(serverConfig.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(serverConfig.ReadTimeout),
		IdleTimeout:  time.Second * serverConfig.IdleTimeout,
	}

	return &Server{
		Server: server,
		Config: serverConfig,
	}
}

// AddController Add the controller to the router
func (h *Server) AddController(c Controller) {
	h.AddRoutes(c.GetRoutes()...)
}

// AddRoutes to the router
func (h *Server) AddRoutes(routes ...Route) {
	router.HandleRoute(routes...)
}

// AddMiddleware to the router
func (h *Server) AddMiddleware(mw ...Middleware) {
	router.Use(mw...)
}

// Start server
func (h *Server) Start() {
	// Start the server
	go func() {
		log.Info("Listening on " + h.Server.Addr)
		h.Server.Handler = router.Handler()
		if err := h.Server.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()
}

// Shutdown server
func (h *Server) Shutdown(ctx context.Context) {
	// Shutdown the server (default context)
	log.Info("Shutting down the server...")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
}

//Vars get vars from a request
func Vars(r *http.Request) map[string]string {
	return router.Vars(r)
}
