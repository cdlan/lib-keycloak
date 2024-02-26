package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

func (c *Config) Introspect(ctx context.Context, token string) (*[]gocloak.ResourcePermission, error) {

	// get socket
	s, err := c.newSocket(ctx)
	if err != nil {
		return nil, err
	}

	res, err := s.client.RetrospectToken(ctx, token, c.ClientID, c.Secret, c.Realm)
	if err != nil {
		return nil, err
	}

	if !*res.Active {
		return nil, fmt.Errorf("token is not active")
	}

	return res.Permissions, nil
}