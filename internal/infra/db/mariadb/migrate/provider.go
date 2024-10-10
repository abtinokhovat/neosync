package migrate

import (
	"embed"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var fsCoreMigrations embed.FS

func Provide() migrate.EmbedFileSystemMigrationSource {
	return migrate.EmbedFileSystemMigrationSource{
		FileSystem: fsCoreMigrations,
		Root:       "migrations",
	}
}
