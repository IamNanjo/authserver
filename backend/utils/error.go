package utils

import (
	"encoding/json"
	"net/http"

	"github.com/IamNanjo/authserver/components"
)

func Error(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	accept := r.Header.Get("Accept")

	if accept == "*/*" && r.Header.Get("HX-Request") != "" {
		accept = "text/html"
	}

	w.WriteHeader(statusCode)

	switch accept {
	case "text/html":
		components.Error(message, nil).Render(r.Context(), w)
	default:
		json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	}
}
