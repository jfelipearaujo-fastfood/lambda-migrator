package orders_db

import (
	"context"
	"log/slog"

	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"

	_ "embed"
)

type Migrator struct {
	DbService database.DatabaseSQLService
}

//go:embed db.sql
var query string

func NewMigrator(dbService database.DatabaseSQLService) *Migrator {
	return &Migrator{
		DbService: dbService,
	}
}

func (m *Migrator) Migrate(ctx context.Context) error {
	dbInstance := m.DbService.GetInstance()

	_, err := dbInstance.Exec(query)
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, "migration completed", "database", m.DbService.GetDbName())

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

	slog.InfoContext(ctx, "tables founded after migration", "tables", tableNames)

	return nil
}
