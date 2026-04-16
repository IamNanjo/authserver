package embedded

import "embed"

//go:embed migrations/*.sql
var DbMigrations embed.FS
