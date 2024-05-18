package tests

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/handler"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "embed"

	_ "github.com/lib/pq"
)

var opts = godog.Options{
	Format:      "pretty",
	Paths:       []string{"features"},
	Output:      colors.Colored(os.Stdout),
	Concurrency: 4,
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestFeatures(t *testing.T) {
	o := opts
	o.TestingT = t

	status := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options:             &o,
	}.Run()

	if status == 2 {
		t.SkipNow()
	}

	if status != 0 {
		t.Fatalf("zero status code expected, %d received", status)
	}
}

// Steps
const dbEngine string = "postgres"

//go:embed testdata/orders_db_init.sql
var queryOrderDbInit string

//go:embed testdata/payments_db_init.sql
var queryPaymentsDbInit string

//go:embed testdata/productions_db_init.sql
var queryProductionsDbInit string

//go:embed testdata/customers_db_init.sql
var queryCustomersDbInit string

const featureKey CtxKeyType = "feature"

type feature struct {
	urls map[string]string
}

var state = NewState[feature](featureKey)

func iHaveToMigrateTheDatabases(ctx context.Context) (context.Context, error) {
	feat := state.retrieve(ctx)

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
	feat := state.retrieve(ctx)

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
	feat := state.retrieve(ctx)

	expectedOrdersTables := []string{"orders", "order_items", "order_payments"}
	expectedPaymentsTables := []string{"payments", "payment_items"}
	expectedProductionsTables := []string{"orders", "order_items"}
	expectedCustomersTables := []string{"customers"}

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

type testContext struct {
	network    *testcontainers.DockerNetwork
	containers []testcontainers.Container
}

var (
	containers = make(map[string]testContext)
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		network, err := network.New(ctx, network.WithCheckDuplicate(), network.WithDriver("bridge"))
		if err != nil {
			return ctx, err
		}

		ordersContainer, ctx, err := createPostgresContainer(ctx, network, "order")
		if err != nil {
			return ctx, err
		}

		paymentsContainer, ctx, err := createPostgresContainer(ctx, network, "payment")
		if err != nil {
			return ctx, err
		}

		productionsContainer, ctx, err := createPostgresContainer(ctx, network, "production")
		if err != nil {
			return ctx, err
		}

		customersContainer, ctx, err := createPostgresContainer(ctx, network, "customer")
		if err != nil {
			return ctx, err
		}

		containers[sc.Id] = testContext{
			network: network,
			containers: []testcontainers.Container{
				ordersContainer,
				paymentsContainer,
				productionsContainer,
				customersContainer,
			},
		}

		return ctx, nil
	})

	ctx.Step(`^I have to migrate the databases$`, iHaveToMigrateTheDatabases)
	ctx.Step(`^I migrate the databases$`, iMigrateTheDatabases)
	ctx.Step(`^the databases should be migrated$`, theDatabasesShouldBeMigrated)

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if err != nil {
			return ctx, err
		}

		tc := containers[sc.Id]

		for _, c := range tc.containers {
			err := c.Terminate(ctx)
			if err != nil {
				return ctx, err
			}
		}

		err = tc.network.Remove(ctx)

		return ctx, err
	})
}

func createPostgresContainer(
	ctx context.Context,
	network *testcontainers.DockerNetwork,
	dbPrefix string,
) (testcontainers.Container, context.Context, error) {
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: "postgres:16.0",
			ExposedPorts: []string{
				"5432",
			},
			Env: map[string]string{
				"POSTGRES_DB":       fmt.Sprintf("%s_db", dbPrefix),
				"POSTGRES_USER":     dbPrefix,
				"POSTGRES_PASSWORD": dbPrefix,
			},
			Networks: []string{
				network.Name,
			},
			NetworkAliases: map[string][]string{
				network.Name: {
					"test",
				},
			},
			WaitingFor: wait.ForLog("PostgreSQL init process complete; ready for start up").WithStartupTimeout(120 * time.Second),
		},
		Started: true,
	})
	if err != nil {
		return nil, ctx, fmt.Errorf("failed to start postgres container: %w", err)
	}

	postgresIp, err := container.Host(ctx)
	if err != nil {
		return nil, ctx, fmt.Errorf("failed to get postgres ip: %w", err)
	}

	postgresPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, ctx, fmt.Errorf("failed to get postgres port: %w", err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s_db?sslmode=disable", dbPrefix, dbPrefix, postgresIp, postgresPort.Port(), dbPrefix)

	feat := state.retrieve(ctx)

	if feat.urls == nil {
		feat.urls = make(map[string]string)
	}

	feat.urls[dbPrefix] = connStr

	return container, state.enrich(ctx, feat), nil
}
