package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// NewRestController create new RestController
func NewRestController(router *mux.Router, repo Repository) *RestController {
	ctrl := RestController{
		Repository: repo,
	}
	router.HandleFunc("/users", ctrl.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", ctrl.GetUserByID).Methods("GET")
	router.HandleFunc("/users", ctrl.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", ctrl.DeleteUserByID).Methods("DELETE")
	return &ctrl
}

// RestController for http transport
type RestController struct {
	Repository Repository
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Users Controller Shutdown")
}

// GetAllUsers handler to request the
func (c *RestController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Repository.FindAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// GetUserByID handler to request the
func (c *RestController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := c.Repository.FindByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
	err := c.Repository.DeleteByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
