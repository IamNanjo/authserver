package backend

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/IamNanjo/authserver/internal/backend/api"
	"github.com/IamNanjo/authserver/internal/backend/api/passkey"
	"github.com/IamNanjo/authserver/internal/backend/api/password"
	"github.com/IamNanjo/authserver/internal/backend/middleware"
	"github.com/IamNanjo/authserver/internal/backend/routes"
	"github.com/IamNanjo/authserver/internal/config"
	"github.com/IamNanjo/authserver/internal/embedded"

	"github.com/IamNanjo/go-logging"
	"github.com/go-webauthn/webauthn/webauthn"
)

func StartServer(ctx context.Context) {
	webAuthnEnabled := false

	if config.Parsed.WebAuthn.Id != "" && len(config.Parsed.WebAuthn.Origins) != 0 {
		webAuthnEnabled = true
	}

	if webAuthnEnabled {
		passkey.WebAuthnConfig = &webauthn.Config{
			RPDisplayName: config.Parsed.WebAuthn.DisplayName,
			RPID:          config.Parsed.WebAuthn.Id,
			RPOrigins:     config.Parsed.WebAuthn.Origins,
		}

		webAuthn, err := webauthn.New(passkey.WebAuthnConfig)
		if err != nil {
			logging.Default.Fatal("Failed to initialize Webauthn. %v", err)
		}

		passkey.WebAuthn = webAuthn
	}

	var handler http.ServeMux
	server := http.Server{
		Addr:        config.Parsed.Address,
		BaseContext: func(l net.Listener) context.Context { return ctx },
		Handler:     &handler,
	}

	handler.Handle("GET /static/", http.FileServer(http.FS(embedded.StaticFiles)))

	handler.HandleFunc("GET /api/user/exists/{$}", middleware.WithMiddleware(api.UserExists, middleware.Logger, middleware.Locale))

	handler.HandleFunc("POST /api/register/password", middleware.WithMiddleware(password.PasswordRegister))
	handler.HandleFunc("POST /api/auth/password", password.PasswordLogin)

	if webAuthnEnabled {
		handler.HandleFunc("POST /api/passkey/register/begin/{$}", passkey.PasskeyRegisterBegin)
		handler.HandleFunc("POST /api/passkey/register/finish/{$}", passkey.PasskeyRegisterFinish)
		handler.HandleFunc("POST /api/passkey/auth/begin/{$}", passkey.PasskeyLoginBegin)
		handler.HandleFunc("POST /api/passkey/auth/finish/{$}", passkey.PasskeyLoginFinish)
	}

	handler.HandleFunc("GET /register/{$}", routes.RegisterPage)
	handler.HandleFunc("GET /error/{$}", routes.Error)
	handler.HandleFunc("GET /{$}", routes.Index)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logging.Default.Err("Failed to shutdown HTTP server: %v", err)
		}
	}()

	<-ctx.Done()
	server.Shutdown(context.Background())
}
