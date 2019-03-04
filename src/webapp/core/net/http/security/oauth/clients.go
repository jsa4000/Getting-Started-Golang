package oauth

import (
	"context"
	"fmt"
	pcrypt "webapp/core/crypto/password"
	net "webapp/core/net/http"
)

// Client structure to retrieve (fetch) the user information
type Client struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Secret string   `json:"secret"`
	Scopes []string `json:"scopes"`
}

// ClientService Interface
type ClientService interface {
	Fetch(ctx context.Context, name string) (*Client, error)
}

// ValidateClient compares the name and secret to be the same as the client
func ValidateClient(client *Client, name, secret string) bool {
	return name == client.Name && pcrypt.Compare(client.Secret, secret)
}

// Clients to implement the Client
type Clients map[string]*Client

// Fetch implements ClientService interface
func (c *Clients) Fetch(ctx context.Context, name string) (*Client, error) {
	client, ok := (*c)[name]
	if !ok {
		return nil, net.ErrNotFound.From(fmt.Errorf("Client %s has not been found", name))
	}
	return client, nil
}

// ClientsBuilder build struct
type ClientsBuilder struct {
	Clients
	current *Client
}

// NewClientsBuilder Create a new DefaultClientsBuilder
func NewClientsBuilder() *ClientsBuilder {
	return &ClientsBuilder{
		Clients: make(map[string]*Client),
	}
}

// WithClient set the interface to use for fetching user info
func (c *ClientsBuilder) WithClient(name string) *ClientsBuilder {
	if c.current != nil {
		c.Clients[c.current.Name] = c.current
	}
	c.current = &Client{
		Name: name,
	}
	return c
}

// WithSecret set the interface to use for fetching user info
func (c *ClientsBuilder) WithSecret(secret string) *ClientsBuilder {
	c.current.Secret = pcrypt.New(secret)
	return c
}

// WithScope set the interface to use for fetching user info
func (c *ClientsBuilder) WithScope(scopes ...string) *ClientsBuilder {
	c.current.Scopes = append(c.current.Scopes, scopes...)
	return c
}

// Build default users struct
func (c *ClientsBuilder) Build() *Clients {
	if c.current != nil {
		c.Clients[c.current.Name] = c.current
	}
	return &c.Clients
}
