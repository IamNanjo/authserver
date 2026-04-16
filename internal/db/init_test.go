package db

import (
	"context"
	"os"
	"testing"

	"github.com/IamNanjo/go-logging"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	if err := Initialize(ctx, "/tmp/authserver/authserver_test.db"); err != nil {
		logging.Default.Fatal("Database initialization failed: %v\n", err)
	}

	if err := Db.Ping(); err != nil {
		logging.Default.Fatal("Connection failed%+v\n", err)
	}

	// Truncate tables
	if _, err := Db.ExecContext(ctx, "DELETE FROM App;"); err != nil {
		logging.Default.Err("Delete failed: %v\n", err)
	}
	if _, err := Db.ExecContext(ctx, "DELETE FROM Domain;"); err != nil {
		logging.Default.Err("Delete failed: %v\n", err)
	}
	if _, err := Db.ExecContext(ctx, "DELETE FROM User;"); err != nil {
		logging.Default.Err("Delete failed: %v\n", err)
	}
	if _, err := Db.ExecContext(ctx, "DELETE FROM AppUser;"); err != nil {
		logging.Default.Err("Delete failed: %v\n", err)
	}

	os.Exit(m.Run())
}
