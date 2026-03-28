package db

import (
	"testing"
)

func TestSetup(t *testing.T) {
	err := Initialize("/tmp/authserver_test.db")
	if err != nil {
		t.Fatalf("Database initialization failed %v", err)
	}

	// Truncate tables and reset ID counters
	Q().db.Exec(`
		DELETE FROM App;
		DELETE FROM Domain;
		DELETE FROM User;
		DELETE FROM AppUser;
	`)
}
