package servers

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"webapp/core/config"
	log "webapp/core/logging"
	trans "webapp/core/transport"

	"github.com/gorilla/mux"
)

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

// HomeHandler handler for the HomePage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Info("Home Page")
}

// LoggingMiddleware decorator (closure)
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Received Request ", fmt.Sprintf("uri=%s args=%s ", r.RequestURI, mux.Vars(r)))
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

	// Create the routings
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.Use(LoggingMiddleware)

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
func (h *HTTPServer) AddRoutes(routes []trans.HTTPRoute) {
	for _, r := range routes {
		h.Router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
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
