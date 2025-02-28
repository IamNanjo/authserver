# Authentication server

[![Go](https://github.com/IamNanjo/auth-server/actions/workflows/go.yml/badge.svg)](https://github.com/IamNanjo/auth-server/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/IamNanjo/authserver)](https://goreportcard.com/report/github.com/IamNanjo/authserver)

```
This project is in very early development.
It is not currently recommended to use this project anywhere.
```

A simple authentication server written in Go.
Supports password authentication using argon2 hashes as well as passkey authentication.

The authentication page ensures that the redirect path is on the same domain as one of the domains that is configured for the selected app.
This prevents redirecting to malicious sites after authentication, even if the authentication cookie is not valid for the selected app.

## Database

When a new user signs up, they will be able to use any apps that use this authentication server and they will be able to connect their own apps to it.
Connected apps will have a list of domains where the authentication cookie is valid.
The creator of the app will have the ability to change the app name, description and domains freely, as well as add other existing users as managers.

There is always at least one administrator user, who can delete any app connections and users.
The first created user will always be an administrator.
Administrators can give other existing users the administrator role.

The app uses an SQLite3 database that is created in the same directory where the binary is (`./authserver.db`)

![ER diagram](ER-diagram.svg)
