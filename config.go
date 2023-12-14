package keycloak

import (
	"os"
	"strings"
)

// Config is compatible with github.com/spf13/viper yaml unmarshal
type Config struct {

	// Url is the url of your KeyCloak server ending in `/auth/`
	// ES: https://kc.domain.example/auth/
	Url string `mapstructure:"base_url"`

	// ClientID is the oauth2 client_id
	ClientID string `mapstructure:"client_id"`

	// Secret is the oauth2 client_secret
	Secret string `mapstructure:"client_secret"`

	// Realm is the name of the realm
	Realm string `mapstructure:"realm"`

	SetupActions []string `mapstructure:"setup_actions"`
}

// LoadVarsFromEnv checks if each var exists inside os environment variables.
// If it does exist, overrides the existing value
func (c *Config) LoadVarsFromEnv() {

	BaseUrl, ok := os.LookupEnv("KC_URL")
	if ok {
		c.Url = BaseUrl
	}

	ClientID, ok := os.LookupEnv("KC_CLIENT_ID")
	if ok {
		c.ClientID = ClientID
	}

	Secret, ok := os.LookupEnv("KC_CLIENT_SECRET")
	if ok {
		c.Secret = Secret
	}

	Realm, ok := os.LookupEnv("KC_REALM")
	if ok {
		c.Realm = Realm
	}

	CommaSeparatedActions, ok := os.LookupEnv("KC_SETUP_ACTIONS")
	if ok {
		rawActions := strings.Split(CommaSeparatedActions, ",")

		for _, rawAction := range rawActions {

			cleanedAction := strings.ToUpper(strings.TrimSpace(rawAction))
			c.SetupActions = append(c.SetupActions, cleanedAction)
		}
	}
}

// Default returns the default config for the structure
func Default() Config {

	return Config{
		SetupActions: *DefaultActions(),
	}
}
