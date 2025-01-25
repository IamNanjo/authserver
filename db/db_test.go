package db

import (
	"context"
	"testing"
)

func TestConnection(t *testing.T) {
	connection := Connection()

	err := connection.PingContext(context.Background())
	if err != nil {
		t.Errorf("Database connection failed. %v", err)
	}
}
