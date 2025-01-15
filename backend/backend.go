package backend

import (
	"github.com/IamNanjo/authserver/backend/api"
	"github.com/IamNanjo/authserver/backend/routes"
	"io/fs"
	"net/http"
)

func StartServer(addr string, staticFiles fs.FS) {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	http.HandleFunc("/api/auth/password", api.PasswordAuth)

	http.HandleFunc("/", routes.Index)

	http.ListenAndServe(addr, nil)
}
