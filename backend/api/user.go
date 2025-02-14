package api

import (
	"encoding/json"
	"github.com/IamNanjo/authserver/db"
	"net/http"
)

func UserExists(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	email := query.Get("email")
	username := query.Get("username")

	if email != "" {
		_, err := db.GetUserByEmail(email)
		if err == nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(ErrorResponse{
				Reason: "email",
				Error:  "A user with this email already exists",
			})
			return
		}
	}

	if username != "" {
		_, err := db.GetUserByUsername(username)
		if err == nil {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(ErrorResponse{
				Reason: "username",
				Error:  "A user with this username already exists",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
