package backend

import (
	"github.com/IamNanjo/authserver/backend/routes"
	"io/fs"
	"net/http"
)

func StartServer(addr string, staticFiles fs.FS) {
	http.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	http.HandleFunc("/", routes.Index)

	http.ListenAndServe(addr, nil)
}
