package migrator

import (
	"context"
	"log/slog"

	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"

	_ "embed"
)

type Migrator struct {
	DbService database.DatabaseSQLService
}

func NewMigrator(dbService database.DatabaseSQLService) *Migrator {
	return &Migrator{
		DbService: dbService,
	}
}

func (m *Migrator) Migrate(ctx context.Context, query string) error {
	dbInstance := m.DbService.GetInstance()

	_, err := dbInstance.Exec(query)
	if err != nil {
		return err
	}

	rows, err := dbInstance.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return err
	}
	defer rows.Close()

	tableNames := make([]string, 0)

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		tableNames = append(tableNames, tableName)
	}

	slog.InfoContext(ctx, "tables founded after migration", "database", m.DbService.GetDbName(), "tables", tableNames)

	return nil
}
