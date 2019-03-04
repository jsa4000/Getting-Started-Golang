package security

import (
	"net/http"
	"sort"
	log "webapp/core/logging"
	net "webapp/core/net/http"
)

// AuthHandler type redefinition
type AuthHandler = FilterHandler

// FilterHandler interface to manage the authorization method
type FilterHandler interface {
	Priority() int
	Matches(url string) (Target, bool)
	Handle(w http.ResponseWriter, r *http.Request, target Target) error
}

// BaseHandler struct to handle access control methods
type BaseHandler struct {
	Targets
	Prior int
}

// Priority any target with the given URL
func (b *BaseHandler) Priority() int {
	return b.Prior
}

// Matches any target with the given URL
func (b *BaseHandler) Matches(url string) (Target, bool) {
	for _, t := range b.Targets {
		if t.Matches(url) {
			return t, true
		}
	}
	return nil, false
}

// Handle handler to manage access control methods
func (b *BaseHandler) Handle(w http.ResponseWriter, r *http.Request, target Target) error {
	log.Debugf("Handle Request for %s", net.RemoveURLParams(r.RequestURI))
	return nil
}

// byPriority array for FilterHandler
type byPriority []FilterHandler

func (c byPriority) Len() int {
	return len(c)
}
func (c byPriority) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byPriority) Less(i, j int) bool {
	return c[i].Priority() < c[j].Priority()
}

// SortFilters function to short middleware by priority
func SortFilters(m []FilterHandler, asc bool) []FilterHandler {
	result := make([]FilterHandler, 0, len(m))
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
