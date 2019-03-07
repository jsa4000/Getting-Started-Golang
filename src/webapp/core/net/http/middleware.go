package http

import (
	"net/http"
	"sort"
)

// RouteInfoKey Key used to extract rout information for filters
const RouteInfoKey = "route"

// HandlerMid for handle the requests
type HandlerMid func(http.Handler) http.Handler

// Middleware interface for middleware to register
type Middleware interface {
	Handler() HandlerMid
	Priority() int
}

// MiddlewareBase interface for components to register
type MiddlewareBase struct {
	Hdlr HandlerMid
	Prio int
}

// Handler returns the HandlerMid
func (m *MiddlewareBase) Handler() HandlerMid {
	return m.Hdlr
}

// Priority returns the priority
func (m *MiddlewareBase) Priority() int {
	return m.Prio
}

// Write response
func (m *MiddlewareBase) Error(w http.ResponseWriter, err error) {
	Error(w, err)
}

// JSON Sets the error from inner layers
func (m *MiddlewareBase) JSON(w http.ResponseWriter, body interface{}, code int) {
	JSON(w, body, code)
}

// byPriority array for Middleware
type byPriority []Middleware

func (c byPriority) Len() int {
	return len(c)
}
func (c byPriority) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byPriority) Less(i, j int) bool {
	return c[i].Priority() < c[j].Priority()
}

// SortMiddleware function to short middleware by priority
func SortMiddleware(m []Middleware, asc bool) []Middleware {
	// Order the components to be initialized in order
	result := make([]Middleware, 0, len(m))
	for _, c := range m {
		result = append(result, c)
	}
	if !asc {
		sort.Sort(sort.Reverse(byPriority(result)))
		return result
	}
	sort.Sort(byPriority(result))
	return result
}

// SplitMiddleware Splits the middleware into  global and filter depending on the priority
func SplitMiddleware(middleware []Middleware) ([]Middleware, []Middleware) {
	global := make([]Middleware, 0)
	filters := make([]Middleware, 0)
	for _, m := range SortMiddleware(middleware, true) {
		if m.Priority() >= PriorityFilters {
			filters = append(filters, m)
			continue
		}
		global = append(global, m)
	}
	return global, filters
}
