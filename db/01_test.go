package db

import (
	"os"
	"strings"
	"testing"
)

func TestSetup(t *testing.T) {
	dbPath, _ = os.Getwd()
	dbPath = strings.TrimSuffix(dbPath, "db")
	dbPath += "dist/authserver_test" + ".db"

	err := Initialize(&dbPath)
	if err != nil {
		t.Errorf("Database initialization failed. %v", err)
	}

	// Truncate tables and reset ID counters
	Q().db.Exec(`
		DELETE FROM App;
		DELETE FROM Domain;
		DELETE FROM User;
		DELETE FROM AppUser;
	`)
}
