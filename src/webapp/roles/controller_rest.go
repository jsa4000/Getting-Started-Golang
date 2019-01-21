package roles

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
	router.HandleFunc("/roles", ctrl.GetAll).Methods("GET")
	router.HandleFunc("/roles/{id}", ctrl.GetByID).Methods("GET")
	router.HandleFunc("/roles", ctrl.Create).Methods("POST")
	router.HandleFunc("/roles/{id}", ctrl.DeleteByID).Methods("DELETE")
	return &ctrl
}

// Close gracefully shutdown rest controller
func (c *RestController) Close() {
	log.Info("Role Controller Shutdown")
}

// GetAll handler to request the
func (c *RestController) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := c.Service.GetAll(r.Context(), &GetAllRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Roles)
}

// GetByID handler to request the
func (c *RestController) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := GetByIDRequest {ID : vars["id"]}
	res, err := c.Service.GetByID(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Role)
}

// Create handler to request the
func (c *RestController) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req CreateRequest
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	res, err := c.Service.Create(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Role)
}

// DeleteByID handler to request the
func (c *RestController) DeleteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := DeleteByIDRequest {ID : vars["id"]}
	_, err := c.Service.DeleteByID(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
