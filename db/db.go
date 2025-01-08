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
	db := Connection()
	db.MustExec(Schema)
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

func GetApp(id string) (App, error) {
	app := App{}

	err := Connection().Get(&app, "SELECT * FROM App WHERE id=$1 LIMIT 1", id)

	return app, err
}

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
