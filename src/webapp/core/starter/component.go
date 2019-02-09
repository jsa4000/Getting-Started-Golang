package starter

import (
	"context"
	"sort"
)

// Component interface for components to register
type Component interface {
	Init(ctx context.Context)
	Close(ctx context.Context)
	Priority() int
}

// byPriority array for component
type byPriority []Component

func (c byPriority) Len() int {
	return len(c)
}
func (c byPriority) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byPriority) Less(i, j int) bool {
	return c[i].Priority() < c[j].Priority()
}

func sortedComponents(asc bool) []Component {
	comps := Components()
	// Order the components to be initialized in order
	result := make([]Component, 0, len(comps))
	for _, c := range comps {
		result = append(result, c)
	}
	sort.Sort(byPriority(result))
	return result
}

// Init initialize defaults values
func Init(ctx context.Context) {
	// Initialize all the components autoloaded
	for _, c := range sortedComponents(true) {
		c.Init(ctx)
	}
}

// Shutdown initialize defaults values
func Shutdown(ctx context.Context) {
	// Shutdown all the components autoloaded
	for _, c := range sortedComponents(false) {
		c.Close(ctx)
	}
}
