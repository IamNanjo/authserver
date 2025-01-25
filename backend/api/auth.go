package api

import (
	"github.com/IamNanjo/authserver/components"
	"github.com/IamNanjo/authserver/db"
	"github.com/IamNanjo/authserver/hash"
	"net/http"
	"time"
)

func PasswordAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	app := query.Get("app")
	redirect := query.Get("redirect")

	if app == "" {
		components.Error("App parameter not set", nil).Render(r.Context(), w)
		return
	}

	if redirect == "" {
		components.Error("Redirect parameter not set", nil).Render(r.Context(), w)
		return
	}

	emailOrUsername := r.PostFormValue("email-or-username")
	password := r.PostFormValue("password")

	if emailOrUsername == "" {
		components.Error("Email or username is missing", nil).Render(r.Context(), w)
		return
	} else if password == "" {
		components.Error("Password is missing", nil).Render(r.Context(), w)
		return
	}

	user, err := db.GetUserByEmailOrUsername(emailOrUsername)
	if err != nil {
		components.Error("User not found", nil).Render(r.Context(), w)
		return
	}

	passwordIsCorrect, err := hash.HashValidate([]byte(password), user.Password)
	if err != nil || !passwordIsCorrect {
		components.Error("Incorrect password", nil).Render(r.Context(), w)
		return
	}

	sessionId, err := db.GenerateId(128)
	if err != nil {
		components.Error("Could not generate session ID. Please try again", nil).Render(r.Context(), w)
		return
	}

	maxAge := 60 * 60 * 24 * 90

	cookie := http.Cookie{
		Name:     "Auth",
		Path:     "/",
		Value:    sessionId,
		MaxAge:   maxAge,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("HX-Location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}

func PasskeyAuth(w http.ResponseWriter, r *http.Request) {}
