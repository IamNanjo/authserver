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
	addr := os.Getenv("AUTH_SERVER_PORT")
	if addr == "" {
		addr = ":8080"
	}

	// Override environment variable with cli flag
	flag.StringVar(&addr, "addr", addr, "Listen address")

	flag.Parse()

	err := db.Initialize(nil)

	if err != nil {
		os.Stderr.WriteString("Could not initialize database. Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString("Starting server on " + addr + "\n")
	backend.StartServer(addr, staticFiles)
}
