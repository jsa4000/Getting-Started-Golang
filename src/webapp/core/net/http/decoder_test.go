package http

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/go-test/deep"
)

// CreateTokenRequest request
type TestRequest struct {
	ClientID       string `json:"client_id" validate:"min=0,max=1024"`
	ClientSecret   string `json:"client_secret,omitempty" validate:"min=0,max=1024"`
	UserName       string `json:"username" validate:"min=0,max=255"`
	RedirectURI    string `json:"redirect_uri" validate:"min=0,max=4096"`
	ResponseType   string `json:"response_type" validate:"min=0,max=255"`
	ExpirationTime int    `json:"expires_in,omitempty"`
	Enabled        bool   `json:"enabled,omitempty"`
}

func TestDecoder(t *testing.T) {
	expeted := TestRequest{
		ClientID:       "my-client-name",
		ClientSecret:   "my-client-secret",
		UserName:       "my-user-name",
		ExpirationTime: 22,
		Enabled:        true,
	}
	url := fmt.Sprintf("/test?client_id=%s&client_secret=%s&username=%s&expires_in=%d&enabled=%s",
		expeted.ClientID, expeted.ClientSecret, expeted.UserName,
		expeted.ExpirationTime, strconv.FormatBool(expeted.Enabled))

	req, _ := http.NewRequest("GET", url, nil)
	message := TestRequest{}
	DecodeParams(req, &message, "json")

	if diff := deep.Equal(expeted, message); diff != nil {
		t.Error(diff)
	}
}
