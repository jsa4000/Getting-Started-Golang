package server

import (
	"fmt"
	"net/http"
	log "webapp/core/logging"

	"github.com/gorilla/mux"
)

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
