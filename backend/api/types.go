package api

import "github.com/go-webauthn/webauthn/webauthn"

type User struct {
	Id          []byte
	Name        string
	Credentials []webauthn.Credential
}

func (u User) WebAuthnID() []byte {
	return u.Id
}

func (u User) WebAuthnName() string {
	return u.Name
}

func (u User) WebAuthnDisplayName() string {
	return u.Name
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

type EmailAndUsernameRequestBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
