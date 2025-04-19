package api

import (
	"net/http"
	"strings"

	"github.com/IamNanjo/authserver/backend/utils"
	"github.com/IamNanjo/authserver/db"
)

func PasswordRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.PostForm.Get("email")
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == "" {
		utils.Error(w, r, http.StatusBadRequest, "No username provided")
		return
	}

	if email != "" && !utils.IsValidEmail(email) {
		utils.Error(w, r, http.StatusBadRequest, "Invalid email provided")
		return
	}

	_, err := db.CreateUser(username, email, password)
	if err != nil {
		errMsg := err.Error()
		after, conflict := strings.CutPrefix(errMsg, "constraint failed: UNIQUE constraint failed: User.")

		if !conflict {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		reason := strings.Split(after, " ")[0]
		errorMessage := "Unknown conflict with existing user"

		switch reason {
		case "name":
			errorMessage = "A user with this username already exists"
		case "email":
			errorMessage = "A user with this email already exists"
		}

		utils.Error(w, r, http.StatusConflict, errorMessage)

		return
	}

	utils.Redirect(w, r, "/", http.StatusMovedPermanently)
}
