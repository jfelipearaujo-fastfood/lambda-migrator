package main

import (
	"context"
	_ "embed"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/environment"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/environment/loader"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/service/migrator"
)

type MigrateConfig struct {
	DbEngine string
	DbConfig *environment.DatabaseConfig
	Query    string
}

//go:embed scripts/orders_db_init.sql
var queryOrderDbInit string

//go:embed scripts/payments_db_init.sql
var queryPaymentsDbInit string

//go:embed scripts/productions_db_init.sql
var queryProductionsDbInit string

func handler(ctx context.Context) error {
	loader := loader.NewLoader()

	config, err := loader.GetEnvironment(ctx)
	if err != nil {
		return err
	}

	migrationConfigs := []MigrateConfig{
		{
			DbEngine: "postgres",
			DbConfig: config.DbOrdersConfig,
			Query:    queryOrderDbInit,
		},
		{
			DbEngine: "postgres",
			DbConfig: config.DbPaymentsConfig,
			Query:    queryPaymentsDbInit,
		},
		{
			DbEngine: "postgres",
			DbConfig: config.DbProductionsConfig,
			Query:    queryProductionsDbInit,
		},
	}

	for _, migrationConfig := range migrationConfigs {
		slog.InfoContext(ctx, "starting migration", "engine", migrationConfig.DbEngine, "database", migrationConfig.DbConfig.Name)

		db, err := database.NewDbSQLService(ctx, migrationConfig.DbEngine, migrationConfig.DbConfig.Name, migrationConfig.DbConfig.Url)
		if err != nil {
			slog.ErrorContext(ctx, "error creating database service", "error", err)
			continue // we don't want to stop the migration if one of the databases fails
		}

		migrator := migrator.NewMigrator(db)

		if err := migrator.Migrate(ctx, migrationConfig.Query); err != nil {
			slog.ErrorContext(ctx, "error migrating database", "error", err)
			continue
		}

		slog.InfoContext(ctx, "migration completed", "engine", migrationConfig.DbEngine, "database", migrationConfig.DbConfig.Name)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
