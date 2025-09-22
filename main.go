package main

import (
	"embed"
	"os"

	"github.com/IamNanjo/authserver/backend"
	"github.com/IamNanjo/authserver/config"
	"github.com/IamNanjo/authserver/db"
)

//go:embed static/**
var staticFiles embed.FS

func main() {
	var err error

	appConfig := (&config.Config{}).ParseConfig()

	if appConfig.DatabaseURI != "" {
		err = db.Initialize(&appConfig.DatabaseURI)
	} else {
		err = db.Initialize(nil)
	}

	if err != nil {
		os.Stderr.WriteString("Could not initialize database. Error: " + err.Error() + "\n")
		os.Exit(1)
	}

	os.Stdout.WriteString("Starting server on " + appConfig.Address + "\n")
	backend.StartServer(appConfig, staticFiles)
}
