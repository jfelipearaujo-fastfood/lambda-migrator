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
	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, 5432, dbUser, dbPass, dbName)

	slog.Info("connecting to the database")

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

	// listing the tables
	rows, err := conn.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		slog.Error("error while trying to list the tables", "error", err)
		return err
	}
	defer rows.Close()

	tableNames := make([]string, 0)

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			slog.Error("error while trying to scan the table name", "error", err)
			return err
		}
		tableNames = append(tableNames, tableName)
	}

	slog.Info("tables", "tables", tableNames)

	return nil
}

func main() {
	lambda.Start(handler)
}
