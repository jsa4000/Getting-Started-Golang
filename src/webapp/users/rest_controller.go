package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// NewRestController create new RestController
func NewRestController(router *mux.Router) *RestController {
	ctrl := RestController{}
	router.HandleFunc("/users", ctrl.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", ctrl.GetUserByID).Methods("GET")
	router.HandleFunc("/users", ctrl.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", ctrl.DeleteUserByID).Methods("DELETE")
	return &ctrl
}

// RestController for http transport
type RestController struct {
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Users Controller Shutdown")
}

// GetAllUsers handler to request the
func (c *RestController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v\n", vars["Users"])
}

// GetUserByID handler to request the
func (c *RestController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v\n", vars["Users"])
}

// CreateUser handler to request the
func (c *RestController) CreateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v\n", vars["Users"])
}

// DeleteUserByID handler to request the
func (c *RestController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Users: %v\n", vars["Users"])
}
