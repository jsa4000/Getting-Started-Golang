package net

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"

	"github.com/gorilla/mux"
)

// HTTPHandler for handle the requests
type HTTPHandler func(w http.ResponseWriter, r *http.Request)

// HTTPMiddleware for handle the requests
type HTTPMiddleware func(http.Handler) http.Handler

// HTTPRoute route
type HTTPRoute struct {
	Path    string
	Method  string
	Handler HTTPHandler
	secured bool
	roles   []string
}

// HTTPConfig main app configuration
type HTTPConfig struct {
	Name         string        `config:"app.name:ServerApp"`
	LogLevel     string        `config:"logging.level:info"`
	Port         int           `config:"server.port:8080"`
	WriteTimeout int           `config:"server.writeTimeout:60"`
	ReadTimeout  int           `config:"server.readTimeout:60"`
	IdleTimeout  time.Duration `config:"server.idleTimeout:60"`
	Status       bool
}

// LoggingMiddleware decorator (closure)
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Received Request ", fmt.Sprintf("uri=%s args=%s ", r.RequestURI, mux.Vars(r)))
		next.ServeHTTP(w, r)
	})
}

// CustomHeaders decorator (closure)
func CustomHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// HTTPServer struct
type HTTPServer struct {
	Router *mux.Router
	Server *http.Server
	Config HTTPConfig
}

// NewHTTPServer create
func NewHTTPServer() *HTTPServer {
	// Read the Configuration
	serverConfig := HTTPConfig{}
	config.ReadFields(&serverConfig)

	// Create the router
	router := mux.NewRouter()

	// Create server with parameters
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", serverConfig.Port),
		WriteTimeout: time.Second * time.Duration(serverConfig.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(serverConfig.ReadTimeout),
		IdleTimeout:  time.Second * serverConfig.IdleTimeout,
		Handler:      router,
	}

	return &HTTPServer{
		Router: router,
		Server: server,
		Config: serverConfig,
	}
}

// AddRoutes to the router
func (h *HTTPServer) AddRoutes(routes ...HTTPRoute) {
	for _, r := range routes {
		h.Router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}

// AddMiddleware to the router
func (h *HTTPServer) AddMiddleware(mw ...HTTPMiddleware) {
	for _, m := range mw {
		h.Router.Use(mux.MiddlewareFunc(m))
	}
}

// Start server
func (h *HTTPServer) Start() {
	// Start the server
	go func() {
		log.Info("Listening on " + h.Server.Addr)
		if err := h.Server.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()
}

// Shutdown server
func (h *HTTPServer) Shutdown(ctx context.Context) {
	// Shutdown the server (default context)
	log.Info("Shutting down the server...")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)
}

//GetVars get vars from a request
func GetVars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
