package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/roles"
	"webapp/users"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// HomeHandler handler for the HomePage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Roles: %v\n", vars["Roles"])
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {

	// Create a channel to detect interrupt signal from os
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Set the log formatter
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Info("Starting the Server...")

	// Create the router
	router := mux.NewRouter()

	// Create the routings
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.Use(loggingMiddleware)

	// Create Repositiories
	usersRepo := users.NewMockRepository()

	// Create controllers
	rolesRestCtrl := roles.NewRestController(router)
	usresRestCtrl := users.NewRestController(router, usersRepo)

	// Create server with parameters
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Start the server
	go func() {
		log.Info("Listening on " + srv.Addr)
		log.Info("Press Ctrl+c to shutdown the server")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Waits until an interrupt is sent from the OS
	<-stop

	log.Info("Shutting down the server...")

	// Shutdown the server (defaukt context)
	srv.Shutdown(context.Background())

	// Shutdown controllers
	rolesRestCtrl.Close()
	usresRestCtrl.Close()

	// Close Repositories
	usersRepo.Close()

	log.Info("Server gracefully stopped")

}
