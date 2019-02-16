package http

import (
	"net/http"
	"sort"
	"strings"
)

// HandlerMid for handle the requests
type HandlerMid func(http.Handler) http.Handler

// Middleware interface for middleware to register
type Middleware interface {
	Handler() HandlerMid
	Priority() int
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

// Contains function to short middleware by priority
func Contains(source string, values []string) bool {
	for _, substr := range values {
		if strings.Contains(source, substr) {
			return true
		}
	}
	return false
}
