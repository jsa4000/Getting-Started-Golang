package users

import (
	"encoding/json"
	"net/http"

	log "webapp/core/logging"
	valid "webapp/core/validation"

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
	router.HandleFunc("/users", ctrl.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", ctrl.GetByID).Methods("GET")
	router.HandleFunc("/users", ctrl.Create).Methods("POST")
	router.HandleFunc("/users/{id}", ctrl.DeleteByID).Methods("DELETE")
	return &ctrl
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Users Controller Shutdown")
}

// GetAll handler to request the
func (c *RestController) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := c.Service.GetAll(r.Context(), &GetAllRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Users)
}

// GetByID handler to request the
func (c *RestController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := GetByIDRequest{ID: vars["id"]}
	res, err := c.Service.GetByID(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.User)
}

// Create handler to request the
func (c *RestController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req CreateRequest
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	valid, err := valid.Validate(&req)
	if !valid && err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, err := c.Service.Create(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.User)
}

// DeleteByID handler to request the
func (c *RestController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := DeleteByIDRequest{ID: vars["id"]}
	_, err := c.Service.DeleteByID(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
