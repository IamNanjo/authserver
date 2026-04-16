package password

import (
	"net/http"
	"strings"

	"github.com/IamNanjo/authserver/internal/backend/utils"
	"github.com/IamNanjo/authserver/internal/db"
)

func PasswordRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()

	email := r.PostForm.Get("email")
	username := r.PostForm.Get("username")
	_ = r.PostForm.Get("password")

	if username == "" {
		utils.Error(w, r, http.StatusBadRequest, "No username provided")
		return
	}

	if email != "" && !utils.IsValidEmail(email) {
		utils.Error(w, r, http.StatusBadRequest, "Invalid email provided")
		return
	}

	tx, err := db.Db.BeginTx(ctx, nil)
	defer tx.Commit()
	txQ := db.Q.WithTx(tx)
	_, err = txQ.CreateUser(r.Context(), db.CreateUserParams{Id: "", Name: username, Email: &email})
	// TODO: Add password for user
	// txQ.AddPassword(ctx, db.AddPasswordParams{Password: &password})
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

	utils.Redirect(w, r, "/")
}
