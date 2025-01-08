package routes

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/IamNanjo/authserver/db"
	"github.com/IamNanjo/authserver/pages"
	"github.com/a-h/templ"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	appId := query.Get("app")
	if appId == "" {
		Error(w, http.StatusBadRequest, "Invalid authentication URL. No app ID specified")
		return
	}

	redirectTo := query.Get("redirect")
	if redirectTo == "" {
		Error(w, http.StatusBadRequest, "Invalid authentication URL. No redirect page specified")
		return
	}

	redirectURL, err := url.Parse(redirectTo)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid redirect URL")
		return
	}

	app, err := db.GetApp(appId)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid authentication URL. App not found")
		return
	}

	redirectURLIsAllowed := false

	for _, domain := range app.Domains {
		hostname := redirectURL.Hostname()
		if domain.Name == hostname || (domain.Name[0] == '.' && strings.HasSuffix(domain.Name, redirectURL.Hostname())) {
			redirectURLIsAllowed = true
		}
	}

	if !redirectURLIsAllowed {
		Error(w, http.StatusBadRequest, "Invalid authentication URL. Redirect page is not on the app domains")
		return
	}

	page := pages.Index(pages.PageDataIndex{App: app, RedirectTo: templ.SafeURL(redirectTo)})
	err = page.Render(context.Background(), w)
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
