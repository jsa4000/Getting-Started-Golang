package mngmt

import "context"

// Checker component that returns Health information
type Checker interface {
	Status() *Health
}

// HealthRequest struct to request the health
type HealthRequest struct {
}

// HealthResponse struct with the response for the health
type HealthResponse struct {
	Health *GlobalHealth
}

// MetricsRequest struct to request the metrics
type MetricsRequest struct {
}

// MetricsResponse struct with the response for the metrics
type MetricsResponse struct {
	Metrics Metrics
}

// Service Interface for Management
type Service interface {
	Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error)
	Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error)
}
