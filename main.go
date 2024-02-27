package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"github.com/aws/aws-lambda-go/lambda"
)

//go:embed scripts/init.sql
var query string

func handler(ctx context.Context) error {
	// load env vars
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// connect to the database
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	slog.Info("connecting to the database")

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		slog.Error("error while trying to connect to the database", "error", err)
		return err
	}
	defer conn.Close()

	// execute the sql
	slog.Info("executing the query")

	_, err = conn.Exec(query)
	if err != nil {
		slog.Error("error while trying to execute the query", "error", err)
	}

	slog.Info("completed")

	return nil
}

func main() {
	lambda.Start(handler)
}
