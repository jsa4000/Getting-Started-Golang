package http

// Declare default priorities to load the components
const (
	// PriorityMax Maximun init priority
	PriorityMax int = -101
	// PriorityLogging Logging init priority
	PriorityLogging = -99
	// PriorityHeaders Headers init priority
	PriorityHeaders = -89
	// PriorityAuthorization Security init priority
	PriorityAuthorization = -79
	// PriorityResourceFilters Security init priority
	PriorityResourceFilters = -69

	//PriorityNormal
	PriorityNormal = 0

	// PriorityMin Maximun init priority
	PriorityMin int = 101
)
