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
	ClientID       string `param:"client_id" validate:"min=0,max=1024"`
	ClientSecret   string `param:"client_secret,omitempty" validate:"min=0,max=1024"`
	UserName       string `param:"username" validate:"min=0,max=255"`
	RedirectURI    string `param:"redirect_uri" validate:"min=0,max=4096"`
	ResponseType   string `param:"response_type" validate:"min=0,max=255"`
	ExpirationTime int    `param:"expires_in,omitempty"`
	Enabled        bool   `param:"enabled,omitempty"`
}

func TestDecoder(t *testing.T) {
	expected := TestRequest{
		ClientID:       "my-client-name",
		ClientSecret:   "my-client-secret",
		UserName:       "my-user-name",
		ExpirationTime: 22,
		Enabled:        true,
	}
	url := fmt.Sprintf("/test?client_id=%s&client_secret=%s&username=%s&expires_in=%d&enabled=%s",
		expected.ClientID, expected.ClientSecret, expected.UserName,
		expected.ExpirationTime, strconv.FormatBool(expected.Enabled))

	req, _ := http.NewRequest("GET", url, nil)
	message := TestRequest{}
	DecodeParams(req, &message)

	if diff := deep.Equal(expected, message); diff != nil {
		t.Error(diff)
	}
}
