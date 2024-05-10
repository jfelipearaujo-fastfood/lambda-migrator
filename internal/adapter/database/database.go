package database

import (
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/lib/pq"
)

type DatabaseSQLService interface {
	GetInstance() *sql.DB
	GetDbName() string
	Ping(ctx context.Context) error
}

type DatabaseNoSQLService interface {
	GetInstance() *mongo.Client
	GetDbName() string
	Ping(ctx context.Context) error
}
