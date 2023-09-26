package keycloak

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

// SearchUser search for a user based on email
func (c *Config) SearchUser(ctx context.Context, email string) (*gocloak.User, error) {

	// get socket
	s, err := c.newSocket(ctx)
	if err != nil {
		return nil, err
	}

	params := gocloak.GetUsersParams{
		Email: gocloak.StringP(email),
	}

	users, err := s.client.GetUsers(ctx, s.token.AccessToken, c.Realm, params)
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return users[0], nil
	}

	return nil, nil
}

// Actions returns a list of actions: verifyMail, updatePassword, configureTOTP
func (c *Config) Actions() *[]string {

	return &[]string{
		"VERIFY_EMAIL",
		"UPDATE_PASSWORD",
		"CONFIGURE_TOTP",
	}
}

// CreateUser create a new user
func (c *Config) CreateUser(ctx context.Context, user gocloak.User) error {

	// get socket
	s, err := c.newSocket(ctx)
	if err != nil {
		return err
	}

	// create user
	if _, err := s.client.CreateUser(ctx, s.token.AccessToken, c.Realm, user); err != nil {
		return err
	}

	return nil
}

// UpdateUser updates a user with the data from the param user
func (c *Config) UpdateUser(ctx context.Context, user gocloak.User) error {

	// get socket
	s, err := c.newSocket(ctx)
	if err != nil {
		return err
	}

	if err := s.client.UpdateUser(ctx, s.token.AccessToken, c.Realm, user); err != nil {
		return err
	}

	return nil
}

// DeleteUser deletes a user by its email
func (c *Config) DeleteUser(ctx context.Context, email string) error {

	// get socket
	s, err := c.newSocket(ctx)
	if err != nil {
		return err
	}

	// check if already exist
	found, err := c.SearchUser(ctx, email)
	if err != nil {
		return err
	}

	// if user found => delete
	if found != nil {

		err := s.client.DeleteUser(ctx, s.token.AccessToken, c.Realm, *found.ID)
		if err != nil {
			return err
		}
		return nil
	} else {

		return fmt.Errorf("user not found in kc")
	}
}

// SendSetupMail send mail to user with specified email to execute actions.
// actions can be "VERIFY_EMAIL","UPDATE_PASSWORD","UPDATE_PROFILE","CONFIGURE_TOTP"
func (c *Config) SendSetupMail(ctx context.Context, email string, actions []string) error {

	// get socket
	s, err := c.newSocket(ctx)
	if err != nil {
		return err
	}

	// search user
	found, err := c.SearchUser(ctx, email)
	if err != nil {
		return err
	}

	if found != nil {

		payload := gocloak.ExecuteActionsEmail{
			UserID:  found.ID,
			Actions: &actions,
		}

		err := s.client.ExecuteActionsEmail(ctx, s.token.AccessToken, c.Realm, payload)
		if err != nil {
			return err
		}

		return nil
	} else {

		return fmt.Errorf("user not found in kc")
	}
}