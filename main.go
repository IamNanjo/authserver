package main

import (
	"embed"
	"github.com/IamNanjo/authserver/backend"
	"os"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	addr := os.Getenv("AUTH_SERVER_PORT")
	if addr == "" {
		addr = ":8080"
	}

	sFiles, err := staticFiles.ReadDir("static")
	if err != nil {
		panic(err)
	}

	for _, file := range sFiles {
		os.Stdout.WriteString("./" + file.Name() + "\n")
	}

	os.Stdout.WriteString("Listening on " + addr + "\n")
	backend.StartServer(addr, staticFiles)
}
