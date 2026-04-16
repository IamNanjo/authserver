package db

import (
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/IamNanjo/authserver/internal/embedded"
	"github.com/IamNanjo/go-logging/pkg/format"
)

type MigrationFile struct {
	id       int
	filename string
	content  string
}

const dir = "migrations"

func GetMigrations(latest int64) ([]MigrationFile, error) {
	migrationFiles, err := embedded.DbMigrations.ReadDir(dir)
	if err != nil {
		return nil, format.Err("Failed to read migrations directory: %v", err)
	}

	migrations := make([]MigrationFile, 0, len(migrationFiles))

	for _, f := range migrationFiles {
		filename := f.Name()

		id, err := strconv.Atoi(strings.TrimSuffix(filename, ".sql"))
		if err != nil || int64(id) <= latest {
			continue
		}

		content, err := embedded.DbMigrations.ReadFile(filepath.Join(dir, filename))
		if err != nil {
			return nil, format.Err("Failed to read migration file %q: %v", filename, err)
		}

		migrations = append(migrations, MigrationFile{id: id, filename: filename, content: string(content)})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].id < migrations[j].id
	})

	return migrations, nil
}
