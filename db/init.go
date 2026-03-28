package db

import (
	"context"
	"database/sql"

	"github.com/IamNanjo/go-logging"
	"github.com/IamNanjo/go-logging/pkg/format"
	_ "modernc.org/sqlite"
)

var dbPath string
var connection *sql.DB
var dbQ *DBQueries
var dbTx *DBTransactions

// Create tables according to migrations.
// Migrations are run in ascending order based on filename.
// Path can be nil to use default path.
func Initialize(path string) error {
	dbPath = path
	tx := Tx()

	var err error
	connection, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return format.Err("Failed to open database connection %w", err)
	}

	latestMigration, err := tx.GetLatestMigration(context.Background())
	if err != nil {
		latestMigration = 0
	}

	migrations, err := GetMigrations(latestMigration)
	if err != nil {
		return format.Err("Could not get migrations %w", err)
	}

	for _, migration := range migrations {
		_, err = tx.Tx.Exec(migration.content)
		if err != nil {
			return format.Err("Migration %s failed %w", migration.filename, err)
		}
	}

	err = tx.Tx.Commit()
	if err != nil {
		return format.Err("Failed to commit migrations %w", err)
	}

	migrationsFinished := len(migrations)
	if migrationsFinished != 0 {
		return format.Err("Finished %d database migration(s)\n", migrationsFinished)
	}

	return nil
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
		logging.Fatal("Could not connect to database %s\n", dbPath)
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
		logging.Fatal("Could not connect to database %s\n", dbPath)
	}

	// Pragma options
	connection.Exec(`
		PRAGMA foreign_keys = ON;
		PRAGMA journal_mode = WAL;
		PRAGMA auto_vacuum = FULL;
	`)

	tx, err := connection.Begin()
	if err != nil {
		logging.Fatal("Could not start transaction: %+v\n", err)
	}

	dbTx = &DBTransactions{Queries: *New(tx), Tx: tx}
	return dbTx
}
