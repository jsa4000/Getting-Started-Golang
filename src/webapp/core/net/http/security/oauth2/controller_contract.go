package oauth2

// Grant Types: OAuth 2 provides several "grant types" for different use cases.

// The specification describes five grants for acquiring an access token:
// - Authorization code grant (In-progress)
// - Implicit grant (Supported)
// - Resource owner credentials grant (Supported)
// - Client credentials grant (Supported)
// - Refresh token grant

const (
	// GrantTypeAuthorizationCode the client will redirect the user to the
	// authorization server
	GrantTypeAuthorizationCode = "authorization_code"
	// GrantTypePassword great user experience for trusted first party clients
	// both on the web and in native device applications
	GrantTypePassword = "password"
	// GrantTypeClientCredentials suitable for machine-to-machine authentication
	// where a specific user’s permission to access data is not required
	GrantTypeClientCredentials = "client_credentials"
	// GrantTypeRefreshToken Access tokens eventually expire; however some grants
	// respond with a refresh token which enables the client to get a new access
	// token without requiring the user to be redirected.
	GrantTypeRefreshToken = "refresh_token"
	// ResponseTypeCode intended used before authorization code which is exchanged ç
	// for an access token
	ResponseTypeCode = "code" // response_type
	// ResponseTypeImplicit intended to be used for user-agent-based clients (single page web)
	// that can’t keep a client secret because all of the application code and store
	// is easily accessible.
	ResponseTypeImplicit = "token" // response_type
)

// BasicOauth2Request request
type BasicOauth2Request struct {
	ClientID     string `json:"client_id" param:"client_id" validate:"min=0,max=1024"`
	ClientSecret string `json:"client_secret" param:"client_secret" validate:"min=0,max=1024"`
	UserName     string `json:"username" param:"username" validate:"min=0,max=255"`
	Password     string `json:"password" param:"password" validate:"min=0,max=1024"`
	GranType     string `json:"grant_type" param:"grant_type" validate:"min=0,max=255"`
	Scope        string `json:"scope" param:"scope" validate:"min=0,max=4096"`
	RedirectURI  string `json:"redirect_uri" param:"redirect_uri" validate:"min=0,max=4096"`
	ResponseType string `json:"response_type" param:"response_type" validate:"min=0,max=255"`
	State        string `json:"state" param:"state" validate:"min=0,max=2048"`
	Code         string `json:"code,omitempty" param:"code,omitempty" validate:"min=0,max=2048"`
}

// BasicOauth2Response Response
type BasicOauth2Response struct {
	AccessToken    string `json:"access_token" param:"access_token"`
	RefreshToken   string `json:"refresh_token,omitempty" param:"refresh_token,omitempty"`
	TokenType      string `json:"token_type,omitempty" param:"token_type,omitempty"`
	ExpirationTime int    `json:"expires_in" param:"expires_in"`
	State          string `json:"state,omitempty" param:"state,omitempty"`
	Code           string `json:"code,omitempty" param:"code,omitempty"`
}

// CheckTokenRequest struct request
type CheckTokenRequest struct {
	ClientID     string `json:"client_id" param:"client_id" validate:"min=0,max=1024"`
	ClientSecret string `json:"client_secret" param:"client_secret" validate:"min=0,max=1024"`
	Token        string `json:"token" param:"token"`
}

// CheckTokenResponse struct Response
type CheckTokenResponse struct {
	Data  interface{}
	Valid bool
}
