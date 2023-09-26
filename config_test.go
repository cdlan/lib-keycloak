package keycloak

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadVarsFromEnv(t *testing.T) {

	url := "https://kc.domain.example/auth/"
	cid := "client_id"
	cs 	:= "client_secret"
	r	:= "my_realm"

	// set env vars
	err := os.Setenv("KC_URL", url)
	assert.NoError(t, err)

	err = os.Setenv("KC_CLIENT_ID", cid)
	assert.NoError(t, err)

	err = os.Setenv("KC_CLIENT_SECRET", cs)
	assert.NoError(t, err)

	err = os.Setenv("KC_REALM", r)
	assert.NoError(t, err)

	C := Config{}
	C.LoadVarsFromEnv()

	assert.Equal(t, url, 	C.Url)
	assert.Equal(t, cid, 	C.ClientID)
	assert.Equal(t, cs, 	C.Secret)
	assert.Equal(t, r, 		C.Realm)
}