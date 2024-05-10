package database

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSQLDatabase(t *testing.T) {
	t.Run("Should connect to SQL database", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, _, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbName := "test"
		dbUrl := "postgres://test:test@localhost:54321/test?sslmode=disable"

		service, err := NewDbSQLService(ctx, "postgres", dbName, dbUrl)
		assert.NoError(t, err)

		service.Client = db

		// Act
		res := service.GetInstance()

		// Assert
		assert.NotNil(t, res)
	})

	t.Run("Should return database name", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, _, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbName := "test"
		dbUrl := "postgres://test:test@localhost:54321/test?sslmode=disable"

		service, err := NewDbSQLService(ctx, "postgres", dbName, dbUrl)
		assert.NoError(t, err)

		service.Client = db

		// Act
		res := service.GetDbName()

		// Assert
		assert.Equal(t, service.DbName, res)
	})

	t.Run("Should return error if a connection error occurs", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, _, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbName := "test"
		dbUrl := "postgres://test:test@localhost:54321/test?sslmode=disable"

		// Act
		service, err := NewDbSQLService(ctx, "wrong", dbName, dbUrl)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, service)
	})

	t.Run("Should ping the database", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPing().WillDelayFor(time.Second)

		dbName := "test"
		dbUrl := "postgres://test:test@localhost:54321/test?sslmode=disable"

		service, err := NewDbSQLService(ctx, "postgres", dbName, dbUrl)
		assert.NoError(t, err)

		service.Client = db

		// Act
		err = service.Ping(ctx)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})

	t.Run("Should return error if a ping error occurs", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPing().WillReturnError(assert.AnError)

		dbName := "test"
		dbUrl := "postgres://test:test@localhost:54321/test?sslmode=disable"

		service, err := NewDbSQLService(ctx, "postgres", dbName, dbUrl)
		assert.NoError(t, err)

		service.Client = db

		// Act
		err = service.Ping(ctx)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
	})
}
