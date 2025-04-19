package routes

import (
	"net/http"
	"net/url"

	"github.com/IamNanjo/authserver/backend/utils"
	"github.com/IamNanjo/authserver/db"
	"github.com/IamNanjo/authserver/pages"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	appId := query.Get("app")
	redirect := query.Get("redirect")

	if appId == "" {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. No app ID specified", http.StatusMovedPermanently)
		return
	}

	if redirect == "" {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. No redirect page specified", http.StatusMovedPermanently)
		return
	}

	redirectUrl, err := url.Parse(redirect)
	if err != nil {
		utils.Redirect(w, r, "/error?status=400&error=Invalid redirect URL", http.StatusMovedPermanently)
		return
	}

	app, err := db.GetAppById(appId)
	if err != nil {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. App not found", http.StatusMovedPermanently)
		return
	}

	if !ValidateRedirectURL(app.Domains, *redirectUrl) {
		utils.Redirect(w, r, "/error?status=400&error=Invalid authentication URL. Redirect page is not on the app domains", http.StatusMovedPermanently)
		return
	}

	err = pages.Register(app, redirect).Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
