package db

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"

	"github.com/IamNanjo/go-logging"
	"github.com/IamNanjo/go-logging/pkg/format"
	_ "modernc.org/sqlite"
)

var Q *Queries
var Db *sql.DB

// Create tables according to migrations.
// Migrations are run in ascending order based on filename.
// Path can be nil to use default path.
func Initialize(ctx context.Context, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return format.Err("Failed to create directories for DB: %v", err)
	}

	connection, err := sql.Open(
		"sqlite",
		"file:"+path+"?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=auto_vacuum(FULL)",
	)
	if err != nil {
		return format.Err("Failed to open database connection %w", err)
	}

	Q = New(connection)
	Db = connection

	tx, err := connection.BeginTx(ctx, nil)
	if err != nil {
		return format.Err("Failed to start transaction: %v", err)
	}
	qTx := Q.WithTx(tx)

	latestMigration, err := qTx.GetLatestMigration(ctx)
	if err != nil {
		latestMigration = 0
	}

	migrations, err := GetMigrations(latestMigration)
	if err != nil {
		return format.Err("Failed to get migrations %w", err)
	}

	for _, migration := range migrations {
		if _, err = tx.ExecContext(ctx, migration.content); err != nil {
			return format.Err("Migration %s failed %w", migration.filename, err)
		}
		if _, err := tx.ExecContext(ctx, "INSERT INTO Migration VALUES (?)", migration.id); err != nil {
			return format.Err("Failed to set latest migration as %d: %v", migration.id, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return format.Err("Failed to commit migrations %w", err)
	}

	if migrationsFinished := len(migrations); migrationsFinished != 0 {
		logging.Default.Ok("Finished %d database migration(s)\n", migrationsFinished)
	}

	return nil
}
