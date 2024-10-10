package migrator

import (
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"log/slog"
	"neosync/internal/logger"
)

type Config struct {
	Apply bool
}

type Migrator struct {
	cfg              Config
	tableName        string
	dialect          string
	connectionString string
	migrations       migrate.MigrationSource
}

func New(cfg Config, connectionString string, migrations migrate.MigrationSource) Migrator {
	return Migrator{
		cfg:              cfg,
		tableName:        "migrations",
		connectionString: connectionString,
		dialect:          "mysql",
		migrations:       migrations,
	}
}

func (m *Migrator) Up() {
	if !m.cfg.Apply {
		return
	}

	db, err := sql.Open(m.dialect, m.connectionString)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	migrate.SetTable(m.tableName)

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Up)
	if err != nil {
		panic(fmt.Errorf("can't apply migrations: %v", err))
	}

	logger.L().Info("successfully applied migrations!", slog.Int("count", n))
}

func (m *Migrator) Down() {
	if !m.cfg.Apply {
		return
	}

	db, err := sql.Open(m.dialect, m.connectionString)
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}

	migrate.SetTable(m.tableName)

	n, err := migrate.Exec(db, m.dialect, m.migrations, migrate.Down)
	if err != nil {
		panic(fmt.Errorf("can't rollback migrations: %v", err))
	}

	logger.L().Info("rollback migrations!", slog.Int("count", n))
}
