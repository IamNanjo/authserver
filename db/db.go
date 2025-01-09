package db

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

var dbPath = ""

// Create tables according to Schema (schema.go)
func Initialize() {
	Connection().MustExec(Schema)
	os.Stdout.WriteString("Database initialized according to schema\n")
}

func getPath() {
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	path, err := filepath.EvalSymlinks(exe)
	if err != nil {
		panic(err)
	}

	dbPath = filepath.Dir(path) + string(os.PathSeparator) + "authserver.db"
}

func Connection() *sqlx.DB {
	if dbPath == "" {
		getPath()
	}

	return sqlx.MustConnect("sqlite", dbPath)
}

// Populates App.Domains
func GetApp(id string) (App, error) {
	app := App{}

	conn := Connection()
	err := conn.Get(&app, "SELECT * FROM App WHERE id=$1 LIMIT 1", id)
	err = conn.Select(&app.Domains, "SELECT * FROM Domain WHERE app=$1", id)

	return app, err
}

// Does not populate App.Domains
func GetApps() ([]App, error) {
	apps := []App{}

	err := Connection().Select(&apps, "SELECT * FROM App WHERE visibility = 1")

	return apps, err
}

func GetAppUsers(id string) ([]UserWithAppRole, error) {
	users := []UserWithAppRole{}

	err := Connection().Select(
		&users,
		`SELECT
			id,
			name,
			password,
			email,
			u.role as role,
			au.role as app_role
		FROM User u
		INNER JOIN AppUser au
			ON u.id = au.user
		WHERE app=$1`, id,
	)

	return users, err
}
