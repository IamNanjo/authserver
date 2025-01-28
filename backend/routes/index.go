package routes

import (
	"context"
	"github.com/IamNanjo/authserver/db"
	"github.com/IamNanjo/authserver/pages"
	"net/http"
	"net/url"
	"strings"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	appId := query.Get("app")
	if appId == "" {
		err := pages.Index().Render(context.Background(), w)
		if err != nil {
			http.Redirect(w, r, "/error?status=500&error=Could not render the page", http.StatusMovedPermanently)
			http.Redirect(w, r, "/error?status=500&error=", http.StatusMovedPermanently)
			return
		}
		return
	}

	redirectTo := query.Get("redirect")
	if redirectTo == "" {
		http.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. No redirect page specified", http.StatusMovedPermanently)
		return
	}

	redirectURL, err := url.Parse(redirectTo)
	if err != nil {
		http.Redirect(w, r, "/error?status=400&error=Invalid redirect URL", http.StatusMovedPermanently)
		return
	}

	app, err := db.GetAppById(appId)
	if err != nil {
		http.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. App not found", http.StatusMovedPermanently)
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
		http.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. Redirect page is not on the app domains", http.StatusMovedPermanently)
		return
	}

	err = pages.Auth(app, redirectTo).Render(context.Background(), w)
	if err != nil {
		http.Redirect(w, r, "/error?status=500&error=Could not render the page", http.StatusMovedPermanently)
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			getIndex(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			err := pages.Error("Page not found").Render(r.Context(), w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	default:
		http.Redirect(w, r, "/error?status=405&error="+"Invalid method "+r.Method+" for route "+r.URL.Path, http.StatusMovedPermanently)
	}
}
