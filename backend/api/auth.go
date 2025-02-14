package api

import "net/http"

func AuthStatus(w http.ResponseWriter, r *http.Request) {
	authCookie, err := r.Cookie("Auth")
	if err != nil || !AuthCookieIsValid(authCookie) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
