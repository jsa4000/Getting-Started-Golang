package starter

// Declare default priorities to load the components
const (
	// PriorityMax Maximun init priority
	PriorityMax int = -101
	// PriorityConfig Maximun init priority
	PriorityConfig = -100
	// PriorityLogging Logging init priority
	PriorityLogging = -99
	// PriorityValidation Logging init priority
	PriorityValidation = -89
	// PriorityStorage Storage init priority
	PriorityStorage = -79
	// PriorityNet Logging init priority
	PriorityNet = -49

	//PriorityNormal
	PriorityNormal = 0

	// PriorityMin Maximun init priority
	PriorityMin int = 101
)
