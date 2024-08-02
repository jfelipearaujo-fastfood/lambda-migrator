package tests

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/handler"
	"github.com/jfelipearaujo/testcontainers/pkg/container"
	"github.com/jfelipearaujo/testcontainers/pkg/container/postgres"
	"github.com/jfelipearaujo/testcontainers/pkg/network"
	"github.com/jfelipearaujo/testcontainers/pkg/state"
	"github.com/jfelipearaujo/testcontainers/pkg/testsuite"
	"github.com/testcontainers/testcontainers-go"

	_ "embed"

	_ "github.com/lib/pq"
)

const dbEngine string = "postgres"

//go:embed testdata/orders_db_init.sql
var queryOrderDbInit string

//go:embed testdata/payments_db_init.sql
var queryPaymentsDbInit string

//go:embed testdata/productions_db_init.sql
var queryProductionsDbInit string

//go:embed testdata/customers_db_init.sql
var queryCustomersDbInit string

type feature struct {
	urls map[string]string
}

var testState = state.NewState[feature]()

var containers = container.NewGroup()

func TestFeatures(t *testing.T) {
	testsuite.NewTestSuite(t,
		initializeScenario,
		testsuite.WithPaths("features"),
		testsuite.WithConcurrency(0),
	)
}

func initializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		ntwrkDefinition := network.NewNetwork()

		network, err := ntwrkDefinition.Build(ctx)
		if err != nil {
			return ctx, fmt.Errorf("failed to build the network: %w", err)
		}

		ordersContainer, ctx, err := createPostgresContainer(ctx, ntwrkDefinition, network, "order")
		if err != nil {
			return ctx, err
		}

		paymentsContainer, ctx, err := createPostgresContainer(ctx, ntwrkDefinition, network, "payment")
		if err != nil {
			return ctx, err
		}

		productionsContainer, ctx, err := createPostgresContainer(ctx, ntwrkDefinition, network, "production")
		if err != nil {
			return ctx, err
		}

		customersContainer, ctx, err := createPostgresContainer(ctx, ntwrkDefinition, network, "customer")
		if err != nil {
			return ctx, err
		}

		containers[sc.Id] = container.BuildGroupContainer(
			container.WithDockerContainer(ordersContainer),
			container.WithDockerContainer(paymentsContainer),
			container.WithDockerContainer(productionsContainer),
			container.WithDockerContainer(customersContainer),
		)

		return ctx, nil
	})

	ctx.Step(`^I have to migrate the databases$`, iHaveToMigrateTheDatabases)
	ctx.Step(`^I migrate the databases$`, iMigrateTheDatabases)
	ctx.Step(`^the databases should be migrated$`, theDatabasesShouldBeMigrated)

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		group := containers[sc.Id]

		return container.DestroyGroup(ctx, group)
	})
}

func iHaveToMigrateTheDatabases(ctx context.Context) (context.Context, error) {
	feat := testState.Retrieve(ctx)

	err := pingDatabase(ctx, feat.urls["order"])
	if err != nil {
		return ctx, fmt.Errorf("failed to ping orders database: %w", err)
	}

	err = pingDatabase(ctx, feat.urls["payment"])
	if err != nil {
		return ctx, fmt.Errorf("failed to ping payments database: %w", err)
	}

	err = pingDatabase(ctx, feat.urls["production"])
	if err != nil {
		return ctx, fmt.Errorf("failed to ping productions database: %w", err)
	}

	err = pingDatabase(ctx, feat.urls["customer"])
	if err != nil {
		return ctx, fmt.Errorf("failed to ping customers database: %w", err)
	}

	return ctx, nil
}

