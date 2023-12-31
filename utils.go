package keycloak

import "github.com/Nerzal/gocloak/v13"

// PasswordCredentialRep generate credential representation with just this password as a temporary password
func PasswordCredentialRep(pass string) gocloak.CredentialRepresentation {

	passType := "password"
	trueVal := true

	return gocloak.CredentialRepresentation{
		Type:      &passType,
		Value:     &pass,
		Temporary: &trueVal,
	}
}

// DefaultActions returns a list of actions: verifyMail, updatePassword, configureTOTP
func DefaultActions() *[]string {

	return &[]string{
		"VERIFY_EMAIL",
		"UPDATE_PASSWORD",
		"CONFIGURE_TOTP",
	}
}
