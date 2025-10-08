package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/IamNanjo/authserver/backend"
	"github.com/IamNanjo/authserver/config"
	"github.com/IamNanjo/authserver/db"
)

//go:embed static/**
var staticFiles embed.FS

func main() {
	var err error

	appConfig := config.Parse()

	if appConfig.DatabaseURI == nil || *appConfig.DatabaseURI == "" {
		err = db.Initialize(appConfig.DatabaseURI)
	} else {
		err = db.Initialize(nil)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not initialize database. Error: %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Starting server on %s\n", *appConfig.Address)
	backend.StartServer(appConfig, staticFiles)
}
