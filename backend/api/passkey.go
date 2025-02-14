package api

import (
	"encoding/json"
	"github.com/IamNanjo/authserver/db"
	"github.com/go-webauthn/webauthn/webauthn"
	"net/http"
)

var WebAuthnConfig *webauthn.Config
var WebAuthn *webauthn.WebAuthn

func PasskeyBeginRegister(w http.ResponseWriter, r *http.Request) {
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

	user, err := db.GetUserByEmailOrUsername(emailOrUsername)
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

func PasskeyFinishRegister(w http.ResponseWriter, r *http.Request) {
}

func PasskeyBeginLogin(w http.ResponseWriter, r *http.Request) {
}

func PasskeyFinishLogin(w http.ResponseWriter, r *http.Request) {
}
