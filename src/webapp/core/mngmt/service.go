package mngmt

import "context"

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

// RuntimeRequest struct to request the metrics
type RuntimeRequest struct {
}

// RuntimeResponse struct with the response for the metrics
type RuntimeResponse struct {
	Runtime *Runtime
}

// Service Interface for Management
type Service interface {
	Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error)
	Runtime(ctx context.Context, req *RuntimeRequest) (*RuntimeResponse, error)
	Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error)
}
