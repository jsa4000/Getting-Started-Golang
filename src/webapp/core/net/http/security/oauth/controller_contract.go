package oauth

// Grant Types: OAuth 2 provides several "grant types" for different use cases.

// The grant types defined are:
// - Authorization Code
// - Password (Resource Owner password credentials grant)
// - Client credentials
// - Implicit

// CreateTokenRequest request
type CreateTokenRequest struct {
	ClientID     string   `json:"client_id" validate:"min=0,max=1024"`
	ClientSecret string   `json:"client_secret" validate:"min=0,max=1024"`
	UserName     string   `json:"username" validate:"min=0,max=255"`
	Password     string   `json:"password" validate:"min=0,max=1024"`
	GranType     string   `json:"grant_type" validate:"min=0,max=255"`
	Scope        []string `json:"scope" validate:"min=0,max=4096"`
	RedirectURI  string   `json:"redirect_uri" validate:"min=0,max=4096"`
	ResponseType string   `json:"response_type" validate:"min=0,max=255"`
}

// CreateTokenResponse Response
type CreateTokenResponse struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token,omitempty"`
	TokenType      string `json:"token_type,omitempty"`
	ExpirationTime int    `json:"expires_in"`
}

// CheckTokenRequest struct request
type CheckTokenRequest struct {
	Token string `json:"token" validate:",required"`
}

// CheckTokenResponse struct Response
type CheckTokenResponse struct {
	Data  interface{}
	Valid bool
}
