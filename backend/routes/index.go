package routes

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/IamNanjo/authserver/backend/utils"
	"github.com/IamNanjo/authserver/db"
	"github.com/IamNanjo/authserver/pages"
)

func ValidateRedirectURL(domains []db.Domain, url url.URL) bool {
	redirectURLIsAllowed := false

	for _, domain := range domains {
		hostname := url.Hostname()
		if domain.Name == hostname || (domain.Name[0] == '.' && strings.HasSuffix(domain.Name, url.Hostname())) {
			redirectURLIsAllowed = true
		}
	}

	return redirectURLIsAllowed
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	appId := query.Get("app")
	if appId == "" {
		err := pages.Index().Render(r.Context(), w)
		if err != nil {
			utils.Redirect(w, r, "/error?status=500&error=Could not render the page")
			return
		}
		return
	}

	redirectTo := query.Get("redirect")
	if redirectTo == "" {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. No redirect page specified")
		return
	}

	redirectURL, err := url.Parse(redirectTo)
	if err != nil {
		utils.Redirect(w, r, "/error?status=400&error=Invalid redirect URL")
		return
	}

	app, err := db.Q().GetApp(r.Context(), appId)
	if err != nil {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. App not found")
		return
	}

	domains, err := db.Q().GetAppDomains(r.Context(), appId)
	if err != nil {
		utils.Redirect(w, r, "/error?status=400&error=App domains not configured")
		return
	}

	redirectURLIsAllowed := ValidateRedirectURL(domains, *redirectURL)

	if !redirectURLIsAllowed {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. Redirect page is not on the app domains")
		return
	}

	err = pages.Auth(app, redirectTo).Render(r.Context(), w)
	if err != nil {
		utils.Redirect(w, r, "/error?status=500&error=Could not render the page")
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
		utils.Redirect(w, r, "/error?status=405&error="+"Invalid method "+r.Method+" for route "+r.URL.Path)
	}
}
