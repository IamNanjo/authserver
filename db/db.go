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

	dbPath = filepath.Dir(path) + "authserver.db"
}

func Connection() *sqlx.DB {
	if dbPath == "" {
		getPath()
	}

	return sqlx.MustConnect("sqlite", dbPath)
}
