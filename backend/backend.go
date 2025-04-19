package backend

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/IamNanjo/authserver/backend/api"
	"github.com/IamNanjo/authserver/backend/routes"
	"github.com/go-webauthn/webauthn/webauthn"
)

func StartServer(addr string, staticFiles fs.FS) {
	webAuthnEnabled := false
	webAuthnId := os.Getenv("AUTHSERVER_WEBAUTHN_RPID")
	webAuthnOrigins := strings.Split(os.Getenv("AUTHSERVER_WEBAUTHN_RPORIGINS"), ",")

	for i, origin := range webAuthnOrigins {
		webAuthnOrigins[i] = strings.TrimSpace(origin)
	}

	if webAuthnId != "" && len(webAuthnOrigins) != 0 {
		webAuthnEnabled = true
	}

	if webAuthnEnabled {
		api.WebAuthnConfig = &webauthn.Config{
			RPDisplayName: "Authentication Service",
			RPID:          webAuthnId,
			RPOrigins:     webAuthnOrigins,
		}

		webAuthn, err := webauthn.New(api.WebAuthnConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not initialize Webauthn. %v", err)
			os.Exit(1)
		}

		api.WebAuthn = webAuthn
	}

	http.Handle("GET /static/", http.FileServer(http.FS(staticFiles)))

	http.HandleFunc("GET /api/user/exists/{$}", api.UserExists)

	http.HandleFunc("POST /api/register/password", api.PasswordRegister)
	http.HandleFunc("POST /api/auth/password", api.PasswordAuth)

	if webAuthnEnabled {
		http.HandleFunc("POST /api/passkey/register/begin/{$}", api.PasskeyRegisterBegin)
		http.HandleFunc("POST /api/passkey/register/finish/{$}", api.PasskeyRegisterFinish)
		http.HandleFunc("POST /api/passkey/auth/begin/{$}", api.PasskeyLoginBegin)
		http.HandleFunc("POST /api/passkey/auth/finish/{$}", api.PasskeyLoginFinish)
	}

	http.HandleFunc("GET /register/{$}", routes.RegisterPage)
	http.HandleFunc("GET /error/{$}", routes.Error)
	http.HandleFunc("GET /{$}", routes.Index)

	http.ListenAndServe(addr, nil)
}
