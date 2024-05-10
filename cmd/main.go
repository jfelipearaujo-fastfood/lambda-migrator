package main

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/environment/loader"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/service"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/service/orders_db"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/service/payments_db"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/service/productions_db"
)

const (
	SQL_DB_ENGINE = "postgres"
)

func handler(ctx context.Context) error {
	loader := loader.NewLoader()

	config, err := loader.GetEnvironment(ctx)
	if err != nil {
		return err
	}

	ordersDb, err := database.NewDbSQLService(ctx, SQL_DB_ENGINE, config.DbOrdersConfig.Name, config.DbOrdersConfig.Url)
	if err != nil {
		return err
	}

	paymentsDb, err := database.NewDbSQLService(ctx, SQL_DB_ENGINE, config.DbPaymentsConfig.Name, config.DbPaymentsConfig.Url)
	if err != nil {
		return err
	}

	productionsDb, err := database.NewDbSQLService(ctx, SQL_DB_ENGINE, config.DbProductionsConfig.Name, config.DbProductionsConfig.Url)
	if err != nil {
		return err
	}

	migrators := []service.Migrator{
		orders_db.NewMigrator(ordersDb),
		payments_db.NewMigrator(paymentsDb),
		productions_db.NewMigrator(productionsDb),
	}

	slog.InfoContext(ctx, "starting migration", "migrators", migrators)

	for _, migrator := range migrators {
		if err := migrator.Migrate(ctx); err != nil {
			return err
		}
	}

	slog.InfoContext(ctx, "migration completed")

	return nil
}

func main() {
	lambda.Start(handler)
}
