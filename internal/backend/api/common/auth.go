package common

import "net/http"

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

func AuthStatus(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("Auth")
	if err != nil || !AuthCookieIsValid(authCookie) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
