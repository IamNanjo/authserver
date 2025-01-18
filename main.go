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

	db.Initialize(nil)

	os.Stdout.WriteString("Listening on " + addr + "\n")
	backend.StartServer(addr, staticFiles)
}
