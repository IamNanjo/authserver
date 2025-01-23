package db

import (
	"embed"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed migrations/**
var migrationFs embed.FS

type MigrationFile struct {
	id       int
	filename string
	content  []byte
}

func GetMigrations(latest int) ([]MigrationFile, error) {
	migrationFiles, err := migrationFs.ReadDir("migrations")
	if err != nil {
		return nil, err
	}

	migrations := make([]MigrationFile, 0, len(migrationFiles))

	for _, f := range migrationFiles {
		filename := f.Name()

		id, err := strconv.Atoi(strings.TrimSuffix(filename, ".sql"))
		if err != nil {
			continue
		}

		if f.Type().IsDir() || !strings.HasSuffix(filename, ".sql") || id <= latest {
			continue
		}

		content, err := migrationFs.ReadFile("migrations/" + filename)
		if err != nil {
			os.Stderr.WriteString("Could not read migration file: " + filename + "\n")
			os.Exit(1)
		}

		migrations = append(migrations, MigrationFile{id: id, filename: filename, content: content})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].id < migrations[j].id
	})

	return migrations, nil
}
