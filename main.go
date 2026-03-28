package main

import (
	"embed"

	"github.com/IamNanjo/authserver/backend"
	"github.com/IamNanjo/authserver/config"
	"github.com/IamNanjo/authserver/db"

	"github.com/IamNanjo/go-logging"
)

//go:embed static/**
var staticFiles embed.FS

func main() {
	if err := config.Parse(); err != nil {
		logging.Fatal("Failed to parse config %v", err)
	}

	if err := db.Initialize(config.Parsed.DatabaseURI); err != nil {
		logging.Fatal("Could not initialize database %+v\n", err)
	}

	logging.Pending("Starting server on %s\n", config.Parsed.Address)
	backend.StartServer(staticFiles)
}
