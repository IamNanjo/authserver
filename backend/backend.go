package backend

import (
	"io/fs"
	"net/http"

	"github.com/IamNanjo/authserver/backend/api"
	"github.com/IamNanjo/authserver/backend/routes"
	"github.com/IamNanjo/authserver/config"

	"github.com/IamNanjo/go-logging"
	"github.com/go-webauthn/webauthn/webauthn"
)

func StartServer(staticFiles fs.FS) {
	webAuthnEnabled := false

	if config.Parsed.WebAuthn.Id != "" && len(config.Parsed.WebAuthn.Origins) != 0 {
		webAuthnEnabled = true
	}

	if webAuthnEnabled {
		api.WebAuthnConfig = &webauthn.Config{
			RPDisplayName: config.Parsed.WebAuthn.DisplayName,
			RPID:          config.Parsed.WebAuthn.Id,
			RPOrigins:     config.Parsed.WebAuthn.Origins,
		}

		webAuthn, err := webauthn.New(api.WebAuthnConfig)
		if err != nil {
			logging.Fatal("Could not initialize Webauthn. %v", err)
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

	http.ListenAndServe(config.Parsed.Address, nil)
}
