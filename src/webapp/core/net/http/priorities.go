package http

// Declare default priorities to load the components
const (
	// PriorityMax Maximun init priority
	PriorityMax int = -101
	// PriorityLogging Logging init priority
	PriorityLogging = -99
	// PriorityHeaders Headers init priority
	PriorityHeaders = -89
	// PriorityAuthentication Security init priority
	PriorityAuthentication = -79
	// PriorityAuthorization Security init priority
	PriorityAuthorization = -78

	//PriorityNormal
	PriorityNormal = 0

	// PriorityMin Maximun init priority
	PriorityMin int = 101
)
