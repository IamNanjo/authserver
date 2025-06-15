package db

import (
	"testing"
)

func TestConnection(t *testing.T) {
	connection := Connection()

	err := connection.Ping()
	if err != nil {
		t.Errorf("Database connection failed. %v", err)
	}
}
