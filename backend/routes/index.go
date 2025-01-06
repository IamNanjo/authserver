package routes

import (
	"context"
	"github.com/IamNanjo/authserver/pages"
	"github.com/a-h/templ"
	"net/http"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	app := query.Get("app")
	if app == "" {
		Error(w, http.StatusBadRequest, "Invalid authentication URL. No app ID specified")
		return
	}

	redirectTo := query.Get("redirect")
	if redirectTo == "" {
		Error(w, http.StatusBadRequest, "Invalid authentication URL. No redirect page specified")
		return
	}

	page := pages.Index(pages.PageDataIndex{App: app, RedirectTo: templ.SafeURL(redirectTo)})
	err := page.Render(context.Background(), w)
	if err != nil {
		Error(w, http.StatusInternalServerError, "Could not render the page")
		return
	}
}

func postIndex(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getIndex(w, r)
	case "POST":
		postIndex(w, r)
	default:
		Error(w, http.StatusMethodNotAllowed, "Invalid method "+r.Method+" for route "+r.URL.Path)
	}
}
