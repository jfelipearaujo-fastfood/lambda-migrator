package main

import (
	"context"
	_ "embed"
	"log/slog"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/environment/loader"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/handler"
)

const dbEngine string = "postgres"

//go:embed scripts/orders_db_init.sql
var queryOrderDbInit string

//go:embed scripts/payments_db_init.sql
var queryPaymentsDbInit string

//go:embed scripts/productions_db_init.sql
var queryProductionsDbInit string

//go:embed scripts/customers_db_init.sql
var queryCustomersDbInit string

func main() {
	ctx := context.Background()

	loader := loader.NewLoader()

	config, err := loader.GetEnvironment(ctx)
	if err != nil {
		slog.Error("error loading environment", "error", err)
	}

	ordersDbService, err := database.NewDbSQLService(ctx, dbEngine, config.DbOrdersConfig.Name, config.DbOrdersConfig.Url)
	if err != nil {
		slog.Error("error creating database service", "error", err)
	}

	paymentsDbService, err := database.NewDbSQLService(ctx, dbEngine, config.DbPaymentsConfig.Name, config.DbPaymentsConfig.Url)
	if err != nil {
		slog.Error("error creating database service", "error", err)
	}

	productionsDbService, err := database.NewDbSQLService(ctx, dbEngine, config.DbProductionsConfig.Name, config.DbProductionsConfig.Url)
	if err != nil {
		slog.Error("error creating database service", "error", err)
	}

	customersDbService, err := database.NewDbSQLService(ctx, dbEngine, config.DbCustomersConfig.Name, config.DbCustomersConfig.Url)
	if err != nil {
		slog.Error("error creating database service", "error", err)
	}

	handler := handler.NewHandler(
		ordersDbService,
		paymentsDbService,
		productionsDbService,
		customersDbService,
		queryOrderDbInit,
		queryPaymentsDbInit,
		queryProductionsDbInit,
		queryCustomersDbInit,
	)

	lambda.Start(handler.Handle)
}
