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

func PasskeyOptions(w http.ResponseWriter, r *http.Request) {

}
