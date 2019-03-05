package mngmt

import "time"

const (
	// StatusOk status is ok
	StatusOk = "OK"
	// StatusError there is an error in the system
	StatusError = "ERROR"
)

// GlobalHealth struct with the info ragring the health of the system
type GlobalHealth struct {
	Status     string    `json:"status"`
	Time       time.Time `json:"time"`
	Components []*Health `json:"components,omitempty"`
}

//NewGlobalHealth returns new Global Heath item
func NewGlobalHealth() *GlobalHealth {
	return &GlobalHealth{
		Status:     StatusError,
		Components: make([]*Health, 0),
	}
}

// Health structure to retrieve (fetch) the user information
type Health struct {
	Name   string        `json:"name"`
	Status string        `json:"status"`
	Values []interface{} `json:"values,omitempty"`
}

// Metrics struct with the info ragarding the health of the system
type Metrics []*Value

// Value struct with the info ragarding the health of the system
type Value struct {
	value interface{}
}
