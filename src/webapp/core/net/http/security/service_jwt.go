package security

import (
	"context"
	"time"
	log "webapp/core/logging"
)

// ServiceJwt Implementation used for the service
type ServiceJwt struct {
	config *Config
}

// NewServiceJwt Create a new ServiceImpl
func NewServiceJwt(config *Config) Service {
	return &ServiceJwt{
		config: config,
	}
}

// CreateToken create the token
func (s *ServiceJwt) CreateToken(ctx context.Context, req *CreateTokenRequest) (*CreateTokenResponse, error) {
	log.Debug("Create Token Request: ", req)

	user, err := s.config.uc.GetUserByName(ctx, req.UserName)

	if err != nil {
		return nil, err
	}

	log.Debug("User Founded: ", user)

	return &CreateTokenResponse{
		Token:  "Bearer 3243",
		Expire: time.Now().Add(time.Duration(int64(time.Millisecond) * int64(s.config.expiretime))),
	}, nil
}

// CheckToken returns deserialized token
func (s *ServiceJwt) CheckToken(ctx context.Context, req *CheckTokenRequest) (*CheckTokenResponse, error) {
	log.Debug("Check Token Request: ", req)
	return &CheckTokenResponse{
		Data: string([]byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)),
	}, nil
}
