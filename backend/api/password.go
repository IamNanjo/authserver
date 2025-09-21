package api

import (
	"net/http"
	"time"

	"github.com/IamNanjo/authserver/backend/utils"
	"github.com/IamNanjo/authserver/db"
	"github.com/IamNanjo/authserver/hash"
)

// Ensures cookie is valid. Also ensures session exists in DB.
func AuthCookieIsValid(cookie *http.Cookie) bool {
	err := cookie.Valid()
	if err != nil {
		return false
	}

	// sessionId := cookie.Value

	// TODO: Ensure session exists and has not expired

	return true
}

func PasswordAuth(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	app := query.Get("app")
	redirect := query.Get("redirect")

	if app == "" {
		utils.Error(w, r, http.StatusBadRequest, "App parameter not set")
		return
	}

	if redirect == "" {
		utils.Error(w, r, http.StatusBadRequest, "Redirect parameter not set")
		return
	}

	emailOrUsername := r.PostFormValue("email-or-username")
	password := r.PostFormValue("password")

	if emailOrUsername == "" {
		utils.Error(w, r, http.StatusBadRequest, "Email or username is missing")
		return
	} else if password == "" {
		utils.Error(w, r, http.StatusBadRequest, "Password is missing")
		return
	}

	user, err := db.Q().GetUserByEmailOrUsername(r.Context(), &emailOrUsername)
	if err != nil {
		utils.Error(w, r, http.StatusNotFound, "User not found")
		return
	}

	if user.Password == nil {
		utils.Error(w, r, http.StatusBadRequest, "Invalid login method for user "+user.Name)
	}

	passwordIsCorrect, err := hash.HashValidate([]byte(password), *user.Password)
	if err != nil || !passwordIsCorrect {
		utils.Error(w, r, http.StatusUnauthorized, "Incorrect password")
		return
	}

	sessionId, err := db.GenerateId(128)
	if err != nil {
		utils.Error(w, r, http.StatusInternalServerError, "Could not generate session ID. Please try again")
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

	utils.Redirect(w, r, "/")
}
