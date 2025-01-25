package api

import (
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

var WebAuthnConfig = &webauthn.Config{
	RPDisplayName: "Auth Service",
	RPID:          "auth.nanjo.dev",
	RPOrigins:     []string{"https://auth.nanjo.dev"},
}

var WebAuthn *webauthn.WebAuthn

func PasskeyBeginRegister(w http.ResponseWriter, r *http.Request) {}

func PasskeyFinishRegister(w http.ResponseWriter, r *http.Request) {}

func PasskeyBeginLogin(w http.ResponseWriter, r *http.Request) {}

func PasskeyFinishLogin(w http.ResponseWriter, r *http.Request) {}
