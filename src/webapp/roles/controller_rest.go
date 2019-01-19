package roles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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
	router.HandleFunc("/roles", ctrl.GetAllRoles).Methods("GET")
	router.HandleFunc("/roles/{id}", ctrl.GetRoleByID).Methods("GET")
	router.HandleFunc("/roles", ctrl.CreateRole).Methods("POST")
	router.HandleFunc("/roles/{id}", ctrl.DeleteRoleByID).Methods("DELETE")
	return &ctrl
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Role Controller Shutdown")
}

// GetAllRoles handler to request the
func (c *RestController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := c.Service.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(roles)
}

// GetRoleByID handler to request the
func (c *RestController) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role, err := c.Service.GetByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}

// CreateRole handler to request the
func (c *RestController) CreateRole(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var role Role
	err := decoder.Decode(&role)
	if err != nil {
		panic(err)
	}
	role, err = c.Service.Create(r.Context(), role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}

// DeleteRoleByID handler to request the
func (c *RestController) DeleteRoleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := c.Service.DeleteByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
