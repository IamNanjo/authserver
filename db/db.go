package db

import (
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
	"strconv"
)

var dbPath string = ""

// Create tables according to migrations.
// Migrations are run in ascending order based on filename.
// Path can be nil to use default path.
func Initialize(path *string) error {
	if path == nil {
		getDefaultPath()
	} else {
		dbPath = *path
	}

	connection := Connection()

	latestMigration := Migration{Id: 0}

	connection.Get(&latestMigration, "SELECT * FROM Migration ORDER BY id DESC LIMIT 1")

	migrations, err := GetMigrations(latestMigration.Id)

	if err != nil {
		os.Stderr.WriteString("Could not get migrations. Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, migration := range migrations {
		tx, err := connection.Begin()
		if err != nil {
			return err
		}

		_, err = tx.Exec(string(migration.content))
		if err != nil {
			os.Stderr.WriteString("Migration failed: " + migration.filename + "\n" + err.Error() + "\n")
			os.Exit(1)
		}

		err = tx.Commit()
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

// Returns connection and panics on error
func Connection() *sqlx.DB {
	return sqlx.MustConnect("sqlite", dbPath)
}
