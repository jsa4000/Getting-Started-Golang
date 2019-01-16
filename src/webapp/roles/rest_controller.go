package roles

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// RestController for http transport
type RestController struct {
}

// NewRestController create new RestController
func NewRestController(router *mux.Router) *RestController {
	ctrl := RestController{}
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
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Roles: %v\n", vars["Roles"])
}

// GetRoleByID handler to request the
func (c *RestController) GetRoleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Roles: %v\n", vars["Roles"])
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Roles: %v\n", vars["Roles"])
}
