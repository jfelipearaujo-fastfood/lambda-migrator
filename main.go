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
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// connect to the database
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbName)

	slog.Info("connecting to the database", "connection_string", connectionStr)

	conn, err := sql.Open("postgres", connectionStr)
	if err != nil {
		slog.Error("error while trying to connect to the database", "error", err)
		return err
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		slog.Error("error while trying to ping the database", "error", err)
		return err
	}

	// execute the sql
	slog.Info("executing the query")

	res, err := conn.Exec(query)
	if err != nil {
		slog.Error("error while trying to execute the query", "error", err)
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		slog.Error("error while trying to get the affected rows", "error", err)
		return err
	}

	slog.Info("completed", "affected_rows", affectedRows)

	return nil
}

func main() {
	lambda.Start(handler)
}
