package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"strconv"

	_ "modernc.org/sqlite"
)

var dbPath string = ""
var connection *sql.DB
var dbQ *DBQueries
var dbTx *DBTransactions

// Create tables according to migrations.
// Migrations are run in ascending order based on filename.
// Path can be nil to use default path.
func Initialize(path *string) error {
	if path == nil {
		getDefaultPath()
	} else {
		dbPath = *path
	}

	tx := Tx()

	latestMigration, err := tx.GetLatestMigration(context.Background())
	if err != nil {
		latestMigration = 0
	}

	migrations, err := GetMigrations(latestMigration)

	if err != nil {
		os.Stderr.WriteString("Could not get migrations. Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, migration := range migrations {
		_, err = tx.Tx.Exec(string(migration.content))
		if err != nil {
			os.Stderr.WriteString("Migration failed: " + migration.filename + "\n" + err.Error() + "\n")
			os.Exit(1)
		}

		err = tx.Tx.Commit()
	}

	migrationsFinished := len(migrations)
	os.Stdout.WriteString("Finished " + strconv.Itoa(migrationsFinished) + " database migration")
	if migrationsFinished == 1 {
		os.Stdout.WriteString("\n")
	} else {
		os.Stdout.WriteString("s\n")
	}

	return nil
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

type DBQueries struct {
	Queries
	db *sql.DB
}

type DBTransactions struct {
	Queries
	Tx *sql.Tx
}

// Ensures connection is active and return queries. Panics on error.
func Q() *DBQueries {
	// Return existing connection if it is still alive
	if dbQ != nil && dbQ.db.Ping() != nil {
		return dbQ
	}

	connection, err := sql.Open("sqlite", dbPath)
	if err != nil {
		os.Stderr.WriteString("Could not connect to database " + dbPath + "\n")
		os.Exit(1)
	}

	// Pragma options
	connection.Exec(`
		PRAGMA foreign_keys = ON;
		PRAGMA journal_mode = WAL;
		PRAGMA auto_vacuum = FULL;
	`)

	dbQ = &DBQueries{Queries: *New(connection), db: connection}
	return dbQ
}

// Ensures connection is active and returns queries with active transaction. Panics on error.
func Tx() *DBTransactions {
	var err error
	if connection == nil || connection.Ping() != nil {
		connection, err = sql.Open("sqlite", dbPath)
	}
	if err != nil {
		os.Stderr.WriteString("Could not connect to database " + dbPath + "\n")
		os.Exit(1)
	}

	// Pragma options
	connection.Exec(`
		PRAGMA foreign_keys = ON;
		PRAGMA journal_mode = WAL;
		PRAGMA auto_vacuum = FULL;
	`)

	tx, err := connection.Begin()
	if err != nil {
		os.Stderr.WriteString("Could not start transaction")
		os.Exit(1)
	}

	dbTx = &DBTransactions{Queries: *New(tx), Tx: tx}
	return dbTx
}
