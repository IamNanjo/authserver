package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/IamNanjo/authserver/components"
	"github.com/IamNanjo/authserver/db"
)

func PasswordRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	expectsJsonResponse := strings.HasPrefix(r.Header.Get("Accept"), "application/json")

	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if username == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := "No username provided"
		if expectsJsonResponse {
			json.NewEncoder(w).Encode(ErrorResponse{Reason: "username", Error: errorMessage})
		} else {
			components.Error(errorMessage, nil).Render(r.Context(), w)
		}
		return
	}

	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := db.CreateUser(username, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
