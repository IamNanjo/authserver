package main

import (
	"embed"
	"flag"
	"github.com/IamNanjo/authserver/backend"
	"github.com/IamNanjo/authserver/db"
	"os"
)

//go:embed static/**
var staticFiles embed.FS

func main() {
	var err error
	addr := os.Getenv("AUTH_SERVER_PORT")
	if addr == "" {
		addr = ":8080"
	}

	dbPath := os.Getenv("AUTH_SERVER_DB")

	// Override environment variables with cli flag
	flag.StringVar(&addr, "addr", addr, "Listen address")
	flag.StringVar(&dbPath, "db", dbPath, "Database path")

	flag.Parse()

	if dbPath != "" {
		err = db.Initialize(&dbPath)
	} else {
		err = db.Initialize(nil)
	}

	if err != nil {
		os.Stderr.WriteString("Could not initialize database. Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString("Starting server on " + addr + "\n")
	backend.StartServer(addr, staticFiles)
}
