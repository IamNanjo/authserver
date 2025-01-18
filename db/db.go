package db

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
)

var dbPath = ""

// Create tables according to Schema (schema.go).
// Path can be nil to use default path.
func Initialize(path *string) {
	if path == nil {
		getDefaultPath()
	} else {
		dbPath = *path
	}
	Connection().MustExec(Schema)
	os.Stdout.WriteString("Database initialized according to schema\n")
}

// Default path is <path of executable>/authserver.db.
// Evaluates symlinks.
func getDefaultPath() {
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

// Returns connection and panics on error
func Connection() *sqlx.DB {
	return sqlx.MustConnect("sqlite", dbPath)
}
