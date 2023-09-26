# keycloak
KeyCloak Go Library

## Install
**TODO**

## Usage
```go
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
