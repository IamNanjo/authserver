package routes

import (
	"net/http"
	"strconv"

	"github.com/IamNanjo/authserver/pages"
)

func Error(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	status := query.Get("status")
	errorMessage := query.Get("error")

	var statusCode int
	var err error

	if errorMessage == "" {
		errorMessage = "Unknown error"
	}

	statusCode, err = strconv.Atoi(status)
	if err != nil {
		statusCode = http.StatusBadRequest
	}

	w.WriteHeader(statusCode)
	err = pages.Error(errorMessage).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
