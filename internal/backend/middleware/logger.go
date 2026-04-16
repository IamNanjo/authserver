package middleware

import "net/http"

func Logger(w http.ResponseWriter, r *http.Request) bool {
	return true
}
