package security

import "time"

// CreateTokenRequest request
type CreateTokenRequest struct {
	UserName string   `json:"username" validate:"min=0,max=255,required"`
	Password string   `json:"password" validate:"min=0,max=1024,required"`
	GranType string   `json:"grant_type" validate:"min=0,max=255,required"`
	Scope    []string `json:"scope" validate:"min=0,max=1024"`
}

// CreateTokenResponse Response
type CreateTokenResponse struct {
	Token          string
	ExpirationTime time.Time
}

// CheckTokenRequest struct request
type CheckTokenRequest struct {
	Token string
}

// CheckTokenResponse struct Response
type CheckTokenResponse struct {
	Data  interface{}
	Valid bool
}
