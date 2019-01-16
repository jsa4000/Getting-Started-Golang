package main

import (
	"fmt"
	"net/http"
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
	// Set the log formatter
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.Debug("Staring Sever...")

	// Create the router
	router := mux.NewRouter()

	// Create the routings
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.Use(loggingMiddleware)

	// Create other controllers
	rc := roles.NewRestController(router)
	uc := users.NewRestController(router)

	// Create server with parameters
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Start the server
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	// Shotdown all controller and services
	rc.Close()
	uc.Close()

	log.Debug("Shutdown Server")

}
