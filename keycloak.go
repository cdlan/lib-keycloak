package keycloak

import (
	"context"
	"fmt"
	
	"github.com/Nerzal/gocloak/v13"
)

type Socket struct {
	client *gocloak.GoCloak
	token  *gocloak.JWT
}

// newSocket Initialize Socket global variable with vars
func (c *Config) newSocket(ctx context.Context) (*Socket, error) {

	s := Socket{
		client: gocloak.NewClient(c.Url),
	}

	// get new token
	token, err := s.client.LoginClient(ctx, c.ClientID, c.Secret, c.Realm)
	if err != nil {
		return nil, fmt.Errorf("something wrong with the credentials or url: %v", err)
	}
	s.token = token

	return &s, nil
}


