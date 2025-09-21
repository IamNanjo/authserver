package api

import (
	"encoding/json"
	"net/http"

	"github.com/IamNanjo/authserver/db"
	"github.com/go-webauthn/webauthn/webauthn"
)

var WebAuthnConfig *webauthn.Config
var WebAuthn *webauthn.WebAuthn

func PasskeyRegisterBegin(w http.ResponseWriter, r *http.Request) {
	body := EmailAndUsernameRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var emailOrUsername string

	if body.Email != "" {
		emailOrUsername = body.Email
	} else if body.Username != "" {
		emailOrUsername = body.Username
	} else {
		return
	}

	user, err := db.Q().GetUserByEmailOrUsername(r.Context(), &emailOrUsername)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	wUser := User{Id: []byte(user.Id), Name: user.Name, Credentials: []webauthn.Credential{}}

	options, _, err := WebAuthn.BeginRegistration(wUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(options)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func PasskeyRegisterFinish(w http.ResponseWriter, r *http.Request) {
}

func PasskeyLoginBegin(w http.ResponseWriter, r *http.Request) {
}

func PasskeyLoginFinish(w http.ResponseWriter, r *http.Request) {
}
