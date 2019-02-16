package http

// Security interface for components to register
type Security interface {
	Middleware() []Middleware
}
