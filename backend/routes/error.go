package routes

import (
	"context"
	"github.com/IamNanjo/authserver/pages"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, err string) {
	w.WriteHeader(statusCode)
	page := pages.Error(err, "")
	page.Render(context.Background(), w)
}
