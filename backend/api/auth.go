package api

import (
	"context"
	"github.com/IamNanjo/authserver/components"
	"net/http"
)

func PasswordAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	emailOrUsername := r.PostFormValue("email-or-username")
	password := r.PostFormValue("password")

	if emailOrUsername == "" {
		components.Error("Email or username is missing", nil).Render(context.Background(), w)
		return
	} else if password == "" {
		components.Error("Password is missing", nil).Render(context.Background(), w)
		return
	}

	w.Header().Set("HX-Location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}
