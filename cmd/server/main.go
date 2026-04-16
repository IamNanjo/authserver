package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/IamNanjo/authserver/internal/backend"
	"github.com/IamNanjo/authserver/internal/config"
	"github.com/IamNanjo/authserver/internal/db"

	"github.com/IamNanjo/go-logging"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := config.Parse(); err != nil {
		logging.Default.Fatal("Failed to parse config %v", err)
	}

	if err := db.Initialize(ctx, config.Parsed.DatabasePath); err != nil {
		logging.Default.Fatal("Failed to initialize database %+v\n", err)
	}

	logging.Default.Info("Starting server on %s\n", config.Parsed.Address)
	backend.StartServer(ctx)
}
