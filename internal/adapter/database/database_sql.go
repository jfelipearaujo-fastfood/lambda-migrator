package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/lib/pq"
)

type DbSQLService struct {
	Client *sql.DB
	DbName string
}

func NewDbSQLService(ctx context.Context, engine string, dbName string, dbUrl string) (*DbSQLService, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	client, err := sql.Open(engine, dbUrl)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error while trying to connect to SQLdatabase '%s'", dbName), "error", err)
		return nil, err
	}

	return &DbSQLService{
		Client: client,
		DbName: dbName,
	}, nil
}

func (d *DbSQLService) GetInstance() *sql.DB {
	return d.Client
}

func (d *DbSQLService) GetDbName() string {
	return d.DbName
}

func (d *DbSQLService) Ping(ctx context.Context) error {
	if err := d.Client.PingContext(ctx); err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("error while trying to ping SQL database '%s'", d.DbName), "error", err)
		return err
	}

	return nil
}
