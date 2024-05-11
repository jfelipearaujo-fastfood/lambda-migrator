package handler

import (
	"context"
	"log/slog"

	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/service/migrator"
)

type Handler struct {
	ordersDbService      database.DatabaseSQLService
	paymentsDbService    database.DatabaseSQLService
	productionsDbService database.DatabaseSQLService

	queryOrdersDbInit      string
	queryPaymentsDbInit    string
	queryProductionsDbInit string
}

func NewHandler(
	ordersDbService database.DatabaseSQLService,
	paymentsDbService database.DatabaseSQLService,
	productionsDbService database.DatabaseSQLService,
	queryOrdersDbInit string,
	queryPaymentsDbInit string,
	queryProductionsDbInit string,
) *Handler {
	return &Handler{
		ordersDbService:      ordersDbService,
		paymentsDbService:    paymentsDbService,
		productionsDbService: productionsDbService,

		queryOrdersDbInit:      queryOrdersDbInit,
		queryPaymentsDbInit:    queryPaymentsDbInit,
		queryProductionsDbInit: queryProductionsDbInit,
	}
}

func (h *Handler) Handle(ctx context.Context) error {
	slog.InfoContext(ctx, "starting migrations")

	err := runMigration(ctx, h.ordersDbService, h.queryOrdersDbInit)
	if err != nil {
		slog.ErrorContext(ctx, "error migrating database", "database", h.ordersDbService.GetDbName(), "error", err)
	}

	err = runMigration(ctx, h.paymentsDbService, h.queryPaymentsDbInit)
	if err != nil {
		slog.ErrorContext(ctx, "error migrating database", "database", h.paymentsDbService.GetDbName(), "error", err)
	}

	err = runMigration(ctx, h.productionsDbService, h.queryProductionsDbInit)
	if err != nil {
		slog.ErrorContext(ctx, "error migrating database", "database", h.productionsDbService.GetDbName(), "error", err)
	}

	return nil
}

func runMigration(ctx context.Context, db database.DatabaseSQLService, query string) error {
	migrator := migrator.NewMigrator(db)

	if err := migrator.Migrate(ctx, query); err != nil {
		return err
	}

	return nil
}
