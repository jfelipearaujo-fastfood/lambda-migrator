package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestNoSQLDatabase(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Should connect to noSQL database", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		dbName := "test"
		dbUrl := "mongodb://test:test@localhost:27017/test?sslmode=disable"

		service, err := NewDbNoSQLService(ctx, dbName, dbUrl)
		assert.NoError(t, err)

		// Act
		res := service.GetInstance()

		// Assert
		assert.NotNil(mt, res)
	})

	mt.Run("Should return database name", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		dbName := "test"
		dbUrl := "mongodb://test:test@localhost:27017/test?sslmode=disable"

		service, err := NewDbNoSQLService(ctx, dbName, dbUrl)
		assert.NoError(t, err)

		// Act
		res := service.GetDbName()

		// Assert
		assert.Equal(mt, service.DbName, res)
	})

	mt.Run("Should return error if a connection error occurs", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		dbName := "test"
		dbUrl := "wrongUrl"

		// Act
		service, err := NewDbNoSQLService(ctx, dbName, dbUrl)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, service)
	})

	mt.Run("Should ping noSQL database", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		dbName := "test"
		dbUrl := "mongodb://test:test@localhost:27017/test?sslmode=disable"

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		service, err := NewDbNoSQLService(ctx, dbName, dbUrl)
		assert.NoError(t, err)

		service.Client = mt.Client

		// Act
		err = service.Ping(ctx)

		// Assert
		assert.NoError(t, err)
	})

	mt.Run("Should return error if a ping error occurs", func(mt *mtest.T) {
		// Arrange
		ctx := context.Background()

		dbName := "test"
		dbUrl := "mongodb://test:test@localhost:27017/test?sslmode=disable"

		service, err := NewDbNoSQLService(ctx, dbName, dbUrl)
		assert.NoError(t, err)

		service.Client = mt.Client

		// Act
		err = service.Ping(ctx)

		// Assert
		assert.Error(t, err)
	})
}
