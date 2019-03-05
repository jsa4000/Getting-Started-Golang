package mngmt

// Checker component that returns Health information
type Checker interface {
	Status() *Health
}
