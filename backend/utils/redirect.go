package utils

import (
	"net/http"
)

// Uses correct redirect method for HTMX and normal requests.
func Redirect(w http.ResponseWriter, r *http.Request, url string, statusCode int) {
	if r.Header.Get("HX-Request") == "" {
		http.Redirect(w, r, url, statusCode)
	} else {
		w.Header().Set("HX-Location", url)
		w.WriteHeader(http.StatusOK)
	}
}
