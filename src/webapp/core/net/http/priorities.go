package http

// Declare default priorities to load the components
const (
	// PriorityMax Maximun init priority
	PriorityMax int = -101
	// PriorityLogging Logging init priority
	PriorityLogging = -99
	// PriorityHeaders Headers init priority
	PriorityHeaders = -89
	// PriorityAuth Security init priority
	PriorityAuth = -79
	// PriorityFilters Security init priority
	PriorityFilters = -69

	//PriorityNormal
	PriorityNormal = 0

	// PriorityMin Maximun init priority
	PriorityMin int = 101
)
