package mngmt

import (
	"context"
	"time"
	"webapp/core/starter"
	ctime "webapp/core/time"
)

// ServiceImpl for management purposes such as metrics, health status,
// environment variables, etc..
type ServiceImpl struct {
	metrics         bool
	health          bool
	metricsSnapshot Metrics
	healthSnapshot  *GlobalHealth
	lastSnapshot    time.Time
	refreshTime     time.Duration
}

// NewServiceImpl creates new service impl
func NewServiceImpl(health bool, metrics bool, seconds int) *ServiceImpl {
	return &ServiceImpl{
		metrics:      metrics,
		health:       health,
		lastSnapshot: time.Time{},
		refreshTime:  time.Duration(seconds) * time.Second,
	}
}

// Snapshot retrieve an snapshot for all the values
func (s *ServiceImpl) Snapshot() {
	now := ctime.Now()
	if now.Sub(s.lastSnapshot) < s.refreshTime {
		return
	}
	// Set the lastSnapshot time. Since it is already eventual, this set will
	// avoid simultaneous requests to take snapshot at the same time.
	s.lastSnapshot = now
	if s.health {
		health := &GlobalHealth{
			Status: StatusOk,
			Time:   s.lastSnapshot,
		}
		for _, c := range starter.Components() {
			if h, ok := c.(Checker); ok {
				status := h.Status()
				if status.Status == StatusError {
					health.Status = StatusError
				}
				health.Components = append(health.Components, status)
			}
		}
		s.healthSnapshot = health
	}
	if s.metrics {
		metrics := make([]*Value, 0)

		s.metricsSnapshot = metrics
	}
}

// Health to retrieve the health of the system
func (s *ServiceImpl) Health(ctx context.Context, req *HealthRequest) (*HealthResponse, error) {
	s.Snapshot()
	return &HealthResponse{
		Health: s.healthSnapshot,
	}, nil
}

// Metrics to caputure an snapshot of the system
func (s *ServiceImpl) Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	s.Snapshot()
	return &MetricsResponse{
		Metrics: s.metricsSnapshot,
	}, nil
}
