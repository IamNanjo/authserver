package db

import (
	"embed"
	"sort"
	"strconv"
	"strings"

	"github.com/IamNanjo/go-logging"
)

//go:embed migrations/*.sql
var migrationFs embed.FS

type MigrationFile struct {
	id       int
	filename string
	content  string
}

func GetMigrations(latest int64) ([]MigrationFile, error) {
	migrationFiles, err := migrationFs.ReadDir("migrations")
	if err != nil {
		return nil, err
	}

	migrations := make([]MigrationFile, 0, len(migrationFiles))

	for _, f := range migrationFiles {
		filename := f.Name()

		id, err := strconv.Atoi(strings.TrimSuffix(filename, ".sql"))
		if err != nil || int64(id) <= latest {
			continue
		}

		content, err := migrationFs.ReadFile("migrations/" + filename)
		if err != nil {
			logging.Fatal("Could not read migration file: %s\n", filename)
		}

		migrations = append(migrations, MigrationFile{id: id, filename: filename, content: string(content)})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].id < migrations[j].id
	})

	return migrations, nil
}
