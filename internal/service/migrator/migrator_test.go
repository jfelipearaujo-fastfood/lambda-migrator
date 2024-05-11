package migrator

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jfelipearaujo-org/lambda-migrator/internal/adapter/database/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMigrator(t *testing.T) {
	t.Run("Should migrate the database", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbService := mocks.NewMockDatabaseSQLService(t)

		dbService.On("GetInstance").
			Return(db).
			Once()

		dbService.On("GetDbName").
			Return("test").
			Once()

		mock.ExpectExec("SELECT 1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT (.+) FROM (.+)?information_schema.tables(.+)?").
			WillReturnRows(sqlmock.NewRows([]string{"table_name"}).AddRow("orders"))

		migrator := NewMigrator(dbService)

		// Act
		err = migrator.Migrate(ctx, "SELECT 1")

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
		dbService.AssertExpectations(t)
	})

	t.Run("Should return error if a migration error occurs", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbService := mocks.NewMockDatabaseSQLService(t)

		dbService.On("GetInstance").
			Return(db).
			Once()

		mock.ExpectExec("SELECT 1").
			WillReturnError(assert.AnError)

		migrator := NewMigrator(dbService)

		// Act
		err = migrator.Migrate(ctx, "SELECT 1")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
		dbService.AssertExpectations(t)
	})

	t.Run("Should return error if the query of tables return an error", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbService := mocks.NewMockDatabaseSQLService(t)

		dbService.On("GetInstance").
			Return(db).
			Once()

		dbService.On("GetDbName").
			Return("test").
			Once()

		mock.ExpectExec("SELECT 1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT (.+) FROM (.+)?information_schema.tables(.+)?").
			WillReturnError(assert.AnError)

		migrator := NewMigrator(dbService)

		// Act
		err = migrator.Migrate(ctx, "SELECT 1")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
		dbService.AssertExpectations(t)
	})

	t.Run("Should return error when try to scan the table names", func(t *testing.T) {
		// Arrange
		ctx := context.Background()

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		dbService := mocks.NewMockDatabaseSQLService(t)

		dbService.On("GetInstance").
			Return(db).
			Once()

		dbService.On("GetDbName").
			Return("test").
			Once()

		mock.ExpectExec("SELECT 1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT (.+) FROM (.+)?information_schema.tables(.+)?").
			WillReturnRows(sqlmock.NewRows([]string{"table_name"}).AddRow(nil))

		migrator := NewMigrator(dbService)

		// Act
		err = migrator.Migrate(ctx, "SELECT 1")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, mock.ExpectationsWereMet())
		dbService.AssertExpectations(t)
	})
}
