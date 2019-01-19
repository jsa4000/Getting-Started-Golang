package users

import (
	"encoding/json"
	"net/http"

	log "webapp/core/logging"

	"github.com/gorilla/mux"
)

// RestController for http transport
type RestController struct {
	Service Service
}

// NewRestController create new RestController
func NewRestController(router *mux.Router, service Service) *RestController {
	ctrl := RestController{
		Service: service,
	}
	router.HandleFunc("/users", ctrl.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", ctrl.GetUserByID).Methods("GET")
	router.HandleFunc("/users", ctrl.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", ctrl.DeleteUserByID).Methods("DELETE")
	return &ctrl
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Users Controller Shutdown")
}

// GetAllUsers handler to request the
func (c *RestController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.GetAll(r.Context())
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
	user, err := c.Service.GetByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// CreateUser handler to request the
func (c *RestController) CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	user, err = c.Service.Create(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// DeleteUserByID handler to request the
func (c *RestController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := c.Service.DeleteByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
