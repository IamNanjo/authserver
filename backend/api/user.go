package api

import (
	"net/http"

	"github.com/IamNanjo/authserver/backend/utils"
	"github.com/IamNanjo/authserver/db"
)

func UserExists(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	email := query.Get("email")
	username := query.Get("username")

	if email != "" {
		_, err := db.GetUserByEmail(email)
		if err == nil {
			utils.Error(w, r, http.StatusConflict, "A user with this email already exists")
			return
		}
	}

	if username != "" {
		_, err := db.GetUserByUsername(username)
		if err == nil {
			utils.Error(w, r, http.StatusConflict, "A user with this username already exists")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
