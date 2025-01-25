package backend

import (
	"github.com/IamNanjo/authserver/backend/api"
	"github.com/IamNanjo/authserver/backend/routes"
	"github.com/go-webauthn/webauthn/webauthn"
	"io/fs"
	"net/http"
)

func StartServer(addr string, staticFiles fs.FS) {
	webAuthn, err := webauthn.New(api.WebAuthnConfig)
	if err != nil {
		panic("Could not initialize WebAuthn")
	}

	api.WebAuthn = webAuthn

	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	http.HandleFunc("/api/auth/password", api.PasswordAuth)
	http.HandleFunc("/api/auth/passkey", api.PasskeyAuth)

	http.HandleFunc("/api/passkey-options", api.PasskeyOptions)

	http.HandleFunc("/", routes.Index)

	http.ListenAndServe(addr, nil)
}
