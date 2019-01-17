package roles

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// RestController for http transport
type RestController struct {
	Repository Repository
}

// NewRestController create new RestController
func NewRestController(router *mux.Router, repo Repository) *RestController {
	ctrl := RestController{
		Repository: repo,
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
	roles, err := c.Repository.FindAll(r.Context())

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
	role, err := c.Repository.FindByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(role)
}

// CreateRole handler to request the
func (c *RestController) CreateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Roles: %v\n", vars["Roles"])
}

// DeleteRoleByID handler to request the
func (c *RestController) DeleteRoleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := c.Repository.DeleteByID(r.Context(), vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
