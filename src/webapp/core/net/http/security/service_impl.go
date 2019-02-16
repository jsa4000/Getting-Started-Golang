package security

import (
	"context"
	"time"
	log "webapp/core/logging"
)

// ServiceImpl Implementation used for the service
type ServiceImpl struct {
}

// NewServiceImpl Create a new ServiceImpl
func NewServiceImpl() Service {
	return &ServiceImpl{}
}

// CreateToken create the token
func (s *ServiceImpl) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	log.Debug("Create Token Request: ", req)
	return &CreateTokenResponse{
		Token:  "Bearer 3243",
		Expire: time.Now().Add(time.Duration(int64(time.Millisecond) * int64(60000))),
	}, nil
}

// CheckToken returns deserialized token
func (s *ServiceImpl) CheckToken(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
	log.Debug("Check Token Request: ", req)
	return &CheckTokenResponse{
		Data: string([]byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)),
	}, nil
}
