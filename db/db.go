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
