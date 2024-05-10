package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "github.com/lib/pq"
)

type DbNoSQLService struct {
	Client *mongo.Client
	DbName string
}

func NewDbNoSQLService(ctx context.Context, dbName string, dbUrl string) (*DbNoSQLService, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		slog.Error(fmt.Sprintf("error while trying to connect to noSQL database '%s'", dbName), "error", err)
		return nil, err
	}

	return &DbNoSQLService{
		Client: client,
		DbName: dbName,
	}, nil
}

func (d *DbNoSQLService) GetInstance() *mongo.Client {
	return d.Client
}

func (d *DbNoSQLService) GetDbName() string {
	return d.DbName
}

func (d *DbNoSQLService) Ping(ctx context.Context) error {
	if err := d.Client.Ping(ctx, nil); err != nil {
		slog.Error(fmt.Sprintf("error while trying to ping noSQL database '%s'", d.DbName), "error", err)
		return err
	}

	return nil
}
