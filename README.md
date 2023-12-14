# keycloak
Wrapper around [github.com/Nerzal/gocloak](https://github.com/Nerzal/gocloak/tree/main) KeyCloak Go Library

## Install
```shell
go get github.com/cdlan/lib-keycloak
```

## Usage
```golang
import keycloak "github.com/cdlan/lib-keycloak"

// initialize object
kc := keycloak.Config{
    Url:        "https://kc.domain.example/auth/",
    ClientID:   "my_client_id",
    Secret:     "my_secret",
    Realm:      "my_realm",
}

ctx := context.Background()

// search user by email
user, err := kc.SearchUser(ctx, "email@domain.example")
if err != nil {
    panic(err)
}
```

## Configuration
The module has built-in support for [github.com/spf13/viper](https://github.com/spf13/viper).

There is also the possibility to load configurations from ENV VARS using
```go
import keycloak "github.com/cdlan/lib-keycloak"

// create default object
kc := keycloak.Default()

// load from ENV VARS
kc.LoadVarsFromEnv()
```

### ENV VARS
- KC_URL: url of KC in the `https://kc.domain.example/auth/` form.
- KC_CLIENT_ID
- KC_CLIENT_SECRET
- KC_REALM
- KC_SETUP_ACTIONS: comma separated actions to execute when new user is created. see [here](https://stackoverflow.com/questions/45328003/setting-required-action-on-new-keycloak-user-account#45433846)
