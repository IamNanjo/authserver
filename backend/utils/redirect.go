package utils

import (
	"net/http"
)

// Uses correct redirect method for HTMX and normal requests.
func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	if r.Header.Get("HX-Request") == "" {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	} else {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(http.StatusOK)
	}
}