func iMigrateTheDatabases(ctx context.Context) (context.Context, error) {
	feat := testState.Retrieve(ctx)

	ordersDbService, err := database.NewDbSQLService(ctx, dbEngine, "orders", feat.urls["order"])
	if err != nil {
		return ctx, fmt.Errorf("error creating database service: %v", err)
	}

	paymentsDbService, err := database.NewDbSQLService(ctx, dbEngine, "payments", feat.urls["payment"])
	if err != nil {
		return ctx, fmt.Errorf("error creating database service: %v", err)
	}

	productionsDbService, err := database.NewDbSQLService(ctx, dbEngine, "productions", feat.urls["production"])
	if err != nil {
		return ctx, fmt.Errorf("error creating database service: %v", err)
	}

	customersDbService, err := database.NewDbSQLService(ctx, dbEngine, "customers", feat.urls["customer"])
	if err != nil {
		return ctx, fmt.Errorf("error creating database service: %v", err)
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

	return ctx, handler.Handle(ctx)
}

func theDatabasesShouldBeMigrated(ctx context.Context) (context.Context, error) {
	feat := testState.Retrieve(ctx)

	expectedOrdersTables := []string{"orders", "order_items", "order_payments"}
	expectedPaymentsTables := []string{"payments", "payment_items"}
	expectedProductionsTables := []string{"orders", "order_items"}
	expectedCustomersTables := []string{"customers", "customer_deletion_requests"}

	ordersTableNames, err := getDatabaseTables(ctx, feat.urls["order"])
	if err != nil {
		return ctx, fmt.Errorf("failed to get orders tables: %w", err)
	}

	paymentsTableNames, err := getDatabaseTables(ctx, feat.urls["payment"])
	if err != nil {
		return ctx, fmt.Errorf("failed to get payments tables: %w", err)
	}

	productionsTableNames, err := getDatabaseTables(ctx, feat.urls["production"])
	if err != nil {
		return ctx, fmt.Errorf("failed to get productions tables: %w", err)
	}

	customersTableNames, err := getDatabaseTables(ctx, feat.urls["customer"])
	if err != nil {
		return ctx, fmt.Errorf("failed to get customers tables: %w", err)
	}

	if !reflect.DeepEqual(expectedOrdersTables, ordersTableNames) {
		return ctx, fmt.Errorf("expected orders tables %v, got %v", expectedOrdersTables, ordersTableNames)
	}

	if !reflect.DeepEqual(expectedPaymentsTables, paymentsTableNames) {
		return ctx, fmt.Errorf("expected payments tables %v, got %v", expectedPaymentsTables, paymentsTableNames)
	}

	if !reflect.DeepEqual(expectedProductionsTables, productionsTableNames) {
		return ctx, fmt.Errorf("expected productions tables %v, got %v", expectedProductionsTables, productionsTableNames)
	}

	if !reflect.DeepEqual(expectedCustomersTables, customersTableNames) {
		return ctx, fmt.Errorf("expected customers tables %v, got %v", expectedCustomersTables, customersTableNames)
	}

	return ctx, nil
}

func pingDatabase(ctx context.Context, dbUrl string) error {
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer conn.Close()

	if err := conn.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping postgres: %w", err)
	}

	return nil
}

func getDatabaseTables(ctx context.Context, dbUrl string) ([]string, error) {
	tableNames := make([]string, 0)

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return tableNames, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return tableNames, fmt.Errorf("failed to query postgres: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return tableNames, fmt.Errorf("failed to scan postgres: %w", err)
		}
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}

func createPostgresContainer(
	ctx context.Context,
	ntwrkDefinition *network.Network,
	network *testcontainers.DockerNetwork,
	dbPrefix string,
) (testcontainers.Container, context.Context, error) {
	dbName := fmt.Sprintf("%s_db", dbPrefix)

	pgDefinition := container.NewContainerDefinition(
		container.WithNetwork(ntwrkDefinition.Alias, network),
		postgres.WithPostgresContainer(),
		container.WithEnvVars(map[string]string{
			"POSTGRES_DB":       dbName,
			"POSTGRES_USER":     dbPrefix,
			"POSTGRES_PASSWORD": dbPrefix,
		}),
		container.WithForceWaitDuration(2*time.Second),
	)

	pgContainer, err := pgDefinition.BuildContainer(ctx)
	if err != nil {
		return nil, ctx, err
	}

	connString, err := postgres.BuildExternalConnectionString(ctx,
		pgContainer,
		postgres.WithNetwork(ntwrkDefinition),
		postgres.WithDatabase(dbName),
		postgres.WithUser(dbPrefix),
		postgres.WithPass(dbPrefix),
	)
	if err != nil {
		return nil, ctx, err
	}

	feat := testState.Retrieve(ctx)

	if feat.urls == nil {
		feat.urls = make(map[string]string)
	}

	feat.urls[dbPrefix] = connString

	return pgContainer, testState.Enrich(ctx, feat), nil
}
